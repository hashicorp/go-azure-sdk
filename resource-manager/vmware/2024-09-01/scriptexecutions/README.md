
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptexecutions` Documentation

The `scriptexecutions` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptexecutions"
```


### Client Initialization

```go
client := scriptexecutions.NewScriptExecutionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScriptExecutionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := scriptexecutions.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

payload := scriptexecutions.ScriptExecution{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ScriptExecutionsClient.Delete`

```go
ctx := context.TODO()
id := scriptexecutions.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ScriptExecutionsClient.Get`

```go
ctx := context.TODO()
id := scriptexecutions.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptExecutionsClient.GetExecutionLogs`

```go
ctx := context.TODO()
id := scriptexecutions.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")
var payload []ScriptOutputStreamType

read, err := client.GetExecutionLogs(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptExecutionsClient.List`

```go
ctx := context.TODO()
id := scriptexecutions.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
