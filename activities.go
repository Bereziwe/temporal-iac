package main

import (
    "context"
    "fmt"
    "log"
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
	// "os/exec"
)

func main() {
    c, err := client.NewClient(client.Options{
		HostPort: "172.172.245.99:7233",
	})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    w := worker.New(c, "terraform-task2", worker.Options{})

    // w.RegisterWorkflow(TerraformWorkflow)
    w.RegisterActivity(DeployTerraformApply)
	// w.RegisterActivity(TerraformInit)

    log.Println("Worker started for task queue: terraform-task2")
    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}


// func TerraformInit(ctx context.Context) error {
//     cmd := exec.Command("terraform", "init")
// 	cmd.Dir = "/home/runner/work/temporal-iac/temporal-iac/terraform"
//     // cmd.Dir = "./terraform" // Replace with your actual path
//     output, err := cmd.CombinedOutput()
//     if err != nil {
//         return fmt.Errorf("terraform init failed: %v\nOutput: %s", err, string(output))
//     }

//     return nil
// }

func DeployTerraformApply(ctx context.Context, config string) (string, error) {
    fmt.Println("Executing Terraform with config:", config)
    // Simulate Terraform logic or call out to scripts/binaries
    return "Terraform apply complete for config: " + config, nil
}
