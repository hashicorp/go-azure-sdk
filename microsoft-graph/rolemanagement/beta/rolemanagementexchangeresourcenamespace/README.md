
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementexchangeresourcenamespace` Documentation

The `rolemanagementexchangeresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagementexchangeresourcenamespace"
```


### Client Initialization

```go
client := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.CreateRoleManagementExchangeResourceNamespace`

```go
ctx := context.TODO()

payload := rolemanagementexchangeresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementExchangeResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.CreateRoleManagementExchangeResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementexchangeresourcenamespace.CreateRoleManagementExchangeResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementExchangeResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.DeleteRoleManagementExchangeResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementExchangeResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.GetRoleManagementExchangeResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementExchangeResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.GetRoleManagementExchangeResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementExchangeResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.ListRoleManagementExchangeResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementExchangeResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListRoleManagementExchangeResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementExchangeResourceNamespaceClient.UpdateRoleManagementExchangeResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagementexchangeresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementExchangeResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
