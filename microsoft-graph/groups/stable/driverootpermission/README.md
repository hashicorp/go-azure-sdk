
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootpermission` Documentation

The `driverootpermission` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootpermission"
```


### Client Initialization

```go
client := driverootpermission.NewDriveRootPermissionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveRootPermissionClient.CreateDriveRootPermission`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveID("groupId", "driveId")

payload := driverootpermission.Permission{
	// ...
}


read, err := client.CreateDriveRootPermission(ctx, id, payload, driverootpermission.DefaultCreateDriveRootPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootPermissionClient.DeleteDriveRootPermission`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveIdRootPermissionID("groupId", "driveId", "permissionId")

read, err := client.DeleteDriveRootPermission(ctx, id, driverootpermission.DefaultDeleteDriveRootPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootPermissionClient.GetDriveRootPermission`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveIdRootPermissionID("groupId", "driveId", "permissionId")

read, err := client.GetDriveRootPermission(ctx, id, driverootpermission.DefaultGetDriveRootPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootPermissionClient.GetDriveRootPermissionsCount`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveID("groupId", "driveId")

read, err := client.GetDriveRootPermissionsCount(ctx, id, driverootpermission.DefaultGetDriveRootPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootPermissionClient.ListDriveRootPermissionGrants`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveIdRootPermissionID("groupId", "driveId", "permissionId")

payload := driverootpermission.ListDriveRootPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveRootPermissionGrants(ctx, id, payload, driverootpermission.DefaultListDriveRootPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveRootPermissionGrantsComplete(ctx, id, payload, driverootpermission.DefaultListDriveRootPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveRootPermissionClient.ListDriveRootPermissions`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveID("groupId", "driveId")

// alternatively `client.ListDriveRootPermissions(ctx, id, driverootpermission.DefaultListDriveRootPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveRootPermissionsComplete(ctx, id, driverootpermission.DefaultListDriveRootPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveRootPermissionClient.UpdateDriveRootPermission`

```go
ctx := context.TODO()
id := driverootpermission.NewGroupIdDriveIdRootPermissionID("groupId", "driveId", "permissionId")

payload := driverootpermission.Permission{
	// ...
}


read, err := client.UpdateDriveRootPermission(ctx, id, payload, driverootpermission.DefaultUpdateDriveRootPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
