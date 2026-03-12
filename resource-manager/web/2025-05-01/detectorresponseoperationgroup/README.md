
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectorresponseoperationgroup` Documentation

The `detectorresponseoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectorresponseoperationgroup"
```


### Client Initialization

```go
client := detectorresponseoperationgroup.NewDetectorResponseOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DetectorResponseOperationGroupClient.DiagnosticsGetSiteDetectorResponseSlot`

```go
ctx := context.TODO()
id := detectorresponseoperationgroup.NewSlotDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "detectorName")

read, err := client.DiagnosticsGetSiteDetectorResponseSlot(ctx, id, detectorresponseoperationgroup.DefaultDiagnosticsGetSiteDetectorResponseSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorResponseOperationGroupClient.DiagnosticsListSiteDetectorResponsesSlot`

```go
ctx := context.TODO()
id := detectorresponseoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.DiagnosticsListSiteDetectorResponsesSlot(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteDetectorResponsesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
