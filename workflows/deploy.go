// workflows/deploy.go
package workflows

import (
"github.com/Bereziwe/temporal-iac/workflows"
"go.temporal.io/sdk/workflow"
"temporal-iac/activities"
"time"
)

func DeployTerraformWorkflow(ctx workflow.Context) error {
opts := workflow.ActivityOptions{
StartToCloseTimeout: time.Minute * 10,
}
ctx = workflow.WithActivityOptions(ctx, opts)

err := workflow.ExecuteActivity(ctx, activities.InitTerraform).Get(ctx, nil)
if err != nil {
return err
}
err = workflow.ExecuteActivity(ctx, activities.PlanTerraform).Get(ctx, nil)
if err != nil {
return err
}
return workflow.ExecuteActivity(ctx, activities.ApplyTerraform).Get(ctx, nil)
}
