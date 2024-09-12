
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevice` Documentation

The `windowsautopilotdeploymentprofileassigneddevice` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeploymentprofileassigneddevice"
```


### Client Initialization

```go
client := windowsautopilotdeploymentprofileassigneddevice.NewWindowsAutopilotDeploymentProfileAssignedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice(ctx, id, payload)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice(ctx, id, payload)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, payload)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceAllowNextEnrollment(ctx, id)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignResourceAccountFromDevice(ctx, id)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

read, err := client.CreateWindowsAutopilotDeploymentProfileAssignedDeviceUnassignUserFromDevice(ctx, id)
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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

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
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperty`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.UpdateWindowsAutopilotDeploymentProfileAssignedDevicePropertyRequest{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeploymentProfileAssignedDeviceProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
