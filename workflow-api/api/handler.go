package api

import (
	"context"
	"log"

	"github.com/emicklei/go-restful"
	"github.com/ooneko/temporal-tutorial/helloworld"
	"go.temporal.io/sdk/client"
)

type runResult struct {
	Message    string
	WorkflowID string
	RunID      string
}

type handler struct {
	client client.Client
}

func (h *handler) run(request *restful.Request, response *restful.Response) {
	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world_workflowID",
		TaskQueue: "hello-world",
	}

	we, err := h.client.ExecuteWorkflow(context.Background(), workflowOptions, helloworld.Workflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
	response.WriteHeader(200)
	response.WriteEntity(runResult{
		Message:    "Workflow started",
		WorkflowID: we.GetID(),
		RunID:      we.GetRunID(),
	})
}

func (h *handler) result(request *restful.Request, response *restful.Response) {
	workflowID := request.QueryParameter("workflow_id")
	runID := request.QueryParameter("run_id")
	we := h.client.GetWorkflow(context.Background(), workflowID, runID)

	// Synchronously wait for the workflow completion.
	var result string
	err := we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
	response.Write([]byte("test"))
}

func (h *handler) progress(request *restful.Request, response *restful.Response) {

}

func (h *handler) cancel(request *restful.Request, response *restful.Response) {
}

func (h *handler) terminate(request *restful.Request, response *restful.Response) {
}

func (h *handler) retryActivity(request *restful.Request, response *restful.Response) {}

func (h *handler) pause(request *restful.Request, response *restful.Response) {}
