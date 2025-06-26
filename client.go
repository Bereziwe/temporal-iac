package main

import (
    "context"
    "log"
    "go.temporal.io/sdk/client"
)

func main() {
    c, err := client.NewClient(client.Options{
        HostPort: "172.172.245.99:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    workflowOptions := client.StartWorkflowOptions{
        TaskQueue: "terraform-task",
    }

    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, DeployTerraformWorkflow, "example-config")
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
