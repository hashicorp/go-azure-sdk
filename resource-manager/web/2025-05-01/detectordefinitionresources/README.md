
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresources` Documentation

The `detectordefinitionresources` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/detectordefinitionresources"
```


### Client Initialization

```go
client := detectordefinitionresources.NewDetectorDefinitionResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DetectorDefinitionResourcesClient.DiagnosticsExecuteSiteDetector`

```go
ctx := context.TODO()
id := detectordefinitionresources.NewDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName", "detectorName")

read, err := client.DiagnosticsExecuteSiteDetector(ctx, id, detectordefinitionresources.DefaultDiagnosticsExecuteSiteDetectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorDefinitionResourcesClient.DiagnosticsGetSiteDetector`

```go
ctx := context.TODO()
id := detectordefinitionresources.NewDiagnosticDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName", "detectorName")

read, err := client.DiagnosticsGetSiteDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DetectorDefinitionResourcesClient.DiagnosticsListSiteDetectors`

```go
ctx := context.TODO()
id := detectordefinitionresources.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName")

// alternatively `client.DiagnosticsListSiteDetectors(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
