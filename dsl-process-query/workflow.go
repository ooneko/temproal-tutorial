/*
Copyright 2021 Baidu ACG/SDC/DET Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dsl_process_query

import (
	"go.temporal.io/sdk/workflow"
	"sync"
	"time"
)

type (
	Workflow struct {
		Variables      map[string]string
		Root           Statement
		activityStatus map[string]bool
		lock           sync.RWMutex
	}

	Statement struct {
		Activity *ActivityInvocation
		Sequence *Sequence
		Parallel *Parallel
	}

	Sequence struct {
		Elements []*Statement
	}

	Parallel struct {
		Branches []*Statement
	}

	ActivityInvocation struct {
		Name      string
		Arguments []string
		Result    string
	}

	executable interface {
		execute(ctx workflow.Context, bindings map[string]string, w *Workflow) error
	}
	activities interface {
		activities() []string
	}
)

type FailPolicy string

var (
	Break    FailPolicy = "Break"
	Continue FailPolicy = "Continue"
)

func (p *Parallel) activities() []string {
	var result []string
	for _, s := range p.Branches {
		result = append(result, s.activities()...)
	}
	return result
}

func (s *Sequence) activities() []string {
	var result []string
	for _, s := range s.Elements {
		result = append(result, s.activities()...)
	}
	return result
}

func (s *Statement) activities() []string {
	switch {
	case s.Parallel != nil:
		return s.Parallel.activities()
	case s.Sequence != nil:
		return s.Sequence.activities()
	case s.Activity != nil:
		return []string{s.Activity.Name}
	default:
		return nil
	}
}

func (a *ActivityInvocation) execute(ctx workflow.Context, bindings map[string]string, w *Workflow) error {
	input := makeInput(a.Arguments, bindings)
	var result string
	err := workflow.ExecuteActivity(ctx, a.Name, input).Get(ctx, &result)
	if err != nil {
		return err
	}
	if a.Result != "" {
		bindings[a.Result] = result
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	w.activityStatus[a.Name] = true
	return nil
}

func (p *Parallel) execute(ctx workflow.Context, bindings map[string]string, w *Workflow) error {
	childCtx, cancel := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)
	var activityErr error
	for _, s := range p.Branches {
		f := executeAsync(s, childCtx, bindings, w)
		selector.AddFuture(f, func(f workflow.Future) {
			err := f.Get(ctx, nil)
			if err != nil {
				// cancel all pending activities
				cancel()
				activityErr = err
			}
		})
	}

	for i := 0; i < len(p.Branches); i++ {
		// wait for one branch
		selector.Select(ctx)
		if activityErr != nil {
			return activityErr
		}
	}
	return nil
}

func (s *Sequence) execute(ctx workflow.Context, bindings map[string]string, w *Workflow) error {
	for _, a := range s.Elements {
		err := a.execute(ctx, bindings, w)
		if err != nil {
			return err
		}
	}
	return nil
}

// execute
// Paraller、Sequence、Activity 只能有一个
func (s *Statement) execute(ctx workflow.Context, bindings map[string]string, w *Workflow) error {
	switch {
	case s.Parallel != nil:
		return s.Parallel.execute(ctx, bindings, w)
	case s.Sequence != nil:
		return s.Sequence.execute(ctx, bindings, w)
	case s.Activity != nil:
		return s.Activity.execute(ctx, bindings, w)
	default:
		return nil
	}
}

func QueryableDSLWorkflow(ctx workflow.Context, dslWorkflow Workflow) ([]byte, error) {
	status := map[string]bool{}
	for _, a := range dslWorkflow.Root.activities() {
		status[a] = false
	}
	dslWorkflow.activityStatus = status

	// setup query handler for query type "state"
	workflow.SetQueryHandler(ctx, "state", func() (map[string]bool, error) {
		dslWorkflow.lock.RLock()
		defer dslWorkflow.lock.RUnlock()
		return dslWorkflow.activityStatus, nil
	})

	bindings := make(map[string]string)
	for k, v := range dslWorkflow.Variables {
		bindings[k] = v
	}

	ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Minute}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)

	err := dslWorkflow.Root.execute(ctx, bindings, &dslWorkflow)
	if err != nil {
		logger.Error("DLS workflow failed", "Error", err)
	}

	logger.Info("DLS workflow completed")
	return nil, err
}

func makeInput(argNames []string, bindings map[string]string) []string {
	var args []string
	for _, arg := range argNames {
		args = append(args, bindings[arg])
	}
	return args
}

func executeAsync(exe executable, ctx workflow.Context, bindings map[string]string, w *Workflow) workflow.Future {
	future, settable := workflow.NewFuture(ctx)
	workflow.Go(ctx, func(ctx workflow.Context) {
		err := exe.execute(ctx, bindings, w)
		settable.Set(nil, err)
	})
	return future
}
