
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/backupresourcestorageconfigsnoncrr` Documentation

The `backupresourcestorageconfigsnoncrr` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-06-01/backupresourcestorageconfigsnoncrr"
```


### Client Initialization

```go
client := backupresourcestorageconfigsnoncrr.NewBackupResourceStorageConfigsNonCRRClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupResourceStorageConfigsNonCRRClient.Get`

```go
ctx := context.TODO()
id := backupresourcestorageconfigsnoncrr.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupResourceStorageConfigsNonCRRClient.Patch`

```go
ctx := context.TODO()
id := backupresourcestorageconfigsnoncrr.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := backupresourcestorageconfigsnoncrr.BackupResourceConfigResource{
	// ...
}


read, err := client.Patch(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupResourceStorageConfigsNonCRRClient.Update`

```go
ctx := context.TODO()
id := backupresourcestorageconfigsnoncrr.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := backupresourcestorageconfigsnoncrr.BackupResourceConfigResource{
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
