
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/recoverypoint` Documentation

The `recoverypoint` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/recoverypoint"
```


### Client Initialization

```go
client := recoverypoint.NewRecoveryPointClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoveryPointClient.Get`

```go
ctx := context.TODO()
id := recoverypoint.NewRecoveryPointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "backupInstanceName", "recoveryPointId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoveryPointClient.List`

```go
ctx := context.TODO()
id := recoverypoint.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "backupInstanceName")

// alternatively `client.List(ctx, id, recoverypoint.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, recoverypoint.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
