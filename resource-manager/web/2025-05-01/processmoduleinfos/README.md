
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfos` Documentation

The `processmoduleinfos` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/processmoduleinfos"
```


### Client Initialization

```go
client := processmoduleinfos.NewProcessModuleInfosClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProcessModuleInfosClient.WebAppsGetInstanceProcessModule`

```go
ctx := context.TODO()
id := processmoduleinfos.NewInstanceProcessModuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "instanceId", "processId", "moduleName")

read, err := client.WebAppsGetInstanceProcessModule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProcessModuleInfosClient.WebAppsListInstanceProcessModules`

```go
ctx := context.TODO()
id := processmoduleinfos.NewInstanceProcessID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "instanceId", "processId")

// alternatively `client.WebAppsListInstanceProcessModules(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceProcessModulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
