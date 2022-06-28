
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/startmenuitem` Documentation

The `startmenuitem` SDK allows for interaction with the Azure Resource Manager Service `desktopvirtualization` (API Version `2022-02-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/startmenuitem"
```


### Client Initialization

```go
client := startmenuitem.NewStartMenuItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `StartMenuItemClient.List`

```go
ctx := context.TODO()
id := startmenuitem.NewApplicationGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "applicationGroupValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
