
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/diagnostics` Documentation

The `diagnostics` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/diagnostics"
```


### Client Initialization

```go
client := diagnostics.NewDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsClient.ExecuteSiteAnalysis`

```go
ctx := context.TODO()
id := diagnostics.NewAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue", "analysisValue")

read, err := client.ExecuteSiteAnalysis(ctx, id, diagnostics.DefaultExecuteSiteAnalysisOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ExecuteSiteAnalysisSlot`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue", "analysisValue")

read, err := client.ExecuteSiteAnalysisSlot(ctx, id, diagnostics.DefaultExecuteSiteAnalysisSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ExecuteSiteDetector`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue", "detectorValue")

read, err := client.ExecuteSiteDetector(ctx, id, diagnostics.DefaultExecuteSiteDetectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ExecuteSiteDetectorSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue", "detectorValue")

read, err := client.ExecuteSiteDetectorSlot(ctx, id, diagnostics.DefaultExecuteSiteDetectorSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetHostingEnvironmentDetectorResponse`

```go
ctx := context.TODO()
id := diagnostics.NewHostingEnvironmentDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentValue", "detectorValue")

read, err := client.GetHostingEnvironmentDetectorResponse(ctx, id, diagnostics.DefaultGetHostingEnvironmentDetectorResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteAnalysis`

```go
ctx := context.TODO()
id := diagnostics.NewAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue", "analysisValue")

read, err := client.GetSiteAnalysis(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteAnalysisSlot`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue", "analysisValue")

read, err := client.GetSiteAnalysisSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDetector`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue", "detectorValue")

read, err := client.GetSiteDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDetectorResponse`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "detectorValue")

read, err := client.GetSiteDetectorResponse(ctx, id, diagnostics.DefaultGetSiteDetectorResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDetectorResponseSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "detectorValue")

read, err := client.GetSiteDetectorResponseSlot(ctx, id, diagnostics.DefaultGetSiteDetectorResponseSlotOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDetectorSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue", "detectorValue")

read, err := client.GetSiteDetectorSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDiagnosticCategory`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue")

read, err := client.GetSiteDiagnosticCategory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.GetSiteDiagnosticCategorySlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue")

read, err := client.GetSiteDiagnosticCategorySlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ListHostingEnvironmentDetectorResponses`

```go
ctx := context.TODO()
id := diagnostics.NewHostingEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostingEnvironmentValue")

// alternatively `client.ListHostingEnvironmentDetectorResponses(ctx, id)` can be used to do batched pagination
items, err := client.ListHostingEnvironmentDetectorResponsesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteAnalyses`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue")

// alternatively `client.ListSiteAnalyses(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteAnalysesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteAnalysesSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue")

// alternatively `client.ListSiteAnalysesSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteAnalysesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDetectorResponses`

```go
ctx := context.TODO()
id := diagnostics.NewSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue")

// alternatively `client.ListSiteDetectorResponses(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDetectorResponsesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDetectorResponsesSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue")

// alternatively `client.ListSiteDetectorResponsesSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDetectorResponsesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDetectors`

```go
ctx := context.TODO()
id := diagnostics.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "diagnosticValue")

// alternatively `client.ListSiteDetectors(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDetectorsSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue", "diagnosticValue")

// alternatively `client.ListSiteDetectorsSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDetectorsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDiagnosticCategories`

```go
ctx := context.TODO()
id := diagnostics.NewSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue")

// alternatively `client.ListSiteDiagnosticCategories(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDiagnosticCategoriesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ListSiteDiagnosticCategoriesSlot`

```go
ctx := context.TODO()
id := diagnostics.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "slotValue")

// alternatively `client.ListSiteDiagnosticCategoriesSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDiagnosticCategoriesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
