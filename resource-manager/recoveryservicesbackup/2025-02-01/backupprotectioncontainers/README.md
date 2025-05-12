
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupprotectioncontainers` Documentation

The `backupprotectioncontainers` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-02-01/backupprotectioncontainers"
```


### Client Initialization

```go
client := backupprotectioncontainers.NewBackupProtectionContainersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupProtectionContainersClient.List`

```go
ctx := context.TODO()
id := backupprotectioncontainers.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.List(ctx, id, backupprotectioncontainers.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupprotectioncontainers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
