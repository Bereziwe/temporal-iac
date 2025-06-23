// client/start.go
package main

import (
    "context"
    "log"

    "go.temporal.io/sdk/client"
    "github.com/Bereziwe/temporal-iac/workflows"
)

// package main
// import "temporal-iac/activities"
// import "github.com/Bereziwe/temporal-iac/activities"
// import (
// "temporal-iac/activities"
// "log"
// // "temporal-iac/activities"
// "temporal-iac/workflows"
// "go.temporal.io/sdk/client"
// )

func main() {
c, err := client.Dial(client.Options{
HostPort: "20.106.237.42:8080/",
})
if err != nil {
log.Fatalln("unable to create client", err)
}
defer c.Close()

workflowOptions := client.StartWorkflowOptions{
ID:        "terraform-deploy-workflow",
TaskQueue: "terraform-task-queue",
}

_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.DeployTerraformWorkflow)
if err != nil {
log.Fatalln("unable to execute workflow", err)
}
}
