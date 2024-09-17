
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/syncagents` Documentation

The `syncagents` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/syncagents"
```


### Client Initialization

```go
client := syncagents.NewSyncAgentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SyncAgentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := syncagents.NewSyncAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "syncAgentValue")

payload := syncagents.SyncAgent{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SyncAgentsClient.Delete`

```go
ctx := context.TODO()
id := syncagents.NewSyncAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "syncAgentValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SyncAgentsClient.GenerateKey`

```go
ctx := context.TODO()
id := syncagents.NewSyncAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "syncAgentValue")

read, err := client.GenerateKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncAgentsClient.Get`

```go
ctx := context.TODO()
id := syncagents.NewSyncAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "syncAgentValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SyncAgentsClient.ListByServer`

```go
ctx := context.TODO()
id := commonids.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SyncAgentsClient.ListLinkedDatabases`

```go
ctx := context.TODO()
id := syncagents.NewSyncAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "syncAgentValue")

// alternatively `client.ListLinkedDatabases(ctx, id)` can be used to do batched pagination
items, err := client.ListLinkedDatabasesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
