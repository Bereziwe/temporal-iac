// worker/main.go
package main


import (
"github.com/Bereziwe/temporal-iac/workflows"
"log"
"temporal-iac/activities"
"temporal-iac/workflows"
"go.temporal.io/sdk/client"
"go.temporal.io/sdk/worker"
)

func main() {
c, err := client.Dial(client.Options{
HostPort: "172.172.245.99:7233", // your Temporal server
})
if err != nil {
log.Fatalln("unable to create client", err)
}
defer c.Close()

w := worker.New(c, "terraform-task-queue", worker.Options{})
w.RegisterWorkflow(workflows.DeployTerraformWorkflow)
w.RegisterActivity(activities.InitTerraform)
w.RegisterActivity(activities.PlanTerraform)
w.RegisterActivity(activities.ApplyTerraform)

err = w.Run(worker.InterruptCh())
if err != nil {
log.Fatalln("unable to start worker", err)
}
}
