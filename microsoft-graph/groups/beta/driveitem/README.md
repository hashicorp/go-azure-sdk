
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitem` Documentation

The `driveitem` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitem"
```


### Client Initialization

```go
client := driveitem.NewDriveItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemClient.AssignDriveItemSensitivityLabel`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.AssignDriveItemSensitivityLabelRequest{
	// ...
}


read, err := client.AssignDriveItemSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CheckinDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.CheckinDriveItemRequest{
	// ...
}


read, err := client.CheckinDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CheckoutDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.CheckoutDriveItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CopyDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.CopyDriveItemRequest{
	// ...
}


read, err := client.CopyDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CreateDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveID("groupIdValue", "driveIdValue")

payload := driveitem.DriveItem{
	// ...
}


read, err := client.CreateDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CreateDriveItemLink`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.CreateDriveItemLinkRequest{
	// ...
}


read, err := client.CreateDriveItemLink(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CreateDriveItemPermanentDelete`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.CreateDriveItemPermanentDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.CreateDriveItemUploadSession`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.CreateDriveItemUploadSessionRequest{
	// ...
}


read, err := client.CreateDriveItemUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.DeleteDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.DeleteDriveItem(ctx, id, driveitem.DefaultDeleteDriveItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.DiscardDriveItemCheckout`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.DiscardDriveItemCheckout(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.ExtractDriveItemSensitivityLabel`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.ExtractDriveItemSensitivityLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.FollowDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.FollowDriveItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.GetDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.GetDriveItem(ctx, id, driveitem.DefaultGetDriveItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.GetDriveItemsCount`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveID("groupIdValue", "driveIdValue")

read, err := client.GetDriveItemsCount(ctx, id, driveitem.DefaultGetDriveItemsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.ListDriveItemInvites`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.ListDriveItemInvitesRequest{
	// ...
}


// alternatively `client.ListDriveItemInvites(ctx, id, payload, driveitem.DefaultListDriveItemInvitesOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemInvitesComplete(ctx, id, payload, driveitem.DefaultListDriveItemInvitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemClient.ListDriveItems`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveID("groupIdValue", "driveIdValue")

// alternatively `client.ListDriveItems(ctx, id, driveitem.DefaultListDriveItemsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemsComplete(ctx, id, driveitem.DefaultListDriveItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemClient.PreviewDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.PreviewDriveItemRequest{
	// ...
}


read, err := client.PreviewDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.RestoreDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.RestoreDriveItemRequest{
	// ...
}


read, err := client.RestoreDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.UnfollowDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.UnfollowDriveItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.UpdateDriveItem`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.DriveItem{
	// ...
}


read, err := client.UpdateDriveItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.ValidateDriveItemPermission`

```go
ctx := context.TODO()
id := driveitem.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitem.ValidateDriveItemPermissionRequest{
	// ...
}


read, err := client.ValidateDriveItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
