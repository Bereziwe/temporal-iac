package main

import (
    "go.temporal.io/sdk/workflow"
    "time"
)

func TerraformWorkflow(ctx workflow.Context, config string) (string, error) {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    var result string
    err := workflow.ExecuteActivity(ctx, TerraformApply, config).Get(ctx, &result)
    if err != nil {
        return "", err
    }

    return result, nil
}


