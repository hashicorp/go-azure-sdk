
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitemoperationgroup` Documentation

The `backupitemoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitemoperationgroup"
```


### Client Initialization

```go
client := backupitemoperationgroup.NewBackupItemOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupItemOperationGroupClient.WebAppsDeleteBackupSlot`

```go
ctx := context.TODO()
id := backupitemoperationgroup.NewSlotBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "backupId")

read, err := client.WebAppsDeleteBackupSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemOperationGroupClient.WebAppsGetBackupStatusSlot`

```go
ctx := context.TODO()
id := backupitemoperationgroup.NewSlotBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "backupId")

read, err := client.WebAppsGetBackupStatusSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemOperationGroupClient.WebAppsListBackupStatusSecretsSlot`

```go
ctx := context.TODO()
id := backupitemoperationgroup.NewSlotBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "backupId")

payload := backupitemoperationgroup.BackupRequest{
	// ...
}


read, err := client.WebAppsListBackupStatusSecretsSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemOperationGroupClient.WebAppsListBackupsSlot`

```go
ctx := context.TODO()
id := backupitemoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListBackupsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListBackupsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupItemOperationGroupClient.WebAppsRestoreSlot`

```go
ctx := context.TODO()
id := backupitemoperationgroup.NewSlotBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "backupId")

payload := backupitemoperationgroup.RestoreRequest{
	// ...
}


if err := client.WebAppsRestoreSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
