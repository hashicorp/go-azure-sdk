
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2024-04-01/vaultextendedinfo` Documentation

The `vaultextendedinfo` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2024-04-01/vaultextendedinfo"
```


### Client Initialization

```go
client := vaultextendedinfo.NewVaultExtendedInfoClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VaultExtendedInfoClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := vaultextendedinfo.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := vaultextendedinfo.VaultExtendedInfoResource{
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


### Example Usage: `VaultExtendedInfoClient.Get`

```go
ctx := context.TODO()
id := vaultextendedinfo.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VaultExtendedInfoClient.Update`

```go
ctx := context.TODO()
id := vaultextendedinfo.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := vaultextendedinfo.VaultExtendedInfoResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
