
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagemententerpriseappresourcenamespace` Documentation

The `rolemanagemententerpriseappresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagemententerpriseappresourcenamespace"
```


### Client Initialization

```go
client := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.CreateRoleManagementEnterpriseAppByIdResourceNamespace`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

payload := rolemanagemententerpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementEnterpriseAppByIdResourceNamespace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.CreateRoleManagementEnterpriseAppByIdResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

payload := rolemanagemententerpriseappresourcenamespace.CreateRoleManagementEnterpriseAppByIdResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementEnterpriseAppByIdResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.DeleteRoleManagementEnterpriseAppByIdResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementEnterpriseAppByIdResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.GetRoleManagementEnterpriseAppByIdResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementEnterpriseAppByIdResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.GetRoleManagementEnterpriseAppByIdResourceNamespaceCount`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

read, err := client.GetRoleManagementEnterpriseAppByIdResourceNamespaceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.ListRoleManagementEnterpriseAppByIdResourceNamespaces`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

// alternatively `client.ListRoleManagementEnterpriseAppByIdResourceNamespaces(ctx, id)` can be used to do batched pagination
items, err := client.ListRoleManagementEnterpriseAppByIdResourceNamespacesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementEnterpriseAppResourceNamespaceClient.UpdateRoleManagementEnterpriseAppByIdResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

payload := rolemanagemententerpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementEnterpriseAppByIdResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
