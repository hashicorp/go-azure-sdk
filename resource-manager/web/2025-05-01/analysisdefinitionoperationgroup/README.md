
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitionoperationgroup` Documentation

The `analysisdefinitionoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitionoperationgroup"
```


### Client Initialization

```go
client := analysisdefinitionoperationgroup.NewAnalysisDefinitionOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AnalysisDefinitionOperationGroupClient.DiagnosticsExecuteSiteAnalysisSlot`

```go
ctx := context.TODO()
id := analysisdefinitionoperationgroup.NewDiagnosticAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName", "analysisName")

read, err := client.DiagnosticsExecuteSiteAnalysisSlot(ctx, id, analysisdefinitionoperationgroup.DefaultDiagnosticsExecuteSiteAnalysisSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AnalysisDefinitionOperationGroupClient.DiagnosticsGetSiteAnalysisSlot`

```go
ctx := context.TODO()
id := analysisdefinitionoperationgroup.NewDiagnosticAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName", "analysisName")

read, err := client.DiagnosticsGetSiteAnalysisSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AnalysisDefinitionOperationGroupClient.DiagnosticsListSiteAnalysesSlot`

```go
ctx := context.TODO()
id := analysisdefinitionoperationgroup.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName")

// alternatively `client.DiagnosticsListSiteAnalysesSlot(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteAnalysesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
