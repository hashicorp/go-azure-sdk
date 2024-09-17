
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/workflowtriggers` Documentation

The `workflowtriggers` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/workflowtriggers"
```


### Client Initialization

```go
client := workflowtriggers.NewWorkflowTriggersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowTriggersClient.Get`

```go
ctx := context.TODO()
id := workflowtriggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "triggerValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowTriggersClient.GetSchemaJson`

```go
ctx := context.TODO()
id := workflowtriggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "triggerValue")

read, err := client.GetSchemaJson(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowTriggersClient.List`

```go
ctx := context.TODO()
id := workflowtriggers.NewManagementWorkflowID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue")

// alternatively `client.List(ctx, id, workflowtriggers.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowtriggers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowTriggersClient.ListCallbackUrl`

```go
ctx := context.TODO()
id := workflowtriggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "triggerValue")

read, err := client.ListCallbackUrl(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowTriggersClient.Run`

```go
ctx := context.TODO()
id := workflowtriggers.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "triggerValue")

if err := client.RunThenPoll(ctx, id); err != nil {
	// handle the error
}
```
