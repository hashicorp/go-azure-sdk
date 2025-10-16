
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretention` Documentation

The `backupslongtermretention` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretention"
```


### Client Initialization

```go
client := backupslongtermretention.NewBackupsLongTermRetentionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupsLongTermRetentionClient.CheckPrerequisites`

```go
ctx := context.TODO()
id := backupslongtermretention.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := backupslongtermretention.BackupRequestBase{
	// ...
}


read, err := client.CheckPrerequisites(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupsLongTermRetentionClient.Get`

```go
ctx := context.TODO()
id := backupslongtermretention.NewLtrBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "ltrBackupOperationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupsLongTermRetentionClient.ListByServer`

```go
ctx := context.TODO()
id := backupslongtermretention.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupsLongTermRetentionClient.Start`

```go
ctx := context.TODO()
id := backupslongtermretention.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := backupslongtermretention.BackupsLongTermRetentionRequest{
	// ...
}


if err := client.StartThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
