
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-02-01/backupprotectionintent` Documentation

The `backupprotectionintent` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-02-01/backupprotectionintent"
```


### Client Initialization

```go
client := backupprotectionintent.NewBackupProtectionIntentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupProtectionIntentClient.List`

```go
ctx := context.TODO()
id := backupprotectionintent.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

// alternatively `client.List(ctx, id, backupprotectionintent.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupprotectionintent.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
