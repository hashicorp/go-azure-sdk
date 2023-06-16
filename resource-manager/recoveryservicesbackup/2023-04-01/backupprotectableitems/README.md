
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-04-01/backupprotectableitems` Documentation

The `backupprotectableitems` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-04-01/backupprotectableitems"
```


### Client Initialization

```go
client := backupprotectableitems.NewBackupProtectableItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupProtectableItemsClient.List`

```go
ctx := context.TODO()
id := backupprotectableitems.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, backupprotectableitems.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupprotectableitems.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
