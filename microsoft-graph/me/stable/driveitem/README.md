
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/driveitem` Documentation

The `driveitem` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/driveitem"
```


### Client Initialization

```go
client := driveitem.NewDriveItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemClient.AssignDriveItemSensitivityLabel`

```go
ctx := context.TODO()
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.AssignDriveItemSensitivityLabelRequest{
	// ...
}


read, err := client.AssignDriveItemSensitivityLabel(ctx, id, payload, driveitem.DefaultAssignDriveItemSensitivityLabelOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.CheckinDriveItemRequest{
	// ...
}


read, err := client.CheckinDriveItem(ctx, id, payload, driveitem.DefaultCheckinDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.CheckoutDriveItem(ctx, id, driveitem.DefaultCheckoutDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.CopyDriveItemRequest{
	// ...
}


read, err := client.CopyDriveItem(ctx, id, payload, driveitem.DefaultCopyDriveItemOperationOptions())
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
id := driveitem.NewMeDriveID("driveId")

payload := driveitem.DriveItem{
	// ...
}


read, err := client.CreateDriveItem(ctx, id, payload, driveitem.DefaultCreateDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.CreateDriveItemLinkRequest{
	// ...
}


read, err := client.CreateDriveItemLink(ctx, id, payload, driveitem.DefaultCreateDriveItemLinkOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.CreateDriveItemPermanentDelete(ctx, id, driveitem.DefaultCreateDriveItemPermanentDeleteOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.CreateDriveItemUploadSessionRequest{
	// ...
}


read, err := client.CreateDriveItemUploadSession(ctx, id, payload, driveitem.DefaultCreateDriveItemUploadSessionOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.DeleteDriveItem(ctx, id, driveitem.DefaultDeleteDriveItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemClient.ExtractDriveItemSensitivityLabels`

```go
ctx := context.TODO()
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.ExtractDriveItemSensitivityLabels(ctx, id, driveitem.DefaultExtractDriveItemSensitivityLabelsOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.FollowDriveItem(ctx, id, driveitem.DefaultFollowDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

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
id := driveitem.NewMeDriveID("driveId")

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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

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
id := driveitem.NewMeDriveID("driveId")

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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.PreviewDriveItemRequest{
	// ...
}


read, err := client.PreviewDriveItem(ctx, id, payload, driveitem.DefaultPreviewDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.RestoreDriveItemRequest{
	// ...
}


read, err := client.RestoreDriveItem(ctx, id, payload, driveitem.DefaultRestoreDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

read, err := client.UnfollowDriveItem(ctx, id, driveitem.DefaultUnfollowDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.DriveItem{
	// ...
}


read, err := client.UpdateDriveItem(ctx, id, payload, driveitem.DefaultUpdateDriveItemOperationOptions())
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
id := driveitem.NewMeDriveIdItemID("driveId", "driveItemId")

payload := driveitem.ValidateDriveItemPermissionRequest{
	// ...
}


read, err := client.ValidateDriveItemPermission(ctx, id, payload, driveitem.DefaultValidateDriveItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
