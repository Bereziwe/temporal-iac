// package main

// import (
//     // "context"
//     // "fmt"
//     "os/exec"
// )

package main

import (
    "context"
    "log"
    "fmt"
    "go.temporal.io/sdk/client"
)

func main() {
    c, err := client.NewClient(client.Options{
        HostPort: "172.172.245.99:7233",
    })
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    workflowOptions := client.StartWorkflowOptions{
        TaskQueue: "terraform-task11", // must match the worker's task queue
    }

    we, err := c.ExecuteWorkflow(
        context.Background(),
        workflowOptions,
        TerraformWorkflow, // replace with your actual workflow function
        "./terraform", // replace with actual argument if needed
    )
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "WorkflowIDtest", we.GetID(), "RunID667877", we.GetRunID())
}

// func TerraformApply() error {
// cmd := exec.Command("terraform", "apply", "-auto-approve", "tfplan")
// cmd.Dir = "./terraform"
// return cmd.Run()
// }


func TerraformApply(ctx context.Context, config string) (string, error) {
    fmt.Println("Executing Terraform with config:", config)
    // Simulate Terraform logic or call out to scripts/binaries
    return "Terraform apply complete for config: " + config, nil
}
