
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfooperationgroup` Documentation

The `processmoduleinfooperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfooperationgroup"
```


### Client Initialization

```go
client := processmoduleinfooperationgroup.NewProcessModuleInfoOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProcessModuleInfoOperationGroupClient.WebAppsGetProcessModule`

```go
ctx := context.TODO()
id := processmoduleinfooperationgroup.NewModuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "processId", "moduleName")

read, err := client.WebAppsGetProcessModule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessModuleInfoOperationGroupClient.WebAppsListProcessModules`

```go
ctx := context.TODO()
id := processmoduleinfooperationgroup.NewProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "processId")

// alternatively `client.WebAppsListProcessModules(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListProcessModulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
