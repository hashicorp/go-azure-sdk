
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/proxyoperations` Documentation

The `proxyoperations` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/proxyoperations"
```


### Client Initialization

```go
client := proxyoperations.NewProxyOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProxyOperationsClient.WorkspacesDiagnose`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

payload := proxyoperations.DiagnoseWorkspaceParameters{
	// ...
}


if err := client.WorkspacesDiagnoseThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesListKeys`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.WorkspacesListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesListNotebookAccessToken`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.WorkspacesListNotebookAccessToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesListNotebookKeys`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.WorkspacesListNotebookKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesListStorageAccountKeys`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.WorkspacesListStorageAccountKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesPrepareNotebook`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

if err := client.WorkspacesPrepareNotebookThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ProxyOperationsClient.WorkspacesResyncKeys`

```go
ctx := context.TODO()
id := proxyoperations.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

if err := client.WorkspacesResyncKeysThenPoll(ctx, id); err != nil {
	// handle the error
}
```
