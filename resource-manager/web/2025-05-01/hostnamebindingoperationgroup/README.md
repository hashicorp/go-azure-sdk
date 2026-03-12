
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindingoperationgroup` Documentation

The `hostnamebindingoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindingoperationgroup"
```


### Client Initialization

```go
client := hostnamebindingoperationgroup.NewHostNameBindingOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HostNameBindingOperationGroupClient.WebAppsCreateOrUpdateHostNameBindingSlot`

```go
ctx := context.TODO()
id := hostnamebindingoperationgroup.NewSlotHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hostNameBindingName")

payload := hostnamebindingoperationgroup.HostNameBinding{
	// ...
}


read, err := client.WebAppsCreateOrUpdateHostNameBindingSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingOperationGroupClient.WebAppsDeleteHostNameBindingSlot`

```go
ctx := context.TODO()
id := hostnamebindingoperationgroup.NewSlotHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hostNameBindingName")

read, err := client.WebAppsDeleteHostNameBindingSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingOperationGroupClient.WebAppsGetHostNameBindingSlot`

```go
ctx := context.TODO()
id := hostnamebindingoperationgroup.NewSlotHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "hostNameBindingName")

read, err := client.WebAppsGetHostNameBindingSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingOperationGroupClient.WebAppsListHostNameBindingsSlot`

```go
ctx := context.TODO()
id := hostnamebindingoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListHostNameBindingsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListHostNameBindingsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
