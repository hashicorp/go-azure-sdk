
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/workflows` Documentation

The `workflows` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/workflows"
```


### Client Initialization

```go
client := workflows.NewWorkflowsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowsClient.Abort`

```go
ctx := context.TODO()
id := workflows.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "workflowId")

read, err := client.Abort(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowsClient.Get`

```go
ctx := context.TODO()
id := workflows.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "workflowId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowsClient.ListByStorageSyncService`

```go
ctx := context.TODO()
id := workflows.NewStorageSyncServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName")

// alternatively `client.ListByStorageSyncService(ctx, id)` can be used to do batched pagination
items, err := client.ListByStorageSyncServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
