package main

import (
    "log"

    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
    "github.com/Bereziwe/temporal-iac/workflows" // Adjust this import path
)

func main() {
    c, err := client.Dial(client.Options{
        HostPort: "172.172.245.99:7233", // or :8080 if that's your SDK port
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    w := worker.New(c, "terraform-task-queue", worker.Options{})

    w.RegisterWorkflow(workflows.DeployWorkflow)
    w.RegisterActivity(workflows.RunTerraformActivity)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}











// // client/start.go
// package main

// import (
//     "context"
//     "log"
//     "go.temporal.io/sdk/client"
// 	"go.temporal.io/sdk/worker"
//     "github.com/Bereziwe/temporal-iac/workflows"
// )

// // package main
// // import "temporal-iac/activities"
// // import "github.com/Bereziwe/temporal-iac/activities"
// // import (
// // "temporal-iac/activities"
// // "log"
// // // "temporal-iac/activities"
// // "temporal-iac/workflows"
// // "go.temporal.io/sdk/client"
// // )

// func main() {
// c, err := client.Dial(client.Options{
// HostPort: "172.172.245.99:7233",
// })
// if err != nil {
// log.Fatalln("unable to create client", err)
// }
// defer c.Close()

// workflowOptions := client.StartWorkflowOptions{
// ID:        "terraform-deploy-workflow",
// TaskQueue: "terraform-task-queue",
// }

// _, err = c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.DeployTerraformWorkflow)
// if err != nil {
// log.Fatalln("unable to execute workflow", err)
//    }
// }
