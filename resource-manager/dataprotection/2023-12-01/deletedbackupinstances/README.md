
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-12-01/deletedbackupinstances` Documentation

The `deletedbackupinstances` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-12-01/deletedbackupinstances"
```


### Client Initialization

```go
client := deletedbackupinstances.NewDeletedBackupInstancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedBackupInstancesClient.Get`

```go
ctx := context.TODO()
id := deletedbackupinstances.NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupInstanceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedBackupInstancesClient.List`

```go
ctx := context.TODO()
id := deletedbackupinstances.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedBackupInstancesClient.Undelete`

```go
ctx := context.TODO()
id := deletedbackupinstances.NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "backupInstanceName")

if err := client.UndeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```
