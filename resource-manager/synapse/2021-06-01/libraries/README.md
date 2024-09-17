
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/libraries` Documentation

The `libraries` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/libraries"
```


### Client Initialization

```go
client := libraries.NewLibrariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LibrariesClient.LibraryGet`

```go
ctx := context.TODO()
id := libraries.NewLibraryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "libraryValue")

read, err := client.LibraryGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LibrariesClient.ListByWorkspace`

```go
ctx := context.TODO()
id := libraries.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
