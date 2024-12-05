
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencegroup` Documentation

The `inferencegroup` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencegroup"
```


### Client Initialization

```go
client := inferencegroup.NewInferenceGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InferenceGroupClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

payload := inferencegroup.InferenceGroupTrackedResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InferenceGroupClient.Delete`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `InferenceGroupClient.Get`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InferenceGroupClient.GetStatus`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

read, err := client.GetStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InferenceGroupClient.List`

```go
ctx := context.TODO()
id := inferencegroup.NewInferencePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName")

// alternatively `client.List(ctx, id, inferencegroup.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, inferencegroup.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InferenceGroupClient.ListSkus`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

// alternatively `client.ListSkus(ctx, id, inferencegroup.DefaultListSkusOperationOptions())` can be used to do batched pagination
items, err := client.ListSkusComplete(ctx, id, inferencegroup.DefaultListSkusOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InferenceGroupClient.Update`

```go
ctx := context.TODO()
id := inferencegroup.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

payload := inferencegroup.PartialMinimalTrackedResourceWithSku{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
