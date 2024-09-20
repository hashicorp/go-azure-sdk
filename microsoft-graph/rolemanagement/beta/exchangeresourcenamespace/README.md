
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespace` Documentation

The `exchangeresourcenamespace` SDK allows for interaction with Microsoft Graph `rolemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespace"
```


### Client Initialization

```go
client := exchangeresourcenamespace.NewExchangeResourceNamespaceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExchangeResourceNamespaceClient.CreateExchangeResourceNamespace`

```go
ctx := context.TODO()

payload := exchangeresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.CreateExchangeResourceNamespace(ctx, payload, exchangeresourcenamespace.DefaultCreateExchangeResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeResourceNamespaceClient.CreateExchangeResourceNamespaceImportResourceAction`

```go
ctx := context.TODO()
id := exchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := exchangeresourcenamespace.CreateExchangeResourceNamespaceImportResourceActionRequest{
	// ...
}


read, err := client.CreateExchangeResourceNamespaceImportResourceAction(ctx, id, payload, exchangeresourcenamespace.DefaultCreateExchangeResourceNamespaceImportResourceActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeResourceNamespaceClient.DeleteExchangeResourceNamespace`

```go
ctx := context.TODO()
id := exchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.DeleteExchangeResourceNamespace(ctx, id, exchangeresourcenamespace.DefaultDeleteExchangeResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeResourceNamespaceClient.GetExchangeResourceNamespace`

```go
ctx := context.TODO()
id := exchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceId")

read, err := client.GetExchangeResourceNamespace(ctx, id, exchangeresourcenamespace.DefaultGetExchangeResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeResourceNamespaceClient.GetExchangeResourceNamespacesCount`

```go
ctx := context.TODO()


read, err := client.GetExchangeResourceNamespacesCount(ctx, exchangeresourcenamespace.DefaultGetExchangeResourceNamespacesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeResourceNamespaceClient.ListExchangeResourceNamespaces`

```go
ctx := context.TODO()


// alternatively `client.ListExchangeResourceNamespaces(ctx, exchangeresourcenamespace.DefaultListExchangeResourceNamespacesOperationOptions())` can be used to do batched pagination
items, err := client.ListExchangeResourceNamespacesComplete(ctx, exchangeresourcenamespace.DefaultListExchangeResourceNamespacesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExchangeResourceNamespaceClient.UpdateExchangeResourceNamespace`

```go
ctx := context.TODO()
id := exchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceID("unifiedRbacResourceNamespaceId")

payload := exchangeresourcenamespace.UnifiedRbacResourceNamespace{
	// ...
}


read, err := client.UpdateExchangeResourceNamespace(ctx, id, payload, exchangeresourcenamespace.DefaultUpdateExchangeResourceNamespaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
