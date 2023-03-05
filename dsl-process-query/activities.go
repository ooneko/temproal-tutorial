package dsl_process_query

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
)

type SampleActivities struct {
}

func (a *SampleActivities) SampleActivity1(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity2(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity3(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity4(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity5(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivityWithError(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	time.Sleep(time.Second * 60)
	return "", fmt.Errorf("something wrong happened")
}
