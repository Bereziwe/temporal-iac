// // worker/main.go
// package main


// import (
// "github.com/Bereziwe/temporal-iac/workflows"
// "github.com/Bereziwe/temporal-iac/worker"
// "log"
// "go.temporal.io/sdk/client"
// "go.temporal.io/sdk/worker"
// )

// func main() {
// c, err := client.Dial(client.Options{
// 	    Namespace: "default"
//         HostPort: "172.172.245.99:7233", // your Temporal server
// })
// if err != nil {
// log.Fatalln("unable to create client", err)
// }
// defer c.Close()

// w := worker.New(c, "terraform-task", worker.Options{})
// w.RegisterActivity(activities.RunTerraformActivity)
// w.RegisterWorkflow(workflows.DeployTerraformWorkflow)
// w.RegisterActivity(activities.InitTerraform)
// w.RegisterActivity(activities.PlanTerraform)
// w.RegisterActivity(activities.ApplyTerraform)

// err = w.Run(worker.InterruptCh())
// if err != nil {
// log.Fatalln("unable to start worker", err)
// }
// }
