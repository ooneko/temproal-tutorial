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
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	var workflowID, queryType string
	flag.StringVar(&workflowID, "w", "", "WorkflowID")
	flag.StringVar(&queryType, "t", "state", "Query type [state]")
	flag.Parse()

	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	resp, err := c.QueryWorkflow(context.Background(), workflowID, "", queryType)
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}
	var result map[string]bool
	if err := resp.Get(&result); err != nil {
		log.Fatalln("Unable to decode query result", err)
	}
	log.Printf("Workflow %s status %v", workflowID, result)
}
