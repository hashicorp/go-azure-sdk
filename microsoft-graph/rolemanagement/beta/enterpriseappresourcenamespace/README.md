
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespace` Documentation

The `enterpriseappresourcenamespace` SDK allows for interaction with the Azure Resource Manager Service `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespace"
```


### Client Initialization

```go
client := enterpriseappresourcenamespace.NewEnterpriseAppResourceNamespaceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.CreateEnterpriseAppResourceNamespace`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

payload := enterpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateEnterpriseAppResourceNamespace(ctx, id, payload)
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
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

payload := enterpriseappresourcenamespace.CreateEnterpriseAppResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateEnterpriseAppResourceNamespaceImportResourceAction(ctx, id, payload)
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
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

read, err := client.DeleteEnterpriseAppResourceNamespace(ctx, id)
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
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

read, err := client.GetEnterpriseAppResourceNamespace(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnterpriseAppResourceNamespaceClient.GetEnterpriseAppResourceNamespaceCount`

```go
ctx := context.TODO()
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

read, err := client.GetEnterpriseAppResourceNamespaceCount(ctx, id)
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
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppID("rbacApplicationIdValue")

// alternatively `client.ListEnterpriseAppResourceNamespaces(ctx, id)` can be used to do batched pagination
items, err := client.ListEnterpriseAppResourceNamespacesComplete(ctx, id)
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
id := enterpriseappresourcenamespace.NewRoleManagementEnterpriseAppIdResourceNamespaceID("rbacApplicationIdValue", "unifiedRbacResourceNamespaceIdValue")

payload := enterpriseappresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateEnterpriseAppResourceNamespace(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
