
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfiguration` Documentation

The `deviceconfiguration` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfiguration"
```


### Client Initialization

```go
client := deviceconfiguration.NewDeviceConfigurationClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceConfigurationClient.AssignDeviceConfigurations`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

payload := deviceconfiguration.AssignDeviceConfigurationsRequest{
	// ...
}


// alternatively `client.AssignDeviceConfigurations(ctx, id, payload, deviceconfiguration.DefaultAssignDeviceConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.AssignDeviceConfigurationsComplete(ctx, id, payload, deviceconfiguration.DefaultAssignDeviceConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceConfigurationClient.CreateDeviceConfiguration`

```go
ctx := context.TODO()

payload := deviceconfiguration.DeviceConfiguration{
	// ...
}


read, err := client.CreateDeviceConfiguration(ctx, payload, deviceconfiguration.DefaultCreateDeviceConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.CreateDeviceConfigurationAssignedAccessMultiModeProfile`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

payload := deviceconfiguration.CreateDeviceConfigurationAssignedAccessMultiModeProfileRequest{
	// ...
}


read, err := client.CreateDeviceConfigurationAssignedAccessMultiModeProfile(ctx, id, payload, deviceconfiguration.DefaultCreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.CreateDeviceConfigurationWindowsPrivacyAccessControl`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

payload := deviceconfiguration.CreateDeviceConfigurationWindowsPrivacyAccessControlRequest{
	// ...
}


read, err := client.CreateDeviceConfigurationWindowsPrivacyAccessControl(ctx, id, payload, deviceconfiguration.DefaultCreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.DeleteDeviceConfiguration`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

read, err := client.DeleteDeviceConfiguration(ctx, id, deviceconfiguration.DefaultDeleteDeviceConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.GetDeviceConfiguration`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

read, err := client.GetDeviceConfiguration(ctx, id, deviceconfiguration.DefaultGetDeviceConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.GetDeviceConfigurationsCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceConfigurationsCount(ctx, deviceconfiguration.DefaultGetDeviceConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.GetDeviceConfigurationsTargetedUsersAndDevices`

```go
ctx := context.TODO()

payload := deviceconfiguration.GetDeviceConfigurationsTargetedUsersAndDevicesRequest{
	// ...
}


// alternatively `client.GetDeviceConfigurationsTargetedUsersAndDevices(ctx, payload, deviceconfiguration.DefaultGetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions())` can be used to do batched pagination
items, err := client.GetDeviceConfigurationsTargetedUsersAndDevicesComplete(ctx, payload, deviceconfiguration.DefaultGetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceConfigurationClient.ListDeviceConfigurationHasPayloadLinks`

```go
ctx := context.TODO()

payload := deviceconfiguration.ListDeviceConfigurationHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListDeviceConfigurationHasPayloadLinks(ctx, payload, deviceconfiguration.DefaultListDeviceConfigurationHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceConfigurationHasPayloadLinksComplete(ctx, payload, deviceconfiguration.DefaultListDeviceConfigurationHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceConfigurationClient.ListDeviceConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceConfigurations(ctx, deviceconfiguration.DefaultListDeviceConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceConfigurationsComplete(ctx, deviceconfiguration.DefaultListDeviceConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceConfigurationClient.UpdateDeviceConfiguration`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationId")

payload := deviceconfiguration.DeviceConfiguration{
	// ...
}


read, err := client.UpdateDeviceConfiguration(ctx, id, payload, deviceconfiguration.DefaultUpdateDeviceConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
