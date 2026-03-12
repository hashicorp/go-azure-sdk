
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjobs` Documentation

The `triggeredwebjobs` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/triggeredwebjobs"
```


### Client Initialization

```go
client := triggeredwebjobs.NewTriggeredWebJobsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggeredWebJobsClient.WebAppsDeleteTriggeredWebJobSlot`

```go
ctx := context.TODO()
id := triggeredwebjobs.NewSlotTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "triggeredWebJobName")

read, err := client.WebAppsDeleteTriggeredWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggeredWebJobsClient.WebAppsGetTriggeredWebJobSlot`

```go
ctx := context.TODO()
id := triggeredwebjobs.NewSlotTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "triggeredWebJobName")

read, err := client.WebAppsGetTriggeredWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggeredWebJobsClient.WebAppsListTriggeredWebJobsSlot`

```go
ctx := context.TODO()
id := triggeredwebjobs.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListTriggeredWebJobsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListTriggeredWebJobsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TriggeredWebJobsClient.WebAppsRunTriggeredWebJobSlot`

```go
ctx := context.TODO()
id := triggeredwebjobs.NewSlotTriggeredWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "triggeredWebJobName")

read, err := client.WebAppsRunTriggeredWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
