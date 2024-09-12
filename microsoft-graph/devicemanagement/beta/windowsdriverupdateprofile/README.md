
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofile` Documentation

The `windowsdriverupdateprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsdriverupdateprofile"
```


### Client Initialization

```go
client := windowsdriverupdateprofile.NewWindowsDriverUpdateProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsDriverUpdateProfileClient.AssignWindowsDriverUpdateProfile`

```go
ctx := context.TODO()
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

payload := windowsdriverupdateprofile.AssignWindowsDriverUpdateProfileRequest{
	// ...
}


read, err := client.AssignWindowsDriverUpdateProfile(ctx, id, payload)
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


read, err := client.CreateWindowsDriverUpdateProfile(ctx, payload)
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
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

payload := windowsdriverupdateprofile.CreateWindowsDriverUpdateProfileExecuteActionRequest{
	// ...
}


read, err := client.CreateWindowsDriverUpdateProfileExecuteAction(ctx, id, payload)
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
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

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
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

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
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

read, err := client.SyncWindowsDriverUpdateProfileInventory(ctx, id)
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
id := windowsdriverupdateprofile.NewDeviceManagementWindowsDriverUpdateProfileID("windowsDriverUpdateProfileIdValue")

payload := windowsdriverupdateprofile.WindowsDriverUpdateProfile{
	// ...
}


read, err := client.UpdateWindowsDriverUpdateProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
