
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2023-09-01/availableservicetiers` Documentation

The `availableservicetiers` SDK allows for interaction with Azure Resource Manager `operationalinsights` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2023-09-01/availableservicetiers"
```


### Client Initialization

```go
client := availableservicetiers.NewAvailableServiceTiersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableServiceTiersClient.ListByWorkspace`

```go
ctx := context.TODO()
id := availableservicetiers.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.ListByWorkspace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
