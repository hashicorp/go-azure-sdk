
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/workflowversions` Documentation

The `workflowversions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/workflowversions"
```


### Client Initialization

```go
client := workflowversions.NewWorkflowVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowVersionsClient.Get`

```go
ctx := context.TODO()
id := workflowversions.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "name", "workflowName", "versionId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowVersionsClient.List`

```go
ctx := context.TODO()
id := workflowversions.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "name", "workflowName")

// alternatively `client.List(ctx, id, workflowversions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowversions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
