
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespace` Documentation

The `defenderresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespace"
```


### Client Initialization

```go
client := defenderresourcenamespace.NewDefenderResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DefenderResourceNamespaceClient.CreateDefenderResourceNamespace`

```go
ctx := context.TODO()

payload := defenderresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateDefenderResourceNamespace(ctx, payload, defenderresourcenamespace.DefaultCreateDefenderResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderResourceNamespaceClient.CreateDefenderResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := defenderresourcenamespace.NewRoleManagementDefenderResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := defenderresourcenamespace.CreateDefenderResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateDefenderResourceNamespaceImportResourceAction(ctx, id, payload, defenderresourcenamespace.DefaultCreateDefenderResourceNamespaceImportResourceActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderResourceNamespaceClient.DeleteDefenderResourceNamespace`

```go
ctx := context.TODO()
id := defenderresourcenamespace.NewRoleManagementDefenderResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.DeleteDefenderResourceNamespace(ctx, id, defenderresourcenamespace.DefaultDeleteDefenderResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderResourceNamespaceClient.GetDefenderResourceNamespace`

```go
ctx := context.TODO()
id := defenderresourcenamespace.NewRoleManagementDefenderResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.GetDefenderResourceNamespace(ctx, id, defenderresourcenamespace.DefaultGetDefenderResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderResourceNamespaceClient.GetDefenderResourceNamespacesCount`

```go
ctx := context.TODO()


read, err := client.GetDefenderResourceNamespacesCount(ctx, defenderresourcenamespace.DefaultGetDefenderResourceNamespacesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DefenderResourceNamespaceClient.ListDefenderResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListDefenderResourceNamespaces(ctx, defenderresourcenamespace.DefaultListDefenderResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListDefenderResourceNamespacesComplete(ctx, defenderresourcenamespace.DefaultListDefenderResourceNamespacesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DefenderResourceNamespaceClient.UpdateDefenderResourceNamespace`

```go
ctx := context.TODO()
id := defenderresourcenamespace.NewRoleManagementDefenderResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := defenderresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateDefenderResourceNamespace(ctx, id, payload, defenderresourcenamespace.DefaultUpdateDefenderResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
