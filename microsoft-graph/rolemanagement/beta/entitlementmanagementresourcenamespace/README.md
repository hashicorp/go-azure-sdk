
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespace` Documentation

The `entitlementmanagementresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespace"
```


### Client Initialization

```go
client := entitlementmanagementresourcenamespace.NewEntitlementManagementResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.CreateEntitlementManagementResourceNamespace`

```go
ctx := context.TODO()

payload := entitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateEntitlementManagementResourceNamespace(ctx, payload, entitlementmanagementresourcenamespace.DefaultCreateEntitlementManagementResourceNamespaceOperationOptions())
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
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := entitlementmanagementresourcenamespace.CreateEntitlementManagementResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateEntitlementManagementResourceNamespaceImportResourceAction(ctx, id, payload, entitlementmanagementresourcenamespace.DefaultCreateEntitlementManagementResourceNamespaceImportResourceActionOperationOptions())
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
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.DeleteEntitlementManagementResourceNamespace(ctx, id, entitlementmanagementresourcenamespace.DefaultDeleteEntitlementManagementResourceNamespaceOperationOptions())
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
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.GetEntitlementManagementResourceNamespace(ctx, id, entitlementmanagementresourcenamespace.DefaultGetEntitlementManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EntitlementManagementResourceNamespaceClient.GetEntitlementManagementResourceNamespacesCount`

```go
ctx := context.TODO()


read, err := client.GetEntitlementManagementResourceNamespacesCount(ctx, entitlementmanagementresourcenamespace.DefaultGetEntitlementManagementResourceNamespacesCountOperationOptions())
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


// alternatively `client.ListEntitlementManagementResourceNamespaces(ctx, entitlementmanagementresourcenamespace.DefaultListEntitlementManagementResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListEntitlementManagementResourceNamespacesComplete(ctx, entitlementmanagementresourcenamespace.DefaultListEntitlementManagementResourceNamespacesOperationOptions())
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
id := entitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := entitlementmanagementresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateEntitlementManagementResourceNamespace(ctx, id, payload, entitlementmanagementresourcenamespace.DefaultUpdateEntitlementManagementResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
