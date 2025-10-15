
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backups` Documentation

The `backups` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2025-08-01/backups"
```


### Client Initialization

```go
client := backups.NewBackupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupsClient.Trigger`

```go
ctx := context.TODO()
id := backups.NewProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupFabricName", "protectionContainerName", "protectedItemName")

payload := backups.BackupRequestResource{
	// ...
}


read, err := client.Trigger(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
