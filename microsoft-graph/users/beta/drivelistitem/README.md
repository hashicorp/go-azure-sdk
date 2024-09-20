
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitem` Documentation

The `drivelistitem` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitem"
```


### Client Initialization

```go
client := drivelistitem.NewDriveListItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveListItemClient.CreateDriveListItem`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveID("userId", "driveId")

payload := drivelistitem.ListItem{
	// ...
}


read, err := client.CreateDriveListItem(ctx, id, payload, drivelistitem.DefaultCreateDriveListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemClient.CreateDriveListItemLink`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveIdListItemID("userId", "driveId", "listItemId")

payload := drivelistitem.CreateDriveListItemLinkRequest{
	// ...
}


read, err := client.CreateDriveListItemLink(ctx, id, payload, drivelistitem.DefaultCreateDriveListItemLinkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemClient.DeleteDriveListItem`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveIdListItemID("userId", "driveId", "listItemId")

read, err := client.DeleteDriveListItem(ctx, id, drivelistitem.DefaultDeleteDriveListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemClient.GetDriveListItem`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveIdListItemID("userId", "driveId", "listItemId")

read, err := client.GetDriveListItem(ctx, id, drivelistitem.DefaultGetDriveListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemClient.ListDriveListItems`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveID("userId", "driveId")

// alternatively `client.ListDriveListItems(ctx, id, drivelistitem.DefaultListDriveListItemsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListItemsComplete(ctx, id, drivelistitem.DefaultListDriveListItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListItemClient.UpdateDriveListItem`

```go
ctx := context.TODO()
id := drivelistitem.NewUserIdDriveIdListItemID("userId", "driveId", "listItemId")

payload := drivelistitem.ListItem{
	// ...
}


read, err := client.UpdateDriveListItem(ctx, id, payload, drivelistitem.DefaultUpdateDriveListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
