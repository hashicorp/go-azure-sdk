
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdateprofile` Documentation

The `windowsqualityupdateprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdateprofile"
```


### Client Initialization

```go
client := windowsqualityupdateprofile.NewWindowsQualityUpdateProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsQualityUpdateProfileClient.AssignWindowsQualityUpdateProfile`

```go
ctx := context.TODO()
id := windowsqualityupdateprofile.NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue")

payload := windowsqualityupdateprofile.AssignWindowsQualityUpdateProfileRequest{
	// ...
}


read, err := client.AssignWindowsQualityUpdateProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.CreateWindowsQualityUpdateProfile`

```go
ctx := context.TODO()

payload := windowsqualityupdateprofile.WindowsQualityUpdateProfile{
	// ...
}


read, err := client.CreateWindowsQualityUpdateProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.DeleteWindowsQualityUpdateProfile`

```go
ctx := context.TODO()
id := windowsqualityupdateprofile.NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue")

read, err := client.DeleteWindowsQualityUpdateProfile(ctx, id, windowsqualityupdateprofile.DefaultDeleteWindowsQualityUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.GetWindowsQualityUpdateProfile`

```go
ctx := context.TODO()
id := windowsqualityupdateprofile.NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue")

read, err := client.GetWindowsQualityUpdateProfile(ctx, id, windowsqualityupdateprofile.DefaultGetWindowsQualityUpdateProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.GetWindowsQualityUpdateProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsQualityUpdateProfilesCount(ctx, windowsqualityupdateprofile.DefaultGetWindowsQualityUpdateProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.ListWindowsQualityUpdateProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsQualityUpdateProfiles(ctx, windowsqualityupdateprofile.DefaultListWindowsQualityUpdateProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsQualityUpdateProfilesComplete(ctx, windowsqualityupdateprofile.DefaultListWindowsQualityUpdateProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsQualityUpdateProfileClient.UpdateWindowsQualityUpdateProfile`

```go
ctx := context.TODO()
id := windowsqualityupdateprofile.NewDeviceManagementWindowsQualityUpdateProfileID("windowsQualityUpdateProfileIdValue")

payload := windowsqualityupdateprofile.WindowsQualityUpdateProfile{
	// ...
}


read, err := client.UpdateWindowsQualityUpdateProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
