
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/global` Documentation

The `global` SDK allows for interaction with Azure Resource Manager `web` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/global"
```


### Client Initialization

```go
client := global.NewGlobalClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GlobalClient.GetDeletedWebApp`

```go
ctx := context.TODO()
id := global.NewDeletedSiteID("12345678-1234-9876-4563-123456789012", "deletedSiteId")

read, err := client.GetDeletedWebApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GlobalClient.GetDeletedWebAppSnapshots`

```go
ctx := context.TODO()
id := global.NewDeletedSiteID("12345678-1234-9876-4563-123456789012", "deletedSiteId")

read, err := client.GetDeletedWebAppSnapshots(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
