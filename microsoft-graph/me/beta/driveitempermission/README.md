
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/driveitempermission` Documentation

The `driveitempermission` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/driveitempermission"
```


### Client Initialization

```go
client := driveitempermission.NewDriveItemPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemPermissionClient.CreateDriveItemPermission`

```go
ctx := context.TODO()
id := driveitempermission.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

payload := driveitempermission.Permission{
	// ...
}


read, err := client.CreateDriveItemPermission(ctx, id, payload)
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
id := driveitempermission.NewMeDriveIdItemIdPermissionID("driveIdValue", "driveItemIdValue", "permissionIdValue")

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
id := driveitempermission.NewMeDriveIdItemIdPermissionID("driveIdValue", "driveItemIdValue", "permissionIdValue")

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
id := driveitempermission.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

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
id := driveitempermission.NewMeDriveIdItemIdPermissionID("driveIdValue", "driveItemIdValue", "permissionIdValue")

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
id := driveitempermission.NewMeDriveIdItemID("driveIdValue", "driveItemIdValue")

// alternatively `client.ListDriveItemPermissions(ctx, id, driveitempermission.DefaultListDriveItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemPermissionsComplete(ctx, id, driveitempermission.DefaultListDriveItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemPermissionClient.RevokeDriveItemPermissionGrant`

```go
ctx := context.TODO()
id := driveitempermission.NewMeDriveIdItemIdPermissionID("driveIdValue", "driveItemIdValue", "permissionIdValue")

payload := driveitempermission.RevokeDriveItemPermissionGrantRequest{
	// ...
}


read, err := client.RevokeDriveItemPermissionGrant(ctx, id, payload)
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
id := driveitempermission.NewMeDriveIdItemIdPermissionID("driveIdValue", "driveItemIdValue", "permissionIdValue")

payload := driveitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
