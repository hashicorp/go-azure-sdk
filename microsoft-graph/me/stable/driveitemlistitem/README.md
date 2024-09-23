
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/driveitemlistitem` Documentation

The `driveitemlistitem` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/driveitemlistitem"
```


### Client Initialization

```go
client := driveitemlistitem.NewDriveItemListItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemListItemClient.CreateDriveItemListItemLink`

```go
ctx := context.TODO()
id := driveitemlistitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitemlistitem.CreateDriveItemListItemLinkRequest{
	// ...
}


read, err := client.CreateDriveItemListItemLink(ctx, id, payload, driveitemlistitem.DefaultCreateDriveItemListItemLinkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemClient.DeleteDriveItemListItem`

```go
ctx := context.TODO()
id := driveitemlistitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.DeleteDriveItemListItem(ctx, id, driveitemlistitem.DefaultDeleteDriveItemListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemClient.GetDriveItemListItem`

```go
ctx := context.TODO()
id := driveitemlistitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.GetDriveItemListItem(ctx, id, driveitemlistitem.DefaultGetDriveItemListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemClient.UpdateDriveItemListItem`

```go
ctx := context.TODO()
id := driveitemlistitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitemlistitem.ListItem{
	// ...
}


read, err := client.UpdateDriveItemListItem(ctx, id, payload, driveitemlistitem.DefaultUpdateDriveItemListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
