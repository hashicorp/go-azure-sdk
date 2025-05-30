
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupengines` Documentation

The `backupengines` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupengines"
```


### Client Initialization

```go
client := backupengines.NewBackupEnginesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupEnginesClient.Get`

```go
ctx := context.TODO()
id := backupengines.NewBackupEngineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupEngineName")

read, err := client.Get(ctx, id, backupengines.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupEnginesClient.List`

```go
ctx := context.TODO()
id := backupengines.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.List(ctx, id, backupengines.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupengines.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
