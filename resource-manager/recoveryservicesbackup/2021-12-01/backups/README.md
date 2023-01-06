
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backups` Documentation

The `backups` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2021-12-01/backups"
```


### Client Initialization

```go
client := backups.NewBackupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupsClient.Trigger`

```go
ctx := context.TODO()
id := backups.NewProtectedItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "fabricValue", "containerValue", "protectedItemValue")

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
