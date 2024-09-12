
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitempermission` Documentation

The `driveitemlistitempermission` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitempermission"
```


### Client Initialization

```go
client := driveitemlistitempermission.NewDriveItemListItemPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveItemListItemPermissionClient.CreateDriveItemListItemPermission`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

payload := driveitemlistitempermission.Permission{
	// ...
}


read, err := client.CreateDriveItemListItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemPermissionClient.DeleteDriveItemListItemPermission`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemIdListItemPermissionID("groupIdValue", "driveIdValue", "driveItemIdValue", "permissionIdValue")

read, err := client.DeleteDriveItemListItemPermission(ctx, id, driveitemlistitempermission.DefaultDeleteDriveItemListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemPermissionClient.GetDriveItemListItemPermission`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemIdListItemPermissionID("groupIdValue", "driveIdValue", "driveItemIdValue", "permissionIdValue")

read, err := client.GetDriveItemListItemPermission(ctx, id, driveitemlistitempermission.DefaultGetDriveItemListItemPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemPermissionClient.GetDriveItemListItemPermissionsCount`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

read, err := client.GetDriveItemListItemPermissionsCount(ctx, id, driveitemlistitempermission.DefaultGetDriveItemListItemPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemPermissionClient.ListDriveItemListItemPermissionGrants`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemIdListItemPermissionID("groupIdValue", "driveIdValue", "driveItemIdValue", "permissionIdValue")

payload := driveitemlistitempermission.ListDriveItemListItemPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveItemListItemPermissionGrants(ctx, id, payload, driveitemlistitempermission.DefaultListDriveItemListItemPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemListItemPermissionGrantsComplete(ctx, id, payload, driveitemlistitempermission.DefaultListDriveItemListItemPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemListItemPermissionClient.ListDriveItemListItemPermissions`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemID("groupIdValue", "driveIdValue", "driveItemIdValue")

// alternatively `client.ListDriveItemListItemPermissions(ctx, id, driveitemlistitempermission.DefaultListDriveItemListItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveItemListItemPermissionsComplete(ctx, id, driveitemlistitempermission.DefaultListDriveItemListItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveItemListItemPermissionClient.RevokeDriveItemListItemPermissionGrant`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemIdListItemPermissionID("groupIdValue", "driveIdValue", "driveItemIdValue", "permissionIdValue")

payload := driveitemlistitempermission.RevokeDriveItemListItemPermissionGrantRequest{
	// ...
}


read, err := client.RevokeDriveItemListItemPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveItemListItemPermissionClient.UpdateDriveItemListItemPermission`

```go
ctx := context.TODO()
id := driveitemlistitempermission.NewGroupIdDriveIdItemIdListItemPermissionID("groupIdValue", "driveIdValue", "driveItemIdValue", "permissionIdValue")

payload := driveitemlistitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveItemListItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
