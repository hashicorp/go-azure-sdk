
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeviceidentity` Documentation

The `windowsautopilotdeviceidentity` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsautopilotdeviceidentity"
```


### Client Initialization

```go
client := windowsautopilotdeviceidentity.NewWindowsAutopilotDeviceIdentityClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.AssignWindowsAutopilotDeviceIdentityResourceAccountToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeviceidentity.AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeviceIdentityResourceAccountToDevice(ctx, id, payload, windowsautopilotdeviceidentity.DefaultAssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.AssignWindowsAutopilotDeviceIdentityUserToDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeviceidentity.AssignWindowsAutopilotDeviceIdentityUserToDeviceRequest{
	// ...
}


read, err := client.AssignWindowsAutopilotDeviceIdentityUserToDevice(ctx, id, payload, windowsautopilotdeviceidentity.DefaultAssignWindowsAutopilotDeviceIdentityUserToDeviceOperationOptions())
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


read, err := client.CreateWindowsAutopilotDeviceIdentity(ctx, payload, windowsautopilotdeviceidentity.DefaultCreateWindowsAutopilotDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.CreateWindowsAutopilotDeviceIdentityAllowNextEnrollment`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeviceIdentityAllowNextEnrollment(ctx, id, windowsautopilotdeviceidentity.DefaultCreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDevice`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDevice(ctx, id, windowsautopilotdeviceidentity.DefaultCreateWindowsAutopilotDeviceIdentityUnassignResourceAccountFromDeviceOperationOptions())
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
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

read, err := client.CreateWindowsAutopilotDeviceIdentityUnassignUserFromDevice(ctx, id, windowsautopilotdeviceidentity.DefaultCreateWindowsAutopilotDeviceIdentityUnassignUserFromDeviceOperationOptions())
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
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

read, err := client.DeleteWindowsAutopilotDeviceIdentity(ctx, id, windowsautopilotdeviceidentity.DefaultDeleteWindowsAutopilotDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.GetWindowsAutopilotDeviceIdentitiesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsAutopilotDeviceIdentitiesCount(ctx, windowsautopilotdeviceidentity.DefaultGetWindowsAutopilotDeviceIdentitiesCountOperationOptions())
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
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

read, err := client.GetWindowsAutopilotDeviceIdentity(ctx, id, windowsautopilotdeviceidentity.DefaultGetWindowsAutopilotDeviceIdentityOperationOptions())
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


// alternatively `client.ListWindowsAutopilotDeviceIdentities(ctx, windowsautopilotdeviceidentity.DefaultListWindowsAutopilotDeviceIdentitiesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsAutopilotDeviceIdentitiesComplete(ctx, windowsautopilotdeviceidentity.DefaultListWindowsAutopilotDeviceIdentitiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.UpdateWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeviceidentity.WindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeviceIdentity(ctx, id, payload, windowsautopilotdeviceidentity.DefaultUpdateWindowsAutopilotDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsAutopilotDeviceIdentityClient.UpdateWindowsAutopilotDeviceIdentityDeviceProperties`

```go
ctx := context.TODO()
id := windowsautopilotdeviceidentity.NewDeviceManagementWindowsAutopilotDeviceIdentityID("windowsAutopilotDeviceIdentityId")

payload := windowsautopilotdeviceidentity.UpdateWindowsAutopilotDeviceIdentityDevicePropertiesRequest{
	// ...
}


read, err := client.UpdateWindowsAutopilotDeviceIdentityDeviceProperties(ctx, id, payload, windowsautopilotdeviceidentity.DefaultUpdateWindowsAutopilotDeviceIdentityDevicePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
