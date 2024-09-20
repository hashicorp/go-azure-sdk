
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofile` Documentation

The `windowsfeatureupdateprofile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofile"
```


### Client Initialization

```go
client := windowsfeatureupdateprofile.NewWindowsFeatureUpdateProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsFeatureUpdateProfileClient.AssignWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId")

payload := windowsfeatureupdateprofile.AssignWindowsFeatureUpdateProfileRequest{
	// ...
}


read, err := client.AssignWindowsFeatureUpdateProfile(ctx, id, payload, windowsfeatureupdateprofile.DefaultAssignWindowsFeatureUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.CreateWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()

payload := windowsfeatureupdateprofile.WindowsFeatureUpdateProfile{
	// ...
}


read, err := client.CreateWindowsFeatureUpdateProfile(ctx, payload, windowsfeatureupdateprofile.DefaultCreateWindowsFeatureUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.DeleteWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId")

read, err := client.DeleteWindowsFeatureUpdateProfile(ctx, id, windowsfeatureupdateprofile.DefaultDeleteWindowsFeatureUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.GetWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId")

read, err := client.GetWindowsFeatureUpdateProfile(ctx, id, windowsfeatureupdateprofile.DefaultGetWindowsFeatureUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.GetWindowsFeatureUpdateProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsFeatureUpdateProfilesCount(ctx, windowsfeatureupdateprofile.DefaultGetWindowsFeatureUpdateProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.ListWindowsFeatureUpdateProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsFeatureUpdateProfiles(ctx, windowsfeatureupdateprofile.DefaultListWindowsFeatureUpdateProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsFeatureUpdateProfilesComplete(ctx, windowsfeatureupdateprofile.DefaultListWindowsFeatureUpdateProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.UpdateWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileId")

payload := windowsfeatureupdateprofile.WindowsFeatureUpdateProfile{
	// ...
}


read, err := client.UpdateWindowsFeatureUpdateProfile(ctx, id, payload, windowsfeatureupdateprofile.DefaultUpdateWindowsFeatureUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
