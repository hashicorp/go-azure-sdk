
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/workflows` Documentation

The `workflows` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/workflows"
```


### Client Initialization

```go
client := workflows.NewWorkflowsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowsClient.RegenerateAccessKey`

```go
ctx := context.TODO()
id := workflows.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue")

payload := workflows.RegenerateActionParameter{
	// ...
}


read, err := client.RegenerateAccessKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowsClient.Validate`

```go
ctx := context.TODO()
id := workflows.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue")

payload := workflows.Workflow{
	// ...
}


read, err := client.Validate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
