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

package main

import (
	"context"
	"flag"
	dsl_process_query "github.com/ooneko/temporal-tutorial/dsl-process-query"
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	var dslConfig string
	flag.StringVar(&dslConfig, "dslConfig", "workflow1.yaml", "dslConfig specify the yaml file for dsl workflow")
	flag.Parse()

	data, err := os.ReadFile(dslConfig)
	if err != nil {
		log.Fatalln("failed to load dsl config file: ", err)
	}
	var dslWorkflow dsl_process_query.Workflow
	if err := yaml.Unmarshal(data, &dslWorkflow); err != nil {
		log.Fatalln("failed to unmarshal dsl workflow: ", err)
	}

	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client: ", err)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        "dsl_" + uuid.New(),
		TaskQueue: "dsl",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, dsl_process_query.QueryableDSLWorkflow, dslWorkflow)
	if err != nil {
		log.Fatalln("Unable to exceed workflow", err)
	}
	log.Println("Started workflow", "Workflow ID", we.GetID(), "RunID", we.GetRunID())
}
