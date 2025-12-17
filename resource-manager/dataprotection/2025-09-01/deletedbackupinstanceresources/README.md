
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/deletedbackupinstanceresources` Documentation

The `deletedbackupinstanceresources` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-09-01/deletedbackupinstanceresources"
```


### Client Initialization

```go
client := deletedbackupinstanceresources.NewDeletedBackupInstanceResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedBackupInstanceResourcesClient.DeletedBackupInstancesGet`

```go
ctx := context.TODO()
id := deletedbackupinstanceresources.NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "deletedBackupInstanceName")

read, err := client.DeletedBackupInstancesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedBackupInstanceResourcesClient.DeletedBackupInstancesList`

```go
ctx := context.TODO()
id := deletedbackupinstanceresources.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName")

// alternatively `client.DeletedBackupInstancesList(ctx, id)` can be used to do batched pagination
items, err := client.DeletedBackupInstancesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedBackupInstanceResourcesClient.DeletedBackupInstancesUndelete`

```go
ctx := context.TODO()
id := deletedbackupinstanceresources.NewDeletedBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultName", "deletedBackupInstanceName")

if err := client.DeletedBackupInstancesUndeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```
