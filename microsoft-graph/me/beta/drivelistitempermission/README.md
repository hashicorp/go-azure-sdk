
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/drivelistitempermission` Documentation

The `drivelistitempermission` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/drivelistitempermission"
```


### Client Initialization

```go
client := drivelistitempermission.NewDriveListItemPermissionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveListItemPermissionClient.CreateDriveListItemPermission`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemID("driveId", "listItemId")

payload := drivelistitempermission.Permission{
	// ...
}


read, err := client.CreateDriveListItemPermission(ctx, id, payload, drivelistitempermission.DefaultCreateDriveListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemPermissionClient.DeleteDriveListItemPermission`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemIdPermissionID("driveId", "listItemId", "permissionId")

read, err := client.DeleteDriveListItemPermission(ctx, id, drivelistitempermission.DefaultDeleteDriveListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemPermissionClient.GetDriveListItemPermission`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemIdPermissionID("driveId", "listItemId", "permissionId")

read, err := client.GetDriveListItemPermission(ctx, id, drivelistitempermission.DefaultGetDriveListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemPermissionClient.GetDriveListItemPermissionsCount`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemID("driveId", "listItemId")

read, err := client.GetDriveListItemPermissionsCount(ctx, id, drivelistitempermission.DefaultGetDriveListItemPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemPermissionClient.ListDriveListItemPermissionGrants`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemIdPermissionID("driveId", "listItemId", "permissionId")

payload := drivelistitempermission.ListDriveListItemPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveListItemPermissionGrants(ctx, id, payload, drivelistitempermission.DefaultListDriveListItemPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListItemPermissionGrantsComplete(ctx, id, payload, drivelistitempermission.DefaultListDriveListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListItemPermissionClient.ListDriveListItemPermissions`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemID("driveId", "listItemId")

// alternatively `client.ListDriveListItemPermissions(ctx, id, drivelistitempermission.DefaultListDriveListItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListItemPermissionsComplete(ctx, id, drivelistitempermission.DefaultListDriveListItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListItemPermissionClient.RevokeDriveListItemPermissionGrants`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemIdPermissionID("driveId", "listItemId", "permissionId")

payload := drivelistitempermission.RevokeDriveListItemPermissionGrantsRequest{
	// ...
}


read, err := client.RevokeDriveListItemPermissionGrants(ctx, id, payload, drivelistitempermission.DefaultRevokeDriveListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListItemPermissionClient.UpdateDriveListItemPermission`

```go
ctx := context.TODO()
id := drivelistitempermission.NewMeDriveIdListItemIdPermissionID("driveId", "listItemId", "permissionId")

payload := drivelistitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveListItemPermission(ctx, id, payload, drivelistitempermission.DefaultUpdateDriveListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
