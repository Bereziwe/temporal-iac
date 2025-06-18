
package workflows

import (
    "context"
    "go.temporal.io/sdk/workflow"
    "your_project/activities"
)

func DeployAzureInfraWorkflow(ctx workflow.Context, terraformDir string) error {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute * 10,
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    return workflow.ExecuteActivity(ctx, activities.TerraformInitApply, terraformDir).Get(ctx, nil)
}
