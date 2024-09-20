
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-07-01-preview/usagedetails` Documentation

The `usagedetails` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2023-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-07-01-preview/usagedetails"
```


### Client Initialization

```go
client := usagedetails.NewUsageDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsageDetailsClient.GenerateDetailedCostReportCreateOperation`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := usagedetails.GenerateDetailedCostReportDefinition{
	// ...
}


if err := client.GenerateDetailedCostReportCreateOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
