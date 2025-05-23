
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/exports` Documentation

The `exports` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/exports"
```


### Client Initialization

```go
client := exports.NewExportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExportsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := exports.NewScopedExportID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "exportName")

payload := exports.Export{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExportsClient.Delete`

```go
ctx := context.TODO()
id := exports.NewScopedExportID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "exportName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExportsClient.Execute`

```go
ctx := context.TODO()
id := exports.NewScopedExportID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "exportName")

payload := exports.ExportRunRequest{
	// ...
}


read, err := client.Execute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExportsClient.Get`

```go
ctx := context.TODO()
id := exports.NewScopedExportID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "exportName")

read, err := client.Get(ctx, id, exports.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExportsClient.GetExecutionHistory`

```go
ctx := context.TODO()
id := exports.NewScopedExportID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "exportName")

read, err := client.GetExecutionHistory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExportsClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id, exports.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
