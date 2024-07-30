
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsautopilotdeviceidentity` Documentation

The `windowsautopilotdeviceidentity` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/windowsautopilotdeviceidentity"
```


### Client Initialization

```go
client := windowsautopilotdeviceidentity.NewWindowsAutopilotDeviceIdentityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.AssignDeviceManagementWindowsAutopilotDeviceIdentityUserToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeviceidentity.AssignDeviceManagementWindowsAutopilotDeviceIdentityUserToDeviceRequest{
	// ...
}


read, err := client.AssignDeviceManagementWindowsAutopilotDeviceIdentityUserToDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.CreateWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()

payload := windowsautopilotdeviceidentity.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.CreateWindowsAutopilotDeviceIdentity(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.CreateWindowsAutopilotDeviceIdentityUnassignUserFromDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

read, err := client.CreateWindowsAutopilotDeviceIdentityUnassignUserFromDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.DeleteWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

read, err := client.DeleteWindowsAutopilotDeviceIdentity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.GetWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

read, err := client.GetWindowsAutopilotDeviceIdentity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.GetWindowsAutopilotDeviceIdentityCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsAutopilotDeviceIdentityCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.ListWindowsAutopilotDeviceIdentities`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsAutopilotDeviceIdentities(ctx)` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeviceIdentitiesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.UpdateDeviceManagementWindowsAutopilotDeviceIdentityDeviceProperty`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeviceidentity.UpdateDeviceManagementWindowsAutopilotDeviceIdentityDevicePropertyRequest{
	// ...
}


read, err := client.UpdateDeviceManagementWindowsAutopilotDeviceIdentityDeviceProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.UpdateWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityIdValue")

payload := windowsautopilotdeviceidentity.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeviceIdentity(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
