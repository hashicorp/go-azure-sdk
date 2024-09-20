
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofile` Documentation

The `windowsdriverupdateprofile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofile"
```


### Client Initialization

```go
client := windowsdriverupdateprofile.NewWindowsDriverUpdateProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsDriverUpdateProfileClient.AssignWindowsDriverUpdateProfile`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

payload := windowsdriverupdateprofile.AssignWindowsDriverUpdateProfileRequest{
	// ...
}


read, err := client.AssignWindowsDriverUpdateProfile(ctx, id, payload, windowsdriverupdateprofile.DefaultAssignWindowsDriverUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.CreateWindowsDriverUpdateProfile`

```go
ctx := context.TODO()

payload := windowsdriverupdateprofile.WindowsDriverUpdateProfile{
	// ...
}


read, err := client.CreateWindowsDriverUpdateProfile(ctx, payload, windowsdriverupdateprofile.DefaultCreateWindowsDriverUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.CreateWindowsDriverUpdateProfileExecuteAction`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

payload := windowsdriverupdateprofile.CreateWindowsDriverUpdateProfileExecuteActionRequest{
	// ...
}


read, err := client.CreateWindowsDriverUpdateProfileExecuteAction(ctx, id, payload, windowsdriverupdateprofile.DefaultCreateWindowsDriverUpdateProfileExecuteActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.DeleteWindowsDriverUpdateProfile`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

read, err := client.DeleteWindowsDriverUpdateProfile(ctx, id, windowsdriverupdateprofile.DefaultDeleteWindowsDriverUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.GetWindowsDriverUpdateProfile`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

read, err := client.GetWindowsDriverUpdateProfile(ctx, id, windowsdriverupdateprofile.DefaultGetWindowsDriverUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.GetWindowsDriverUpdateProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsDriverUpdateProfilesCount(ctx, windowsdriverupdateprofile.DefaultGetWindowsDriverUpdateProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.ListWindowsDriverUpdateProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsDriverUpdateProfiles(ctx, windowsdriverupdateprofile.DefaultListWindowsDriverUpdateProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsDriverUpdateProfilesComplete(ctx, windowsdriverupdateprofile.DefaultListWindowsDriverUpdateProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.SyncWindowsDriverUpdateProfileInventory`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

read, err := client.SyncWindowsDriverUpdateProfileInventory(ctx, id, windowsdriverupdateprofile.DefaultSyncWindowsDriverUpdateProfileInventoryOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsDriverUpdateProfileClient.UpdateWindowsDriverUpdateProfile`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileId")

payload := windowsdriverupdateprofile.WindowsDriverUpdateProfile{
	// ...
}


read, err := client.UpdateWindowsDriverUpdateProfile(ctx, id, payload, windowsdriverupdateprofile.DefaultUpdateWindowsDriverUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
