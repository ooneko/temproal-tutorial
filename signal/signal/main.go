package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/ooneko/temporal-tutorial/signal"
)

func main() {

	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	err = c.SignalWorkflow(context.Background(), signal.WorkflowID, "", signal.SignalChannelName, "Temporal")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
}
