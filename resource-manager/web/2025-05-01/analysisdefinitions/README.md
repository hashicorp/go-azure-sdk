
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitions` Documentation

The `analysisdefinitions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/analysisdefinitions"
```


### Client Initialization

```go
client := analysisdefinitions.NewAnalysisDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AnalysisDefinitionsClient.DiagnosticsExecuteSiteAnalysis`

```go
ctx := context.TODO()
id := analysisdefinitions.NewAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName", "analysisName")

read, err := client.DiagnosticsExecuteSiteAnalysis(ctx, id, analysisdefinitions.DefaultDiagnosticsExecuteSiteAnalysisOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AnalysisDefinitionsClient.DiagnosticsGetSiteAnalysis`

```go
ctx := context.TODO()
id := analysisdefinitions.NewAnalysisID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName", "analysisName")

read, err := client.DiagnosticsGetSiteAnalysis(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AnalysisDefinitionsClient.DiagnosticsListSiteAnalyses`

```go
ctx := context.TODO()
id := analysisdefinitions.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName")

// alternatively `client.DiagnosticsListSiteAnalyses(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteAnalysesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
