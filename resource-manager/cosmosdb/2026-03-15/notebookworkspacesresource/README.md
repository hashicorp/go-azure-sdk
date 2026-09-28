
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/notebookworkspacesresource` Documentation

The `notebookworkspacesresource` SDK allows for interaction with Azure Resource Manager `cosmosdb` (API Version `2026-03-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/notebookworkspacesresource"
```


### Client Initialization

```go
client := notebookworkspacesresource.NewNotebookWorkspacesResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesCreateOrUpdate`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

payload := notebookworkspacesresource.ARMProxyResource{
	// ...
}


if err := client.NotebookWorkspacesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesDelete`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

if err := client.NotebookWorkspacesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesGet`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

read, err := client.NotebookWorkspacesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesListByDatabaseAccount`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

// alternatively `client.NotebookWorkspacesListByDatabaseAccount(ctx, id)` can be used to do batched pagination
items, err := client.NotebookWorkspacesListByDatabaseAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesListConnectionInfo`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

read, err := client.NotebookWorkspacesListConnectionInfo(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesRegenerateAuthToken`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

if err := client.NotebookWorkspacesRegenerateAuthTokenThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NotebookWorkspacesResourceClient.NotebookWorkspacesStart`

```go
ctx := context.TODO()
id := notebookworkspacesresource.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

if err := client.NotebookWorkspacesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```
