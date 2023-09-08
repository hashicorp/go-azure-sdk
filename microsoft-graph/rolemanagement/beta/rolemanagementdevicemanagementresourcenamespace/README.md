
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespace` Documentation

The `rolemanagementdevicemanagementresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespace"
```


### Client Initialization

```go
client := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.CreateRoleManagementDeviceManagementResourceNamespace`

```go
ctx := context.TODO()

payload := rolemanagementdevicemanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementDeviceManagementResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.CreateRoleManagementDeviceManagementResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementdevicemanagementresourcenamespace.CreateRoleManagementDeviceManagementResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementDeviceManagementResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.DeleteRoleManagementDeviceManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementDeviceManagementResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.GetRoleManagementDeviceManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementDeviceManagementResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.GetRoleManagementDeviceManagementResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementDeviceManagementResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.ListRoleManagementDeviceManagementResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementDeviceManagementResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListRoleManagementDeviceManagementResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementDeviceManagementResourceNamespaceClient.UpdateRoleManagementDeviceManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementdevicemanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementDeviceManagementResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
