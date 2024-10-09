
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/longtermretentionbackup` Documentation

The `longtermretentionbackup` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2023-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2023-06-01-preview/longtermretentionbackup"
```


### Client Initialization

```go
client := longtermretentionbackup.NewLongTermRetentionBackupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LongTermRetentionBackupClient.FlexibleServerStartLtrBackup`

```go
ctx := context.TODO()
id := longtermretentionbackup.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := longtermretentionbackup.LtrBackupRequest{
	// ...
}


if err := client.FlexibleServerStartLtrBackupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `LongTermRetentionBackupClient.FlexibleServerTriggerLtrPreBackup`

```go
ctx := context.TODO()
id := longtermretentionbackup.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := longtermretentionbackup.BackupRequestBase{
	// ...
}


read, err := client.FlexibleServerTriggerLtrPreBackup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LongTermRetentionBackupClient.LtrBackupOperationsGet`

```go
ctx := context.TODO()
id := longtermretentionbackup.NewLtrBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "ltrBackupOperationName")

read, err := client.LtrBackupOperationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LongTermRetentionBackupClient.LtrBackupOperationsListByServer`

```go
ctx := context.TODO()
id := longtermretentionbackup.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.LtrBackupOperationsListByServer(ctx, id)` can be used to do batched pagination
items, err := client.LtrBackupOperationsListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
