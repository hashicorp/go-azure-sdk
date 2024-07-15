
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "privateEndpointConnectionValue")

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
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "privateEndpointConnectionValue")

if err := client.ManagedEnvironmentPrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ManagedEnvironmentPrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "privateEndpointConnectionValue")

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
id := privateendpointconnections.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue")

// alternatively `client.ManagedEnvironmentPrivateEndpointConnectionsList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentPrivateEndpointConnectionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
