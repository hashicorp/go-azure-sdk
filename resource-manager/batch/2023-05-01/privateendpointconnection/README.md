
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/privateendpointconnection` Documentation

The `privateendpointconnection` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/privateendpointconnection"
```


### Client Initialization

```go
client := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "privateEndpointConnectionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "privateEndpointConnectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionClient.ListByBatchAccount`

```go
ctx := context.TODO()
id := privateendpointconnection.NewBatchAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.ListByBatchAccount(ctx, id, privateendpointconnection.DefaultListByBatchAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBatchAccountComplete(ctx, id, privateendpointconnection.DefaultListByBatchAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateEndpointConnectionClient.Update`

```go
ctx := context.TODO()
id := privateendpointconnection.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "privateEndpointConnectionName")

payload := privateendpointconnection.PrivateEndpointConnection{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, privateendpointconnection.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
