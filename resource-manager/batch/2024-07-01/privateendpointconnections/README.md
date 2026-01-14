
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionDelete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "privateEndpointConnectionName")

if err := client.PrivateEndpointConnectionDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionGet`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "privateEndpointConnectionName")

read, err := client.PrivateEndpointConnectionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionListByBatchAccount`

```go
ctx := context.TODO()
id := privateendpointconnections.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName")

// alternatively `client.PrivateEndpointConnectionListByBatchAccount(ctx, id, privateendpointconnections.DefaultPrivateEndpointConnectionListByBatchAccountOperationOptions())` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionListByBatchAccountComplete(ctx, id, privateendpointconnections.DefaultPrivateEndpointConnectionListByBatchAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionsClient.PrivateEndpointConnectionUpdate`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "privateEndpointConnectionName")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


if err := client.PrivateEndpointConnectionUpdateThenPoll(ctx, id, payload, privateendpointconnections.DefaultPrivateEndpointConnectionUpdateOperationOptions()); err != nil {
	// handle the error
}
```
