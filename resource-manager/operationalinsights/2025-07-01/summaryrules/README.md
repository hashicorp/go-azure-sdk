
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-07-01/summaryrules` Documentation

The `summaryrules` SDK allows for interaction with Azure Resource Manager `operationalinsights` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-07-01/summaryrules"
```


### Client Initialization

```go
client := summaryrules.NewSummaryRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SummaryRulesClient.SummaryLogsCreateOrUpdate`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

payload := summaryrules.SummaryLogs{
	// ...
}


if err := client.SummaryLogsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsDelete`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

if err := client.SummaryLogsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsGet`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

read, err := client.SummaryLogsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsListByWorkspace`

```go
ctx := context.TODO()
id := summaryrules.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.SummaryLogsListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.SummaryLogsListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsRetryBin`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

payload := summaryrules.SummaryLogsRetryBin{
	// ...
}


if err := client.SummaryLogsRetryBinThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsStart`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

if err := client.SummaryLogsStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SummaryRulesClient.SummaryLogsStop`

```go
ctx := context.TODO()
id := summaryrules.NewSummaryLogID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "summaryLogName")

read, err := client.SummaryLogsStop(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
