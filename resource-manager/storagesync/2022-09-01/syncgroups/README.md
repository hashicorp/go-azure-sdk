
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/syncgroups` Documentation

The `syncgroups` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/syncgroups"
```


### Client Initialization

```go
client := syncgroups.NewSyncGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncGroupsClient.Create`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName")

payload := syncgroups.SyncGroupCreateParameters{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.Delete`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.Get`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.ListByStorageSyncService`

```go
ctx := context.TODO()
id := syncgroups.NewStorageSyncServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName")

// alternatively `client.ListByStorageSyncService(ctx, id)` can be used to do batched pagination
items, err := client.ListByStorageSyncServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
