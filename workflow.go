
func RunTerraformApply(ctx context.Context, dir string) error {
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Dir = dir
    return cmd.Run()
}
