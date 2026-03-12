
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjoboperationgroup` Documentation

The `continuouswebjoboperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/continuouswebjoboperationgroup"
```


### Client Initialization

```go
client := continuouswebjoboperationgroup.NewContinuousWebJobOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContinuousWebJobOperationGroupClient.WebAppsDeleteContinuousWebJobSlot`

```go
ctx := context.TODO()
id := continuouswebjoboperationgroup.NewSlotContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "continuousWebJobName")

read, err := client.WebAppsDeleteContinuousWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobOperationGroupClient.WebAppsGetContinuousWebJobSlot`

```go
ctx := context.TODO()
id := continuouswebjoboperationgroup.NewSlotContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "continuousWebJobName")

read, err := client.WebAppsGetContinuousWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobOperationGroupClient.WebAppsListContinuousWebJobsSlot`

```go
ctx := context.TODO()
id := continuouswebjoboperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListContinuousWebJobsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListContinuousWebJobsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContinuousWebJobOperationGroupClient.WebAppsStartContinuousWebJobSlot`

```go
ctx := context.TODO()
id := continuouswebjoboperationgroup.NewSlotContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "continuousWebJobName")

read, err := client.WebAppsStartContinuousWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContinuousWebJobOperationGroupClient.WebAppsStopContinuousWebJobSlot`

```go
ctx := context.TODO()
id := continuouswebjoboperationgroup.NewSlotContinuousWebJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "continuousWebJobName")

read, err := client.WebAppsStopContinuousWebJobSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
