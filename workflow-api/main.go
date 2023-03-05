package main

import (
	"github.com/ooneko/temporal-tutorial/workflow-api/api"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	apiserver := api.New()
	apiserver.Prepare(signals.SetupSignalHandler())
	apiserver.Run()
}
