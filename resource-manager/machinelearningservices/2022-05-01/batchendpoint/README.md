
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/batchendpoint` Documentation

The `batchendpoint` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/batchendpoint"
```


### Client Initialization

```go
client := batchendpoint.NewBatchEndpointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `BatchEndpointClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := batchendpoint.NewBatchEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "endpointValue")

payload := batchendpoint.BatchEndpointTrackedResource{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BatchEndpointClient.Delete`

```go
ctx := context.TODO()
id := batchendpoint.NewBatchEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "endpointValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `BatchEndpointClient.Get`

```go
ctx := context.TODO()
id := batchendpoint.NewBatchEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "endpointValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchEndpointClient.List`

```go
ctx := context.TODO()
id := batchendpoint.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.List(ctx, id, batchendpoint.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, batchendpoint.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BatchEndpointClient.ListKeys`

```go
ctx := context.TODO()
id := batchendpoint.NewBatchEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "endpointValue")
read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BatchEndpointClient.Update`

```go
ctx := context.TODO()
id := batchendpoint.NewBatchEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "endpointValue")

payload := batchendpoint.PartialMinimalTrackedResourceWithIdentity{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
