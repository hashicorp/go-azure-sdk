
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresourceoperationgroup` Documentation

The `remoteprivateendpointconnectionarmresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresourceoperationgroup"
```


### Client Initialization

```go
client := remoteprivateendpointconnectionarmresourceoperationgroup.NewRemotePrivateEndpointConnectionARMResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourceOperationGroupClient.WebAppsApproveOrRejectPrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresourceoperationgroup.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "privateEndpointConnectionName")

payload := remoteprivateendpointconnectionarmresourceoperationgroup.RemotePrivateEndpointConnectionARMResource{
	// ...
}


if err := client.WebAppsApproveOrRejectPrivateEndpointConnectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourceOperationGroupClient.WebAppsDeletePrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresourceoperationgroup.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "privateEndpointConnectionName")

if err := client.WebAppsDeletePrivateEndpointConnectionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourceOperationGroupClient.WebAppsGetPrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresourceoperationgroup.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "privateEndpointConnectionName")

read, err := client.WebAppsGetPrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourceOperationGroupClient.WebAppsGetPrivateEndpointConnectionList`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsGetPrivateEndpointConnectionList(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsGetPrivateEndpointConnectionListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
