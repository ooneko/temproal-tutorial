package dsl_process_error

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
)

type SampleActivities struct {
}

// 这里模拟了每个任务都有50%的概率失败，一次性跑完所有任务几乎是不可能的。

func (a *SampleActivities) SampleActivity1(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 5 {
		return "Activity1 failed", fmt.Errorf("activity1 failed")
	}
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity2(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 5 {
		return "Activity1 failed", fmt.Errorf("activity1 failed")
	}
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity3(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 5 {
		return "Activity1 failed", fmt.Errorf("activity1 failed")
	}
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity4(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 5 {
		return "Activity1 failed", fmt.Errorf("activity1 failed")
	}
	return "Result_" + name, nil
}

func (a *SampleActivities) SampleActivity5(ctx context.Context, input []string) (string, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	fmt.Printf("Run %s with input %v \n", name, input)
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 5 {
		return "Activity1 failed", fmt.Errorf("activity1 failed")
	}
	return "Result_" + name, nil
}
