
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveroot` Documentation

The `driveroot` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveroot"
```


### Client Initialization

```go
client := driveroot.NewDriveRootClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveRootClient.AssignDriveRootSensitivityLabel`

```go
ctx := context.TODO()
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.AssignDriveRootSensitivityLabelRequest{
	// ...
}


read, err := client.AssignDriveRootSensitivityLabel(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.CheckinDriveRootRequest{
	// ...
}


read, err := client.CheckinDriveRoot(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.CheckoutDriveRoot(ctx, id)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.CopyDriveRootRequest{
	// ...
}


read, err := client.CopyDriveRoot(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.CreateDriveRootLinkRequest{
	// ...
}


read, err := client.CreateDriveRootLink(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.CreateDriveRootPermanentDelete(ctx, id)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.CreateDriveRootUploadSessionRequest{
	// ...
}


read, err := client.CreateDriveRootUploadSession(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.DeleteDriveRoot(ctx, id, driveroot.DefaultDeleteDriveRootOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.DiscardDriveRootCheckout`

```go
ctx := context.TODO()
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.DiscardDriveRootCheckout(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootClient.ExtractDriveRootSensitivityLabel`

```go
ctx := context.TODO()
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.ExtractDriveRootSensitivityLabel(ctx, id)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.FollowDriveRoot(ctx, id)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.PreviewDriveRootRequest{
	// ...
}


read, err := client.PreviewDriveRoot(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.RestoreDriveRootRequest{
	// ...
}


read, err := client.RestoreDriveRoot(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.UnfollowDriveRoot(ctx, id)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.DriveItem{
	// ...
}


read, err := client.UpdateDriveRoot(ctx, id, payload)
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
id := driveroot.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := driveroot.ValidateDriveRootPermissionRequest{
	// ...
}


read, err := client.ValidateDriveRootPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
