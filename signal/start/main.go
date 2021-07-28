package main

import (
	"context"
	"log"

	"github.com/ooneko/temporal-tutorial/signal"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client, err")
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        signal.WorkflowID,
		TaskQueue: signal.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(),
		workflowOptions, signal.Workflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
