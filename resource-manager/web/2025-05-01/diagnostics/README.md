
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnostics` Documentation

The `diagnostics` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnostics"
```


### Client Initialization

```go
client := diagnostics.NewDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsClient.GetSiteDetectorResponse`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "detectorName")

read, err := client.GetSiteDetectorResponse(ctx, id, diagnostics.DefaultGetSiteDetectorResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ListSiteDetectorResponses`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.ListSiteDetectorResponses(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteDetectorResponsesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
