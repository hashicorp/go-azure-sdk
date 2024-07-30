
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


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDeviceRequest{
	// ...
}


read, err := client.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceResourceAccountToDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDeviceRequest{
	// ...
}


read, err := client.AssignDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceUserToDevice(ctx, id, payload)
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

read, err := client.DeleteWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id)
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

read, err := client.GetWindowsAutopilotDeploymentProfileAssignedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.GetWindowsAutopilotDeploymentProfileAssignedDeviceCount`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileID("windowsAutopilotDeploymentProfileIdValue")

read, err := client.GetWindowsAutopilotDeploymentProfileAssignedDeviceCount(ctx, id)
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

// alternatively `client.ListWindowsAutopilotDeploymentProfileAssignedDevices(ctx, id)` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeploymentProfileAssignedDevicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeploymentProfileAssignedDeviceClient.UpdateDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceDeviceProperty`

```go
ctx := context.TODO()
id := windowsautopilotdeploymentprofileassigneddevice.NewDeviceManagementWindowsAutopilotDeploymentProfileIdAssignedDeviceID("windowsAutopilotDeploymentProfileIdValue", "windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeploymentprofileassigneddevice.UpdateDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceDevicePropertyRequest{
	// ...
}


read, err := client.UpdateDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceDeviceProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
