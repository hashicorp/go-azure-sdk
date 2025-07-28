
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/vaultextendedinforesources` Documentation

The `vaultextendedinforesources` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/vaultextendedinforesources"
```


### Client Initialization

```go
client := vaultextendedinforesources.NewVaultExtendedInfoResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VaultExtendedInfoResourcesClient.VaultExtendedInfoCreateOrUpdate`

```go
ctx := context.TODO()
id := vaultextendedinforesources.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := vaultextendedinforesources.VaultExtendedInfoResource{
	// ...
}


read, err := client.VaultExtendedInfoCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VaultExtendedInfoResourcesClient.VaultExtendedInfoGet`

```go
ctx := context.TODO()
id := vaultextendedinforesources.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.VaultExtendedInfoGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VaultExtendedInfoResourcesClient.VaultExtendedInfoUpdate`

```go
ctx := context.TODO()
id := vaultextendedinforesources.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := vaultextendedinforesources.VaultExtendedInfoResource{
	// ...
}


read, err := client.VaultExtendedInfoUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
