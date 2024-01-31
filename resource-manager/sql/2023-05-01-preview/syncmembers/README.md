
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/syncmembers` Documentation

The `syncmembers` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/syncmembers"
```


### Client Initialization

```go
client := syncmembers.NewSyncMembersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncMembersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

payload := syncmembers.SyncMember{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SyncMembersClient.Delete`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncMembersClient.Get`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncMembersClient.ListBySyncGroup`

```go
ctx := context.TODO()
id := syncmembers.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue")

// alternatively `client.ListBySyncGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListBySyncGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncMembersClient.ListMemberSchemas`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

// alternatively `client.ListMemberSchemas(ctx, id)` can be used to do batched pagination
items, err := client.ListMemberSchemasComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncMembersClient.RefreshMemberSchema`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

if err := client.RefreshMemberSchemaThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncMembersClient.Update`

```go
ctx := context.TODO()
id := syncmembers.NewSyncMemberID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "databaseValue", "syncGroupValue", "syncMemberValue")

payload := syncmembers.SyncMember{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
