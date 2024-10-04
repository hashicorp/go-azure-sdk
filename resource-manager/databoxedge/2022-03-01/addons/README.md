
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/addons` Documentation

The `addons` SDK allows for interaction with Azure Resource Manager `databoxedge` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/addons"
```


### Client Initialization

```go
client := addons.NewAddonsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AddonsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := addons.NewAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "roleName", "addonName")

payload := addons.Addon{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AddonsClient.Delete`

```go
ctx := context.TODO()
id := addons.NewAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "roleName", "addonName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AddonsClient.Get`

```go
ctx := context.TODO()
id := addons.NewAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "roleName", "addonName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AddonsClient.ListByRole`

```go
ctx := context.TODO()
id := addons.NewRoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deviceName", "name")

// alternatively `client.ListByRole(ctx, id)` can be used to do batched pagination
items, err := client.ListByRoleComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
