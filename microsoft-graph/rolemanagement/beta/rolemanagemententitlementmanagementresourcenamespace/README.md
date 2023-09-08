
## `github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespace` Documentation

The `rolemanagemententitlementmanagementresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespace"
```


### Client Initialization

```go
client := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.CreateRoleManagementEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()

payload := rolemanagemententitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateRoleManagementEntitlementManagementResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.CreateRoleManagementEntitlementManagementResourceNamespaceByIdImportResourceAction`

```go
ctx := context.TODO()
id := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagemententitlementmanagementresourcenamespace.CreateRoleManagementEntitlementManagementResourceNamespaceByIdImportResourceActionRequest{
	// ...
}


read, err := client.CreateRoleManagementEntitlementManagementResourceNamespaceByIdImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.DeleteRoleManagementEntitlementManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteRoleManagementEntitlementManagementResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.GetRoleManagementEntitlementManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetRoleManagementEntitlementManagementResourceNamespaceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.GetRoleManagementEntitlementManagementResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementEntitlementManagementResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.ListRoleManagementEntitlementManagementResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementEntitlementManagementResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListRoleManagementEntitlementManagementResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementEntitlementManagementResourceNamespaceClient.UpdateRoleManagementEntitlementManagementResourceNamespaceById`

```go
ctx := context.TODO()
id := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := rolemanagemententitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateRoleManagementEntitlementManagementResourceNamespaceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
