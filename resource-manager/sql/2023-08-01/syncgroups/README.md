
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/syncgroups` Documentation

The `syncgroups` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/syncgroups"
```


### Client Initialization

```go
client := syncgroups.NewSyncGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncGroupsClient.CancelSync`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

read, err := client.CancelSync(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

payload := syncgroups.SyncGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.Delete`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.Get`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.ListByDatabase`

```go
ctx := context.TODO()
id := commonids.NewSqlDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName")

// alternatively `client.ListByDatabase(ctx, id)` can be used to do batched pagination
items, err := client.ListByDatabaseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListHubSchemas`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

// alternatively `client.ListHubSchemas(ctx, id)` can be used to do batched pagination
items, err := client.ListHubSchemasComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListLogs`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

// alternatively `client.ListLogs(ctx, id, syncgroups.DefaultListLogsOperationOptions())` can be used to do batched pagination
items, err := client.ListLogsComplete(ctx, id, syncgroups.DefaultListLogsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.ListSyncDatabaseIds`

```go
ctx := context.TODO()
id := syncgroups.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListSyncDatabaseIds(ctx, id)` can be used to do batched pagination
items, err := client.ListSyncDatabaseIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncGroupsClient.RefreshHubSchema`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

if err := client.RefreshHubSchemaThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncGroupsClient.TriggerSync`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

read, err := client.TriggerSync(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncGroupsClient.Update`

```go
ctx := context.TODO()
id := syncgroups.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "databaseName", "syncGroupName")

payload := syncgroups.SyncGroup{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
