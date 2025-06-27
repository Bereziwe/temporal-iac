package main

import (
    "log"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
)

func main() {
    c, err := client.NewClient(client.Options{
        HostPort: "172.172.245.99:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    w := worker.New(c, "terraform-task11", worker.Options{})

    w.RegisterWorkflow(TerraformWorkflow)
    w.RegisterActivity(TerraformApply)

    log.Println("Worker started for task queue: terraform-task")
    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}

