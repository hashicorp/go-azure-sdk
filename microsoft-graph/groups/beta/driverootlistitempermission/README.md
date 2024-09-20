
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitempermission` Documentation

The `driverootlistitempermission` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitempermission"
```


### Client Initialization

```go
client := driverootlistitempermission.NewDriveRootListItemPermissionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveRootListItemPermissionClient.CreateDriveRootListItemPermission`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveID("groupId", "driveId")

payload := driverootlistitempermission.Permission{
	// ...
}


read, err := client.CreateDriveRootListItemPermission(ctx, id, payload, driverootlistitempermission.DefaultCreateDriveRootListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemPermissionClient.DeleteDriveRootListItemPermission`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveIdRootListItemPermissionID("groupId", "driveId", "permissionId")

read, err := client.DeleteDriveRootListItemPermission(ctx, id, driverootlistitempermission.DefaultDeleteDriveRootListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemPermissionClient.GetDriveRootListItemPermission`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveIdRootListItemPermissionID("groupId", "driveId", "permissionId")

read, err := client.GetDriveRootListItemPermission(ctx, id, driverootlistitempermission.DefaultGetDriveRootListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemPermissionClient.GetDriveRootListItemPermissionsCount`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveID("groupId", "driveId")

read, err := client.GetDriveRootListItemPermissionsCount(ctx, id, driverootlistitempermission.DefaultGetDriveRootListItemPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemPermissionClient.ListDriveRootListItemPermissionGrants`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveIdRootListItemPermissionID("groupId", "driveId", "permissionId")

payload := driverootlistitempermission.ListDriveRootListItemPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveRootListItemPermissionGrants(ctx, id, payload, driverootlistitempermission.DefaultListDriveRootListItemPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveRootListItemPermissionGrantsComplete(ctx, id, payload, driverootlistitempermission.DefaultListDriveRootListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveRootListItemPermissionClient.ListDriveRootListItemPermissions`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveID("groupId", "driveId")

// alternatively `client.ListDriveRootListItemPermissions(ctx, id, driverootlistitempermission.DefaultListDriveRootListItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveRootListItemPermissionsComplete(ctx, id, driverootlistitempermission.DefaultListDriveRootListItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveRootListItemPermissionClient.RevokeDriveRootListItemPermissionGrants`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveIdRootListItemPermissionID("groupId", "driveId", "permissionId")

payload := driverootlistitempermission.RevokeDriveRootListItemPermissionGrantsRequest{
	// ...
}


read, err := client.RevokeDriveRootListItemPermissionGrants(ctx, id, payload, driverootlistitempermission.DefaultRevokeDriveRootListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemPermissionClient.UpdateDriveRootListItemPermission`

```go
ctx := context.TODO()
id := driverootlistitempermission.NewGroupIdDriveIdRootListItemPermissionID("groupId", "driveId", "permissionId")

payload := driverootlistitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveRootListItemPermission(ctx, id, payload, driverootlistitempermission.DefaultUpdateDriveRootListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
