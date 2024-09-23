
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevice` Documentation

The `windowsautopilotdeploymentprofileassigneddevice` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevice"
```


### Client Initialization

```go
client := windowsautopilotdeploymentprofileassigneddevice.NewWindowsAutopilotDeploymentProfileAssignedDeviceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice(ctx, id, payload, windowsautopilotdeploymentprofileassigneddevice.DefaultAssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice(ctx, id, payload, windowsautopilotdeploymentprofileassigneddevice.DefaultAssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.CreateWindowsAutopilotDeploymentProfileAssignedDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileId")

payload := windowsautopilotdeploymentprofileassigneddevice.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, payload, windowsautopilotdeploymentprofileassigneddevice.DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.CreateWindowsAutopilotDeploymentProfileAssignedDeviceAllowNextEnrollment`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceAllowNextEnrollment(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceAllowNextEnrollmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDevice(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignUserFromDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignUserFromDevice(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultCreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignUserFromDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.DeleteWindowsAutopilotDeploymentProfileAssignedDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

read, err := client.DeleteWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultDeleteWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.GetWindowsAutopilotDeploymentProfileAssignedDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

read, err := client.GetWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultGetWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.GetWindowsAutopilotDeploymentProfileAssignedDevicesCount`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileId")

read, err := client.GetWindowsAutopilotDeploymentProfileAssignedDevicesCount(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultGetWindowsAutopilotDeploymentProfileAssignedDevicesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.ListWindowsAutopilotDeploymentProfileAssignedDevices`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileId")

// alternatively `client.ListWindowsAutopilotDeploymentProfileAssignedDevices(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultListWindowsAutopilotDeploymentProfileAssignedDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeploymentProfileAssignedDevicesComplete(ctx, id, windowsautopilotdeploymentprofileassigneddevice.DefaultListWindowsAutopilotDeploymentProfileAssignedDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.UpdateWindowsAutopilotDeploymentProfileAssignedDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeploymentprofileassigneddevice.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, payload, windowsautopilotdeploymentprofileassigneddevice.DefaultUpdateWindowsAutopilotDeploymentProfileAssignedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperties`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileId", "windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeploymentprofileassigneddevice.UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesRequest{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperties(ctx, id, payload, windowsautopilotdeploymentprofileassigneddevice.DefaultUpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
