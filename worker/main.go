



package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/Bereziwe/temporal-iac/workflows"
)

func main() {
	serviceClient, err := client.NewClient(client.Options{
		Namespace: "default",
		HostPort:  "172.172.245.99:7233",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	temporalWorker := worker.New(serviceClient, "temporal-terraform-demo", worker.Options{
		WorkerStopTimeout: 30 * time.Second,
	})

	log.Print("registering workflows")
	workflows.Register(temporalWorker)

	if err := temporalWorker.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}














// // worker/main.go
// package main


// import (
// "github.com/Bereziwe/temporal-iac/workflows"
// "github.com/Bereziwe/temporal-iac/worker"
// "log"
// "temporal-iac/activities"
// "temporal-iac/workflows"
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
// temporalWorker := worker.new(serviceClient, "terraform-task", worker.Options{
// 	WorkerStopTimeout: 30 * time.Second,
// })
// w := worker.New(c, "terraform-task", worker.Options{})
// w.RegisterActivity(workflows.RunTerraformActivity)
// w.RegisterWorkflow(workflows.DeployTerraformWorkflow)
// w.RegisterActivity(activities.InitTerraform)
// w.RegisterActivity(activities.PlanTerraform)
// w.RegisterActivity(activities.ApplyTerraform)

// err = w.Run(worker.InterruptCh())
// if err != nil {
// log.Fatalln("unable to start worker", err)
// }
// }
