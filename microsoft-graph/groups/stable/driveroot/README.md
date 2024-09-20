
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveroot` Documentation

The `driveroot` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveroot"
```


### Client Initialization

```go
client := driveroot.NewDriveRootClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveRootClient.AssignDriveRootSensitivityLabel`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.AssignDriveRootSensitivityLabelRequest{
	// ...
}


read, err := client.AssignDriveRootSensitivityLabel(ctx, id, payload, driveroot.DefaultAssignDriveRootSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CheckinDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.CheckinDriveRootRequest{
	// ...
}


read, err := client.CheckinDriveRoot(ctx, id, payload, driveroot.DefaultCheckinDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CheckoutDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.CheckoutDriveRoot(ctx, id, driveroot.DefaultCheckoutDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CopyDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.CopyDriveRootRequest{
	// ...
}


read, err := client.CopyDriveRoot(ctx, id, payload, driveroot.DefaultCopyDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CreateDriveRootLink`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.CreateDriveRootLinkRequest{
	// ...
}


read, err := client.CreateDriveRootLink(ctx, id, payload, driveroot.DefaultCreateDriveRootLinkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CreateDriveRootPermanentDelete`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.CreateDriveRootPermanentDelete(ctx, id, driveroot.DefaultCreateDriveRootPermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.CreateDriveRootUploadSession`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.CreateDriveRootUploadSessionRequest{
	// ...
}


read, err := client.CreateDriveRootUploadSession(ctx, id, payload, driveroot.DefaultCreateDriveRootUploadSessionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.DeleteDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.DeleteDriveRoot(ctx, id, driveroot.DefaultDeleteDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.ExtractDriveRootSensitivityLabels`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.ExtractDriveRootSensitivityLabels(ctx, id, driveroot.DefaultExtractDriveRootSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.FollowDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.FollowDriveRoot(ctx, id, driveroot.DefaultFollowDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.GetDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.GetDriveRoot(ctx, id, driveroot.DefaultGetDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.ListDriveRootInvites`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.ListDriveRootInvitesRequest{
	// ...
}


// alternatively `client.ListDriveRootInvites(ctx, id, payload, driveroot.DefaultListDriveRootInvitesOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveRootInvitesComplete(ctx, id, payload, driveroot.DefaultListDriveRootInvitesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveRootClient.PreviewDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.PreviewDriveRootRequest{
	// ...
}


read, err := client.PreviewDriveRoot(ctx, id, payload, driveroot.DefaultPreviewDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.RestoreDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.RestoreDriveRootRequest{
	// ...
}


read, err := client.RestoreDriveRoot(ctx, id, payload, driveroot.DefaultRestoreDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.UnfollowDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

read, err := client.UnfollowDriveRoot(ctx, id, driveroot.DefaultUnfollowDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.UpdateDriveRoot`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.DriveItem{
	// ...
}


read, err := client.UpdateDriveRoot(ctx, id, payload, driveroot.DefaultUpdateDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.ValidateDriveRootPermission`

```go
ctx := context.TODO()
id := driveroot.NewGroupIdDriveID("groupId", "driveId")

payload := driveroot.ValidateDriveRootPermissionRequest{
	// ...
}


read, err := client.ValidateDriveRootPermission(ctx, id, payload, driveroot.DefaultValidateDriveRootPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
