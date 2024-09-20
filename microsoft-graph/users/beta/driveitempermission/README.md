
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitempermission` Documentation

The `driveitempermission` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitempermission"
```


### Client Initialization

```go
client := driveitempermission.NewDriveItemPermissionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemPermissionClient.CreateDriveItemPermission`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemID("userId", "driveId", "driveItemId")

payload := driveitempermission.Permission{
	// ...
}


read, err := client.CreateDriveItemPermission(ctx, id, payload, driveitempermission.DefaultCreateDriveItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemPermissionClient.DeleteDriveItemPermission`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemIdPermissionID("userId", "driveId", "driveItemId", "permissionId")

read, err := client.DeleteDriveItemPermission(ctx, id, driveitempermission.DefaultDeleteDriveItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemPermissionClient.GetDriveItemPermission`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemIdPermissionID("userId", "driveId", "driveItemId", "permissionId")

read, err := client.GetDriveItemPermission(ctx, id, driveitempermission.DefaultGetDriveItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemPermissionClient.GetDriveItemPermissionsCount`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemID("userId", "driveId", "driveItemId")

read, err := client.GetDriveItemPermissionsCount(ctx, id, driveitempermission.DefaultGetDriveItemPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemPermissionClient.ListDriveItemPermissionGrants`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemIdPermissionID("userId", "driveId", "driveItemId", "permissionId")

payload := driveitempermission.ListDriveItemPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveItemPermissionGrants(ctx, id, payload, driveitempermission.DefaultListDriveItemPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemPermissionGrantsComplete(ctx, id, payload, driveitempermission.DefaultListDriveItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemPermissionClient.ListDriveItemPermissions`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemID("userId", "driveId", "driveItemId")

// alternatively `client.ListDriveItemPermissions(ctx, id, driveitempermission.DefaultListDriveItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemPermissionsComplete(ctx, id, driveitempermission.DefaultListDriveItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemPermissionClient.RevokeDriveItemPermissionGrants`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemIdPermissionID("userId", "driveId", "driveItemId", "permissionId")

payload := driveitempermission.RevokeDriveItemPermissionGrantsRequest{
	// ...
}


read, err := client.RevokeDriveItemPermissionGrants(ctx, id, payload, driveitempermission.DefaultRevokeDriveItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemPermissionClient.UpdateDriveItemPermission`

```go
ctx := context.TODO()
id := driveitempermission.NewUserIdDriveIdItemIdPermissionID("userId", "driveId", "driveItemId", "permissionId")

payload := driveitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveItemPermission(ctx, id, payload, driveitempermission.DefaultUpdateDriveItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
