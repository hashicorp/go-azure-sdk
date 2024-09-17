
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-10-01-preview/encryptionscopes` Documentation

The `encryptionscopes` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2023-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-10-01-preview/encryptionscopes"
```


### Client Initialization

```go
client := encryptionscopes.NewEncryptionScopesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EncryptionScopesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := encryptionscopes.NewEncryptionScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "encryptionScopeValue")

payload := encryptionscopes.EncryptionScope{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EncryptionScopesClient.Delete`

```go
ctx := context.TODO()
id := encryptionscopes.NewEncryptionScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "encryptionScopeValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `EncryptionScopesClient.Get`

```go
ctx := context.TODO()
id := encryptionscopes.NewEncryptionScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "encryptionScopeValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EncryptionScopesClient.List`

```go
ctx := context.TODO()
id := encryptionscopes.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
