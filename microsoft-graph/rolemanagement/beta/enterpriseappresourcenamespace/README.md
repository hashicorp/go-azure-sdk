
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespace` Documentation

The `enterpriseappresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespace"
```


### Client Initialization

```go
client := enterpriseappresourcenamespace.NewEnterpriseAppResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.CreateEnterpriseAppResourceNamespace`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationId")

payload := enterpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateEnterpriseAppResourceNamespace(ctx, id, payload, enterpriseappresourcenamespace.DefaultCreateEnterpriseAppResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.CreateEnterpriseAppResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId")

payload := enterpriseappresourcenamespace.CreateEnterpriseAppResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateEnterpriseAppResourceNamespaceImportResourceAction(ctx, id, payload, enterpriseappresourcenamespace.DefaultCreateEnterpriseAppResourceNamespaceImportResourceActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.DeleteEnterpriseAppResourceNamespace`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId")

read, err := client.DeleteEnterpriseAppResourceNamespace(ctx, id, enterpriseappresourcenamespace.DefaultDeleteEnterpriseAppResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.GetEnterpriseAppResourceNamespace`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId")

read, err := client.GetEnterpriseAppResourceNamespace(ctx, id, enterpriseappresourcenamespace.DefaultGetEnterpriseAppResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.GetEnterpriseAppResourceNamespacesCount`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationId")

read, err := client.GetEnterpriseAppResourceNamespacesCount(ctx, id, enterpriseappresourcenamespace.DefaultGetEnterpriseAppResourceNamespacesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.ListEnterpriseAppResourceNamespaces`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationId")

// alternatively `client.ListEnterpriseAppResourceNamespaces(ctx, id, enterpriseappresourcenamespace.DefaultListEnterpriseAppResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListEnterpriseAppResourceNamespacesComplete(ctx, id, enterpriseappresourcenamespace.DefaultListEnterpriseAppResourceNamespacesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.UpdateEnterpriseAppResourceNamespace`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationId", "unifiedRbacResourceNamespaceId")

payload := enterpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateEnterpriseAppResourceNamespace(ctx, id, payload, enterpriseappresourcenamespace.DefaultUpdateEnterpriseAppResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
