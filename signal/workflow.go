package signal

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, _ string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started")

	err := workflow.SetQueryHandler(ctx, "getItem", func(input []byte) (string, error) {
		return "", nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	channel := workflow.GetSignalChannel(ctx, SignalChannelName)

	var signal string
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(channel, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &signal)
	})

	selector.Select(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err = workflow.ExecuteActivity(ctx, Activity, signal).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return err
	}
	logger.Info("HelloWorld workflow completed.", "result", result)
	return nil
}
