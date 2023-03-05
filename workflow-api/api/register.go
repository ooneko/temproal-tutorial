package api

import (
	"github.com/emicklei/go-restful"
	"go.temporal.io/sdk/client"
	"net/http"
)

func AddToContainer(c *restful.Container, client client.Client) error {
	h := handler{client: client}
	webservice := NewWebService("helloworld")
	webservice.Route(webservice.GET("/run").
		To(h.run).
		Doc("Run Hello World workflow.").
		Returns(http.StatusOK, "OK", runResult{}))

	webservice.Route(webservice.GET("status").
		To(h.result).
		Doc("Run Hello World workflow.").
		Param(webservice.PathParameter("workflow_id", "workflow id")).
		Param(webservice.PathParameter("run_id", "run id")).
		Returns(http.StatusOK, "OK", nil))

	c.Add(webservice)
	return nil
}

func NewWebService(path string) *restful.WebService {
	webservice := restful.WebService{}
	webservice.Path("/workflow" + "/" + path).
		Produces(restful.MIME_JSON)
	return &webservice
}
