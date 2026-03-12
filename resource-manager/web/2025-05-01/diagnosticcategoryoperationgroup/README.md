
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategoryoperationgroup` Documentation

The `diagnosticcategoryoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/diagnosticcategoryoperationgroup"
```


### Client Initialization

```go
client := diagnosticcategoryoperationgroup.NewDiagnosticCategoryOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticCategoryOperationGroupClient.DiagnosticsGetSiteDiagnosticCategorySlot`

```go
ctx := context.TODO()
id := diagnosticcategoryoperationgroup.NewSlotDiagnosticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "diagnosticName")

read, err := client.DiagnosticsGetSiteDiagnosticCategorySlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticCategoryOperationGroupClient.DiagnosticsListSiteDiagnosticCategoriesSlot`

```go
ctx := context.TODO()
id := diagnosticcategoryoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.DiagnosticsListSiteDiagnosticCategoriesSlot(ctx, id)` can be used to do batched pagination
items, err := client.DiagnosticsListSiteDiagnosticCategoriesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
