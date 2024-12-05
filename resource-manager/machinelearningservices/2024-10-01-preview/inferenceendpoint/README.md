
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferenceendpoint` Documentation

The `inferenceendpoint` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferenceendpoint"
```


### Client Initialization

```go
client := inferenceendpoint.NewInferenceEndpointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InferenceEndpointClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := inferenceendpoint.NewInferencePoolEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "endpointName")

payload := inferenceendpoint.InferenceEndpointTrackedResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InferenceEndpointClient.Delete`

```go
ctx := context.TODO()
id := inferenceendpoint.NewInferencePoolEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "endpointName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `InferenceEndpointClient.Get`

```go
ctx := context.TODO()
id := inferenceendpoint.NewInferencePoolEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "endpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InferenceEndpointClient.List`

```go
ctx := context.TODO()
id := inferenceendpoint.NewInferencePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName")

// alternatively `client.List(ctx, id, inferenceendpoint.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, inferenceendpoint.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InferenceEndpointClient.Update`

```go
ctx := context.TODO()
id := inferenceendpoint.NewInferencePoolEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "endpointName")
var payload interface{}

if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
