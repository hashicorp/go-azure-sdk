
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencedeltamodel` Documentation

The `inferencedeltamodel` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencedeltamodel"
```


### Client Initialization

```go
client := inferencedeltamodel.NewInferenceDeltaModelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InferenceDeltaModelClient.InferenceGroupsGetDeltaModelsStatusAsync`

```go
ctx := context.TODO()
id := inferencedeltamodel.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

payload := inferencedeltamodel.DeltaModelStatusRequest{
	// ...
}


read, err := client.InferenceGroupsGetDeltaModelsStatusAsync(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InferenceDeltaModelClient.InferenceGroupsListDeltaModelsAsync`

```go
ctx := context.TODO()
id := inferencedeltamodel.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

payload := inferencedeltamodel.DeltaModelListRequest{
	// ...
}


// alternatively `client.InferenceGroupsListDeltaModelsAsync(ctx, id, payload)` can be used to do batched pagination
items, err := client.InferenceGroupsListDeltaModelsAsyncComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InferenceDeltaModelClient.InferenceGroupsModifyDeltaModelsAsync`

```go
ctx := context.TODO()
id := inferencedeltamodel.NewGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "inferencePoolName", "groupName")

payload := inferencedeltamodel.DeltaModelModifyRequest{
	// ...
}


if err := client.InferenceGroupsModifyDeltaModelsAsyncThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
