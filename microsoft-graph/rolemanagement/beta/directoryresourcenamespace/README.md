
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespace` Documentation

The `directoryresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespace"
```


### Client Initialization

```go
client := directoryresourcenamespace.NewDirectoryResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryResourceNamespaceClient.CreateDirectoryResourceNamespace`

```go
ctx := context.TODO()

payload := directoryresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateDirectoryResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryResourceNamespaceClient.CreateDirectoryResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := directoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := directoryresourcenamespace.CreateDirectoryResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateDirectoryResourceNamespaceImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryResourceNamespaceClient.DeleteDirectoryResourceNamespace`

```go
ctx := context.TODO()
id := directoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteDirectoryResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryResourceNamespaceClient.GetDirectoryResourceNamespace`

```go
ctx := context.TODO()
id := directoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetDirectoryResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryResourceNamespaceClient.GetDirectoryResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryResourceNamespaceClient.ListDirectoryResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryResourceNamespaceClient.UpdateDirectoryResourceNamespace`

```go
ctx := context.TODO()
id := directoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := directoryresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateDirectoryResourceNamespace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
