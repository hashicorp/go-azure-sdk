
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretentionoperations` Documentation

The `backupslongtermretentionoperations` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2025-08-01/backupslongtermretentionoperations"
```


### Client Initialization

```go
client := backupslongtermretentionoperations.NewBackupsLongTermRetentionOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupsLongTermRetentionOperationsClient.BackupsLongTermRetentionGet`

```go
ctx := context.TODO()
id := backupslongtermretentionoperations.NewLtrBackupOperationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "ltrBackupOperationName")

read, err := client.BackupsLongTermRetentionGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupsLongTermRetentionOperationsClient.BackupsLongTermRetentionListByServer`

```go
ctx := context.TODO()
id := backupslongtermretentionoperations.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.BackupsLongTermRetentionListByServer(ctx, id)` can be used to do batched pagination
items, err := client.BackupsLongTermRetentionListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
