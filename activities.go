package main

import (
    "context"
    "fmt"
)

func TerraformApply() error {
cmd := exec.Command("terraform", "apply", "-auto-approve", "tfplan")
cmd.Dir = "./terraform"
return cmd.Run()
}


// func TerraformApply(ctx context.Context, config string) (string, error) {
//     fmt.Println("Executing Terraform with config:", config)
//     // Simulate Terraform logic or call out to scripts/binaries
//     return "Terraform apply complete for config: " + config, nil
// }
