
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistpermission` Documentation

The `drivelistpermission` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistpermission"
```


### Client Initialization

```go
client := drivelistpermission.NewDriveListPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveListPermissionClient.CreateDriveListPermission`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveID("userIdValue", "driveIdValue")

payload := drivelistpermission.Permission{
	// ...
}


read, err := client.CreateDriveListPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListPermissionClient.DeleteDriveListPermission`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveIdListPermissionID("userIdValue", "driveIdValue", "permissionIdValue")

read, err := client.DeleteDriveListPermission(ctx, id, drivelistpermission.DefaultDeleteDriveListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListPermissionClient.GetDriveListPermission`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveIdListPermissionID("userIdValue", "driveIdValue", "permissionIdValue")

read, err := client.GetDriveListPermission(ctx, id, drivelistpermission.DefaultGetDriveListPermissionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListPermissionClient.GetDriveListPermissionsCount`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveID("userIdValue", "driveIdValue")

read, err := client.GetDriveListPermissionsCount(ctx, id, drivelistpermission.DefaultGetDriveListPermissionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListPermissionClient.ListDriveListPermissionGrants`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveIdListPermissionID("userIdValue", "driveIdValue", "permissionIdValue")

payload := drivelistpermission.ListDriveListPermissionGrantsRequest{
	// ...
}


// alternatively `client.ListDriveListPermissionGrants(ctx, id, payload, drivelistpermission.DefaultListDriveListPermissionGrantsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListPermissionGrantsComplete(ctx, id, payload, drivelistpermission.DefaultListDriveListPermissionGrantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListPermissionClient.ListDriveListPermissions`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveID("userIdValue", "driveIdValue")

// alternatively `client.ListDriveListPermissions(ctx, id, drivelistpermission.DefaultListDriveListPermissionsOperationOptions())` can be used to do batched pagination
items, err := client.ListDriveListPermissionsComplete(ctx, id, drivelistpermission.DefaultListDriveListPermissionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DriveListPermissionClient.RevokeDriveListPermissionGrant`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveIdListPermissionID("userIdValue", "driveIdValue", "permissionIdValue")

payload := drivelistpermission.RevokeDriveListPermissionGrantRequest{
	// ...
}


read, err := client.RevokeDriveListPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveListPermissionClient.UpdateDriveListPermission`

```go
ctx := context.TODO()
id := drivelistpermission.NewUserIdDriveIdListPermissionID("userIdValue", "driveIdValue", "permissionIdValue")

payload := drivelistpermission.Permission{
	// ...
}


read, err := client.UpdateDriveListPermission(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
