package main

import (
    "go.temporal.io/sdk/client"
    "log"
)

func main() {
    c, err := client.NewClient(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    workflowOptions := client.StartWorkflowOptions{
        TaskQueue: "terraform-task2",
    }

    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, YourWorkflowFunctionName)
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "terrafformworke", we.GetID(), "56674338977", we.GetRunID())
}


