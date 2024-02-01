
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/workbooksapis` Documentation

The `workbooksapis` SDK allows for interaction with the Azure Resource Manager Service `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/workbooksapis"
```


### Client Initialization

```go
client := workbooksapis.NewWorkbooksAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkbooksAPIsClient.WorkbooksCreateOrUpdate`

```go
ctx := context.TODO()
id := workbooksapis.NewWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workbookValue")

payload := workbooksapis.Workbook{
	// ...
}


read, err := client.WorkbooksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkbooksAPIsClient.WorkbooksDelete`

```go
ctx := context.TODO()
id := workbooksapis.NewWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workbookValue")

read, err := client.WorkbooksDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkbooksAPIsClient.WorkbooksGet`

```go
ctx := context.TODO()
id := workbooksapis.NewWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workbookValue")

read, err := client.WorkbooksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkbooksAPIsClient.WorkbooksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.WorkbooksListByResourceGroup(ctx, id, workbooksapis.DefaultWorkbooksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkbooksAPIsClient.WorkbooksUpdate`

```go
ctx := context.TODO()
id := workbooksapis.NewWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workbookValue")

payload := workbooksapis.Workbook{
	// ...
}


read, err := client.WorkbooksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
