
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitems` Documentation

The `backupitems` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/backupitems"
```


### Client Initialization

```go
client := backupitems.NewBackupItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupItemsClient.WebAppsDeleteBackup`

```go
ctx := context.TODO()
id := backupitems.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "backupId")

read, err := client.WebAppsDeleteBackup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemsClient.WebAppsGetBackupStatus`

```go
ctx := context.TODO()
id := backupitems.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "backupId")

read, err := client.WebAppsGetBackupStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemsClient.WebAppsListBackupStatusSecrets`

```go
ctx := context.TODO()
id := backupitems.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "backupId")

payload := backupitems.BackupRequest{
	// ...
}


read, err := client.WebAppsListBackupStatusSecrets(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupItemsClient.WebAppsListBackups`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListBackups(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListBackupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupItemsClient.WebAppsRestore`

```go
ctx := context.TODO()
id := backupitems.NewBackupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "backupId")

payload := backupitems.RestoreRequest{
	// ...
}


if err := client.WebAppsRestoreThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
