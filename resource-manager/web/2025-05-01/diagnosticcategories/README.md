
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategories` Documentation

The `diagnosticcategories` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategories"
```


### Client Initialization

```go
client := diagnosticcategories.NewDiagnosticCategoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticCategoriesClient.DiagnosticsGetSiteDiagnosticCategory`

```go
ctx := context.TODO()
id := diagnosticcategories.NewDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "diagnosticName")

read, err := client.DiagnosticsGetSiteDiagnosticCategory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticCategoriesClient.DiagnosticsListSiteDiagnosticCategories`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.DiagnosticsListSiteDiagnosticCategories(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteDiagnosticCategoriesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
