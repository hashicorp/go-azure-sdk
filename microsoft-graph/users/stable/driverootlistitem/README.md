
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitem` Documentation

The `driverootlistitem` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitem"
```


### Client Initialization

```go
client := driverootlistitem.NewDriveRootListItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DriveRootListItemClient.CreateDriveRootListItemLink`

```go
ctx := context.TODO()
id := driverootlistitem.NewUserIdDriveID("userId", "driveId")

payload := driverootlistitem.CreateDriveRootListItemLinkRequest{
	// ...
}


read, err := client.CreateDriveRootListItemLink(ctx, id, payload, driverootlistitem.DefaultCreateDriveRootListItemLinkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemClient.DeleteDriveRootListItem`

```go
ctx := context.TODO()
id := driverootlistitem.NewUserIdDriveID("userId", "driveId")

read, err := client.DeleteDriveRootListItem(ctx, id, driverootlistitem.DefaultDeleteDriveRootListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemClient.GetDriveRootListItem`

```go
ctx := context.TODO()
id := driverootlistitem.NewUserIdDriveID("userId", "driveId")

read, err := client.GetDriveRootListItem(ctx, id, driverootlistitem.DefaultGetDriveRootListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DriveRootListItemClient.UpdateDriveRootListItem`

```go
ctx := context.TODO()
id := driverootlistitem.NewUserIdDriveID("userId", "driveId")

payload := driverootlistitem.ListItem{
	// ...
}


read, err := client.UpdateDriveRootListItem(ctx, id, payload, driverootlistitem.DefaultUpdateDriveRootListItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
