// // activities/terraform.go
// package activities

// import (
// "os/exec"
// )

// func InitTerraform() error {
// cmd := exec.Command("terraform", "init")
// cmd.Dir = "./terraform"
// return cmd.Run()
// }

// func PlanTerraform() error {
// cmd := exec.Command("terraform", "plan", "-out=tfplan")
// cmd.Dir = "./terraform"
// return cmd.Run()
// }

// func ApplyTerraform() error {
// cmd := exec.Command("terraform", "apply", "-auto-approve", "tfplan")
// cmd.Dir = "./terraform"
// return cmd.Run()
// }
