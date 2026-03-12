
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresources` Documentation

The `remoteprivateendpointconnectionarmresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/remoteprivateendpointconnectionarmresources"
```


### Client Initialization

```go
client := remoteprivateendpointconnectionarmresources.NewRemotePrivateEndpointConnectionARMResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourcesClient.AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresources.NewHostingEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "privateEndpointConnectionName")

payload := remoteprivateendpointconnectionarmresources.RemotePrivateEndpointConnectionARMResource{
	// ...
}


if err := client.AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourcesClient.AppServiceEnvironmentsDeletePrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresources.NewHostingEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "privateEndpointConnectionName")

if err := client.AppServiceEnvironmentsDeletePrivateEndpointConnectionThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourcesClient.AppServiceEnvironmentsGetPrivateEndpointConnection`

```go
ctx := context.TODO()
id := remoteprivateendpointconnectionarmresources.NewHostingEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName", "privateEndpointConnectionName")

read, err := client.AppServiceEnvironmentsGetPrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RemotePrivateEndpointConnectionARMResourcesClient.AppServiceEnvironmentsGetPrivateEndpointConnectionList`

```go
ctx := context.TODO()
id := commonids.NewAppServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentName")

// alternatively `client.AppServiceEnvironmentsGetPrivateEndpointConnectionList(ctx, id)` can be used to do batched pagination
items, err := client.AppServiceEnvironmentsGetPrivateEndpointConnectionListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
