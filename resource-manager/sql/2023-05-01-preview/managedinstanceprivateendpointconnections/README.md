
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/managedinstanceprivateendpointconnections` Documentation

The `managedinstanceprivateendpointconnections` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/managedinstanceprivateendpointconnections"
```


### Client Initialization

```go
client := managedinstanceprivateendpointconnections.NewManagedInstancePrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstancePrivateEndpointConnectionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedinstanceprivateendpointconnections.NewManagedInstancePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "privateEndpointConnectionValue")

payload := managedinstanceprivateendpointconnections.ManagedInstancePrivateEndpointConnection{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedInstancePrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := managedinstanceprivateendpointconnections.NewManagedInstancePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "privateEndpointConnectionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedInstancePrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := managedinstanceprivateendpointconnections.NewManagedInstancePrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "privateEndpointConnectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedInstancePrivateEndpointConnectionsClient.ListByManagedInstance`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByManagedInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
