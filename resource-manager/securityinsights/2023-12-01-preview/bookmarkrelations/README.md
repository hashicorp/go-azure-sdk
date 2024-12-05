
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/bookmarkrelations` Documentation

The `bookmarkrelations` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/bookmarkrelations"
```


### Client Initialization

```go
client := bookmarkrelations.NewBookmarkRelationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BookmarkRelationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := bookmarkrelations.NewBookmarkRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId", "relationName")

payload := bookmarkrelations.Relation{
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


### Example Usage: `BookmarkRelationsClient.Delete`

```go
ctx := context.TODO()
id := bookmarkrelations.NewBookmarkRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId", "relationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BookmarkRelationsClient.Get`

```go
ctx := context.TODO()
id := bookmarkrelations.NewBookmarkRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId", "relationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BookmarkRelationsClient.List`

```go
ctx := context.TODO()
id := bookmarkrelations.NewBookmarkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "bookmarkId")

// alternatively `client.List(ctx, id, bookmarkrelations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, bookmarkrelations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
