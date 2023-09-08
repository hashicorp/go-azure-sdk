
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementdirectoryresourcenamespace` Documentation

The `rolemanagementdirectoryresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementdirectoryresourcenamespace"
```


### Client Initialization

```go
client := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.CreateRoleManagementDirectoryResourceNamespace`

```go
ctx := context.TODO()

payload := rolemanagementdirectoryresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementDirectoryResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.CreateRoleManagementDirectoryResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementdirectoryresourcenamespace.CreateRoleManagementDirectoryResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementDirectoryResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.DeleteRoleManagementDirectoryResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementDirectoryResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.GetRoleManagementDirectoryResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementDirectoryResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.GetRoleManagementDirectoryResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementDirectoryResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.ListRoleManagementDirectoryResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementDirectoryResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListRoleManagementDirectoryResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementDirectoryResourceNamespaceClient.UpdateRoleManagementDirectoryResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementdirectoryresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementDirectoryResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
