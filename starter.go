package main

import (
    "context"
    "log"
    // "time"
    "go.temporal.io/sdk/client"
    // "your/module/path/workflows" // Replace with the actual module path
)

func main() {
    // Create Temporal client
    c, err := client.NewClient(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    // Define workflow options
    workflowOptions := client.StartWorkflowOptions{
        ID:        "my_workflow_id",
        TaskQueue: "terraform-task2",
    }

    // Start the workflow
    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, MyWorkflow, "Hello from starter.go!")
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "WorkflowID:", we.GetID(), "RunID:", we.GetRunID())
}
