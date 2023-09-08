
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementcloudpcresourcenamespace` Documentation

The `rolemanagementcloudpcresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementcloudpcresourcenamespace"
```


### Client Initialization

```go
client := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.CreateRoleManagementCloudPCResourceNamespace`

```go
ctx := context.TODO()

payload := rolemanagementcloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementCloudPCResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.CreateRoleManagementCloudPCResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementcloudpcresourcenamespace.CreateRoleManagementCloudPCResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementCloudPCResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.DeleteRoleManagementCloudPCResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementCloudPCResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.GetRoleManagementCloudPCResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementCloudPCResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.GetRoleManagementCloudPCResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementCloudPCResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.ListRoleManagementCloudPCResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementCloudPCResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListRoleManagementCloudPCResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementCloudPCResourceNamespaceClient.UpdateRoleManagementCloudPCResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementcloudpcresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementCloudPCResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
