
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/workflowresource` Documentation

The `workflowresource` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2020-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2020-03-01/workflowresource"
```


### Client Initialization

```go
client := workflowresource.NewWorkflowResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowResourceClient.WorkflowsAbort`

```go
ctx := context.TODO()
id := workflowresource.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "workflowIdValue")

read, err := client.WorkflowsAbort(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowResourceClient.WorkflowsGet`

```go
ctx := context.TODO()
id := workflowresource.NewWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue", "workflowIdValue")

read, err := client.WorkflowsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowResourceClient.WorkflowsListByStorageSyncService`

```go
ctx := context.TODO()
id := workflowresource.NewStorageSyncServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue")

read, err := client.WorkflowsListByStorageSyncService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
