
package activities

import (
    "context"
    "os/exec"
)

func TerraformInitApply(ctx context.Context, dir string) error {
    cmds := [][]string{
        {"terraform", "init"},
        {"terraform", "apply", "-auto-approve"},
    }

   for _, args := range cmds {
	    cmd := exec.CommandContext(ctx, args[0], args[1:]...)
        cmd.Dir = dir
        output, err := cmd.CombinedOutput()
        if err != nil {
            return fmt.Errorf("error running %v: %v\nOutput: %s", args, err, output)
      }
    }

    return nil
}
