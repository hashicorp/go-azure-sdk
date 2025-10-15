
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/mhsmprivateendpointconnections` Documentation

The `mhsmprivateendpointconnections` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/mhsmprivateendpointconnections"
```


### Client Initialization

```go
client := mhsmprivateendpointconnections.NewMhsmPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MhsmPrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "privateEndpointConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MhsmPrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "privateEndpointConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MhsmPrivateEndpointConnectionsClient.ListByResource`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewManagedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName")

// alternatively `client.ListByResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MhsmPrivateEndpointConnectionsClient.Put`

```go
ctx := context.TODO()
id := mhsmprivateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMName", "privateEndpointConnectionName")

payload := mhsmprivateendpointconnections.MHSMPrivateEndpointConnection{
	// ...
}


read, err := client.Put(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
