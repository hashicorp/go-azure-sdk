
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/privateendpointconnectionslotoperationgroup` Documentation

The `privateendpointconnectionslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/privateendpointconnectionslotoperationgroup"
```


### Client Initialization

```go
client := privateendpointconnectionslotoperationgroup.NewPrivateEndpointConnectionSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionSlotOperationGroupClient.WebAppsApproveOrRejectPrivateEndpointConnectionSlot`

```go
ctx := context.TODO()
id := privateendpointconnectionslotoperationgroup.NewSlotPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "privateEndpointConnectionName")

payload := privateendpointconnectionslotoperationgroup.RemotePrivateEndpointConnectionARMResource{
	// ...
}


if err := client.WebAppsApproveOrRejectPrivateEndpointConnectionSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionSlotOperationGroupClient.WebAppsDeletePrivateEndpointConnectionSlot`

```go
ctx := context.TODO()
id := privateendpointconnectionslotoperationgroup.NewSlotPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "privateEndpointConnectionName")

if err := client.WebAppsDeletePrivateEndpointConnectionSlotThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionSlotOperationGroupClient.WebAppsGetPrivateEndpointConnectionListSlot`

```go
ctx := context.TODO()
id := privateendpointconnectionslotoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsGetPrivateEndpointConnectionListSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsGetPrivateEndpointConnectionListSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionSlotOperationGroupClient.WebAppsGetPrivateEndpointConnectionSlot`

```go
ctx := context.TODO()
id := privateendpointconnectionslotoperationgroup.NewSlotPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "privateEndpointConnectionName")

read, err := client.WebAppsGetPrivateEndpointConnectionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
