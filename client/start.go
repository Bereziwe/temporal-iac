// client/start.go
package main

"github.com/yourusername/temporal-terraform/workflows"

import (
"github.com/yourusername/temporal-terraform/workflows"
"log"
"your_project/workflows"
"go.temporal.io/sdk/client"
)

func main() {
c, err := client.Dial(client.Options{
HostPort: "localhost:7233",
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
