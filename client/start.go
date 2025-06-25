// client/start.go
package main

import (
    "context"
    "log"
    "go.temporal.io/sdk/client"
	// "go.temporal.io/sdk/worker"
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
HostPort: "172.172.245.99:7233",
})
if err != nil {
log.Fatalln("unable to create client", err)
}
defer c.Close()

workflowOptions := client.StartWorkflowOptions{
ID:        "terraform-deploy4",
TaskQueue: "terraform-task",
}

_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.DeployTerraformWorkflow)
if err != nil {
log.Fatalln("unable to execute workflow", err)
   }
}
