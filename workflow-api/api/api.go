package api

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"

	"go.temporal.io/sdk/client"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *APIServer {
	return &APIServer{Server: &http.Server{Addr: ":8080"}}
}

type APIServer struct {
	Server *http.Server
	// webservice container, where all webservice defines
	container      *restful.Container
	temproalClient client.Client
}

func (s *APIServer) Prepare(ctx context.Context) {
	s.container = restful.NewContainer()

	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	s.temproalClient = c
	go func() {
		<-ctx.Done()
		c.Close()
	}()

	s.InstallAPIS()
	s.Server.Handler = s.container
}

func (s *APIServer) InstallAPIS() {
	must(AddToContainer(s.container, s.temproalClient))
}

func (s *APIServer) Run() error {
	fmt.Printf("Start listening on %s", s.Server.Addr)
	return s.Server.ListenAndServe()
}
