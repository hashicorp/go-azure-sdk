
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/watchlistitems` Documentation

The `watchlistitems` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2021-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/watchlistitems"
```


### Client Initialization

```go
client := watchlistitems.NewWatchlistItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WatchlistItemsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := watchlistitems.NewWatchlistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "watchlistAliasValue", "watchlistItemIdValue")

payload := watchlistitems.WatchlistItem{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WatchlistItemsClient.Delete`

```go
ctx := context.TODO()
id := watchlistitems.NewWatchlistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "watchlistAliasValue", "watchlistItemIdValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WatchlistItemsClient.Get`

```go
ctx := context.TODO()
id := watchlistitems.NewWatchlistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "watchlistAliasValue", "watchlistItemIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WatchlistItemsClient.List`

```go
ctx := context.TODO()
id := watchlistitems.NewWatchlistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "watchlistAliasValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
