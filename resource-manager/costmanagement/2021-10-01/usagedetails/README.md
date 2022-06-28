
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/usagedetails` Documentation

The `usagedetails` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/usagedetails"
```


### Client Initialization

```go
client := usagedetails.NewUsageDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsageDetailsClient.GenerateDetailedCostReportCreateOperation`

```go
ctx := context.TODO()
id := usagedetails.NewScopeID()

payload := usagedetails.GenerateDetailedCostReportDefinition{
	// ...
}


if err := client.GenerateDetailedCostReportCreateOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
