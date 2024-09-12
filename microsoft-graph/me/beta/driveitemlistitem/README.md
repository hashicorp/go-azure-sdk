
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/driveitemlistitem` Documentation

The `driveitemlistitem` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/driveitemlistitem"
```


### Client Initialization

```go
client := driveitemlistitem.NewDriveItemListItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemListItemClient.CreateDriveItemListItemLink`

```go
ctx := context.TODO()
id := driveitemlistitem.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

payload := driveitemlistitem.CreateDriveItemListItemLinkRequest{
	// ...
}


read, err := client.CreateDriveItemListItemLink(ctx, id, payload)
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
id := driveitemlistitem.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

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
id := driveitemlistitem.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

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
id := driveitemlistitem.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

payload := driveitemlistitem.ListItem{
	// ...
}


read, err := client.UpdateDriveItemListItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
