
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/machinelearningcomputes` Documentation

The `machinelearningcomputes` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/machinelearningcomputes"
```


### Client Initialization

```go
client := machinelearningcomputes.NewMachineLearningComputesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeCreateOrUpdate`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")

payload := machinelearningcomputes.ComputeResource{
	// ...
}

future, err := client.ComputeCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeDelete`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
future, err := client.ComputeDelete(ctx, id, machinelearningcomputes.DefaultComputeDeleteOperationOptions())
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeGet`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
read, err := client.ComputeGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachineLearningComputesClient.ComputeList`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.ComputeList(ctx, id, machinelearningcomputes.DefaultComputeListOperationOptions())` can be used to do batched pagination
items, err := client.ComputeListComplete(ctx, id, machinelearningcomputes.DefaultComputeListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachineLearningComputesClient.ComputeListKeys`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
read, err := client.ComputeListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachineLearningComputesClient.ComputeListNodes`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
// alternatively `client.ComputeListNodes(ctx, id)` can be used to do batched pagination
items, err := client.ComputeListNodesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachineLearningComputesClient.ComputeRestart`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
future, err := client.ComputeRestart(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeStart`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
future, err := client.ComputeStart(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeStop`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")
future, err := client.ComputeStop(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `MachineLearningComputesClient.ComputeUpdate`

```go
ctx := context.TODO()
id := machinelearningcomputes.NewComputeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "computeValue")

payload := machinelearningcomputes.ClusterUpdateParameters{
	// ...
}

future, err := client.ComputeUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
