
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/bookmarks` Documentation

The `bookmarks` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/bookmarks"
```


### Client Initialization

```go
client := bookmarks.NewBookmarksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BookmarksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := bookmarks.NewBookmarkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId")

payload := bookmarks.Bookmark{
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


### Example Usage: `BookmarksClient.Delete`

```go
ctx := context.TODO()
id := bookmarks.NewBookmarkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BookmarksClient.Get`

```go
ctx := context.TODO()
id := bookmarks.NewBookmarkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BookmarksClient.List`

```go
ctx := context.TODO()
id := bookmarks.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
