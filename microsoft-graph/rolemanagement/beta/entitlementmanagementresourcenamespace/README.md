
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespace` Documentation

The `entitlementmanagementresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespace"
```


### Client Initialization

```go
client := entitlementmanagementresourcenamespace.NewEntitlementManagementResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.CreateEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()

payload := entitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateEntitlementManagementResourceNamespace(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.CreateEntitlementManagementResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := entitlementmanagementresourcenamespace.CreateEntitlementManagementResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateEntitlementManagementResourceNamespaceImportResourceAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.DeleteEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteEntitlementManagementResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.GetEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

read, err := client.GetEntitlementManagementResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.GetEntitlementManagementResourceNamespaceCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementResourceNamespaceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.ListEntitlementManagementResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListEntitlementManagementResourceNamespaces(ctx)` can be used to do batched pagination
items, err := client.ListEntitlementManagementResourceNamespacesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.UpdateEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceIdValue")

payload := entitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateEntitlementManagementResourceNamespace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
