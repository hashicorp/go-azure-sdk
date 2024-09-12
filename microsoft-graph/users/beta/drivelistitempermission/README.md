
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitempermission` Documentation

The `drivelistitempermission` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitempermission"
```


### Client Initialization

```go
client := drivelistitempermission.NewDriveListItemPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveListItemPermissionClient.CreateDriveListItemPermission`

```go
ctx := context.TODO()
id := drivelistitempermission.NewUserIdDriveIdListItemID("userIdValue", "driveIdValue", "listItemIdValue")

payload := drivelistitempermission.Permission{
	// ...
}


read, err := client.CreateDriveListItemPermission(ctx, id, payload)
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
id := drivelistitempermission.NewUserIdDriveIdListItemIdPermissionID("userIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

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
id := drivelistitempermission.NewUserIdDriveIdListItemIdPermissionID("userIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

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
id := drivelistitempermission.NewUserIdDriveIdListItemID("userIdValue", "driveIdValue", "listItemIdValue")

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
id := drivelistitempermission.NewUserIdDriveIdListItemIdPermissionID("userIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

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
id := drivelistitempermission.NewUserIdDriveIdListItemID("userIdValue", "driveIdValue", "listItemIdValue")

// alternatively `client.ListDriveListItemPermissions(ctx, id, drivelistitempermission.DefaultListDriveListItemPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListItemPermissionsComplete(ctx, id, drivelistitempermission.DefaultListDriveListItemPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListItemPermissionClient.RevokeDriveListItemPermissionGrant`

```go
ctx := context.TODO()
id := drivelistitempermission.NewUserIdDriveIdListItemIdPermissionID("userIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

payload := drivelistitempermission.RevokeDriveListItemPermissionGrantRequest{
	// ...
}


read, err := client.RevokeDriveListItemPermissionGrant(ctx, id, payload)
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
id := drivelistitempermission.NewUserIdDriveIdListItemIdPermissionID("userIdValue", "driveIdValue", "listItemIdValue", "permissionIdValue")

payload := drivelistitempermission.Permission{
	// ...
}


read, err := client.UpdateDriveListItemPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
