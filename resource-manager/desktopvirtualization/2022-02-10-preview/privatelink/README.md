
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/privatelink` Documentation

The `privatelink` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2022-02-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/privatelink"
```


### Client Initialization

```go
client := privatelink.NewPrivateLinkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsDeleteByHostPool`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsDeleteByHostPool(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsDeleteByWorkspace`

```go
ctx := context.TODO()
id := privatelink.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsDeleteByWorkspace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGetByHostPool`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsGetByHostPool(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGetByWorkspace`

```go
ctx := context.TODO()
id := privatelink.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsGetByWorkspace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsListByHostPool`

```go
ctx := context.TODO()
id := privatelink.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue")

// alternatively `client.PrivateEndpointConnectionsListByHostPool(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListByHostPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsListByWorkspace`

```go
ctx := context.TODO()
id := privatelink.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.PrivateEndpointConnectionsListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsUpdateByHostPool`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue", "privateEndpointConnectionValue")

payload := privatelink.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsUpdateByHostPool(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsUpdateByWorkspace`

```go
ctx := context.TODO()
id := privatelink.NewWorkspacePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "privateEndpointConnectionValue")

payload := privatelink.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsUpdateByWorkspace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.ResourcesListByHostPool`

```go
ctx := context.TODO()
id := privatelink.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolValue")

// alternatively `client.ResourcesListByHostPool(ctx, id)` can be used to do batched pagination
items, err := client.ResourcesListByHostPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkClient.ResourcesListByWorkspace`

```go
ctx := context.TODO()
id := privatelink.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.ResourcesListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ResourcesListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
