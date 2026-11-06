
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.ContainerAppPrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


if err := client.ContainerAppPrivateEndpointConnectionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ContainerAppPrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "privateEndpointConnectionName")

if err := client.ContainerAppPrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ContainerAppPrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "privateEndpointConnectionName")

read, err := client.ContainerAppPrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ContainerAppPrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := privateendpointconnections.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ContainerAppPrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.ContainerAppPrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewManagedEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


if err := client.ManagedEnvironmentPrivateEndpointConnectionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewManagedEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "privateEndpointConnectionName")

if err := client.ManagedEnvironmentPrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewManagedEnvironmentPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "privateEndpointConnectionName")

read, err := client.ManagedEnvironmentPrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := privateendpointconnections.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName")

// alternatively `client.ManagedEnvironmentPrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentPrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
