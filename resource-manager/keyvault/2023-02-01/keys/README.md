
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/keys` Documentation

The `keys` SDK allows for interaction with the Azure Resource Manager Service `keyvault` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/keys"
```


### Client Initialization

```go
client := keys.NewKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeysClient.CreateIfNotExist`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "keyValue")

payload := keys.KeyCreateParameters{
	// ...
}


read, err := client.CreateIfNotExist(ctx, id, payload)
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
id := commonids.NewKeyVaultKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "keyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeysClient.GetVersion`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultKeyVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "keyValue", "versionValue")

read, err := client.GetVersion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeysClient.List`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `KeysClient.ListVersions`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "keyValue")

// alternatively `client.ListVersions(ctx, id)` can be used to do batched pagination
items, err := client.ListVersionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
