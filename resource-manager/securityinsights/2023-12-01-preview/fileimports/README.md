
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/fileimports` Documentation

The `fileimports` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/fileimports"
```


### Client Initialization

```go
client := fileimports.NewFileImportsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FileImportsClient.Create`

```go
ctx := context.TODO()
id := fileimports.NewFileImportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "fileImportId")

payload := fileimports.FileImport{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileImportsClient.Delete`

```go
ctx := context.TODO()
id := fileimports.NewFileImportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "fileImportId")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `FileImportsClient.Get`

```go
ctx := context.TODO()
id := fileimports.NewFileImportID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "fileImportId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FileImportsClient.List`

```go
ctx := context.TODO()
id := fileimports.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, fileimports.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, fileimports.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
