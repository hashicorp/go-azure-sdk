
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/keyoperationgroup` Documentation

The `keyoperationgroup` SDK allows for interaction with Azure Resource Manager `keyvault` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2025-05-01/keyoperationgroup"
```


### Client Initialization

```go
client := keyoperationgroup.NewKeyOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeyOperationGroupClient.KeysGetVersion`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultKeyVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "keyName", "versionName")

read, err := client.KeysGetVersion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyOperationGroupClient.KeysListVersions`

```go
ctx := context.TODO()
id := commonids.NewKeyVaultKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "keyName")

// alternatively `client.KeysListVersions(ctx, id)` can be used to do batched pagination
items, err := client.KeysListVersionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
