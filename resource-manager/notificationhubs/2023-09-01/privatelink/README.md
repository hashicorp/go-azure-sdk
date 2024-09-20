
## `github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/privatelink` Documentation

The `privatelink` SDK allows for interaction with Azure Resource Manager `notificationhubs` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/privatelink"
```


### Client Initialization

```go
client := privatelink.NewPrivateLinkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

if err := client.PrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGetGroupId`

```go
ctx := context.TODO()
id := privatelink.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "subResourceName")

read, err := client.PrivateEndpointConnectionsGetGroupId(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := privatelink.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.PrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsListGroupIds`

```go
ctx := context.TODO()
id := privatelink.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.PrivateEndpointConnectionsListGroupIds(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListGroupIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsUpdate`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "privateEndpointConnectionName")

payload := privatelink.PrivateEndpointConnectionResource{
	// ...
}


if err := client.PrivateEndpointConnectionsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
