
## `github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/runbookdraft` Documentation

The `runbookdraft` SDK allows for interaction with Azure Resource Manager `automation` (API Version `2015-10-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/runbookdraft"
```


### Client Initialization

```go
client := runbookdraft.NewRunbookDraftClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RunbookDraftClient.Get`

```go
ctx := context.TODO()
id := runbookdraft.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RunbookDraftClient.GetContent`

```go
ctx := context.TODO()
id := runbookdraft.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")

read, err := client.GetContent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RunbookDraftClient.Publish`

```go
ctx := context.TODO()
id := runbookdraft.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")

if err := client.PublishThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RunbookDraftClient.ReplaceContent`

```go
ctx := context.TODO()
id := runbookdraft.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")
var payload []byte

if err := client.ReplaceContentThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RunbookDraftClient.UndoEdit`

```go
ctx := context.TODO()
id := runbookdraft.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")

read, err := client.UndoEdit(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
