package main

import (
    "context"
    "go.temporal.io/sdk/worker"
    "go.temporal.io/sdk/client"
    "log"
)

func main() {
    c, err := client.NewClient(client.Options{
        HostPort: "172.172.245.99:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    w := worker.New(c, "terraform-task10", worker.Options{})

    // w.RegisterWorkflow(TerraformWorkflow)
    w.RegisterActivity(TerraformApply)
	// w.RegisterActivity(TerraformInit)

    workflowOptions := client.StartWorkflowOptions{
        TaskQueue: "terraform-task10",
    }

    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, TerraformApply, "/home/runner/work/temporal-iac/temporal-iac/terraform")
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "terrafformworke", we.GetID(), "56674338977", we.GetRunID())
}


