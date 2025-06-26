module github.com/Bereziwe/temporal-iac

go 1.21

require (
    go.temporal.io/sdk v1.24.0
    github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.1.0
    github.com/hashicorp/terraform-exec v0.23.0
)

replace github.com/Bereziwe/temporal-iac => ./
