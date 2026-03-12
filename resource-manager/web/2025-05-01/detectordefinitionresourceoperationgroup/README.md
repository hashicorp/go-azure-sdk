
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresourceoperationgroup` Documentation

The `detectordefinitionresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresourceoperationgroup"
```


### Client Initialization

```go
client := detectordefinitionresourceoperationgroup.NewDetectorDefinitionResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DetectorDefinitionResourceOperationGroupClient.DiagnosticsExecuteSiteDetectorSlot`

```go
ctx := context.TODO()
id := detectordefinitionresourceoperationgroup.NewSlotDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName", "detectorName")

read, err := client.DiagnosticsExecuteSiteDetectorSlot(ctx, id, detectordefinitionresourceoperationgroup.DefaultDiagnosticsExecuteSiteDetectorSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorDefinitionResourceOperationGroupClient.DiagnosticsGetSiteDetectorSlot`

```go
ctx := context.TODO()
id := detectordefinitionresourceoperationgroup.NewSlotDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName", "detectorName")

read, err := client.DiagnosticsGetSiteDetectorSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorDefinitionResourceOperationGroupClient.DiagnosticsListSiteDetectorsSlot`

```go
ctx := context.TODO()
id := detectordefinitionresourceoperationgroup.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName")

// alternatively `client.DiagnosticsListSiteDetectorsSlot(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteDetectorsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
