
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/keys` Documentation

The `keys` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/keys"
```


### Client Initialization

```go
client := keys.NewKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := keys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "keyValue")

payload := keys.Key{
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


### Example Usage: `KeysClient.Delete`

```go
ctx := context.TODO()
id := keys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "keyValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeysClient.Get`

```go
ctx := context.TODO()
id := keys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "keyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeysClient.ListByWorkspace`

```go
ctx := context.TODO()
id := keys.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
