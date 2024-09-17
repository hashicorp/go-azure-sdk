
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/backupworkloaditems` Documentation

The `backupworkloaditems` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/backupworkloaditems"
```


### Client Initialization

```go
client := backupworkloaditems.NewBackupWorkloadItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupWorkloadItemsClient.List`

```go
ctx := context.TODO()
id := backupworkloaditems.NewProtectionContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupFabricValue", "protectionContainerValue")

// alternatively `client.List(ctx, id, backupworkloaditems.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, backupworkloaditems.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
