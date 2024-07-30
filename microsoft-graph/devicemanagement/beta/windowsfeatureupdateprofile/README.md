
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofile` Documentation

The `windowsfeatureupdateprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsfeatureupdateprofile"
```


### Client Initialization

```go
client := windowsfeatureupdateprofile.NewWindowsFeatureUpdateProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsFeatureUpdateProfileClient.AssignDeviceManagementWindowsFeatureUpdateProfile`

```go
ctx := context.TODO()
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileIdValue")

payload := windowsfeatureupdateprofile.AssignDeviceManagementWindowsFeatureUpdateProfileRequest{
	// ...
}


read, err := client.AssignDeviceManagementWindowsFeatureUpdateProfile(ctx, id, payload)
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


read, err := client.CreateWindowsFeatureUpdateProfile(ctx, payload)
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
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileIdValue")

read, err := client.DeleteWindowsFeatureUpdateProfile(ctx, id)
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
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileIdValue")

read, err := client.GetWindowsFeatureUpdateProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsFeatureUpdateProfileClient.GetWindowsFeatureUpdateProfileCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsFeatureUpdateProfileCount(ctx)
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


// alternatively `client.ListWindowsFeatureUpdateProfiles(ctx)` can be used to do batched pagination
items, err := client.ListWindowsFeatureUpdateProfilesComplete(ctx)
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
id := windowsfeatureupdateprofile.NewDeviceManagementWindowsFeatureUpdateProfileID("windowsFeatureUpdateProfileIdValue")

payload := windowsfeatureupdateprofile.WindowsFeatureUpdateProfile{
	// ...
}


read, err := client.UpdateWindowsFeatureUpdateProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
