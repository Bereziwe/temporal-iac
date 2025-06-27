package main

import (
    "go.temporal.io/sdk/workflow"
    "fmt"
	// "os/exec"
)

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

func TerraformApply(ctx workflow.Context, config string) (string, error) {
    fmt.Println("Executing Terraform with config:", config)
    // Simulate Terraform logic or call out to scripts/binaries
    return "Terraform apply complete for config: " + config, nil
}
