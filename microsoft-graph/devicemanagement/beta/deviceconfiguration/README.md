
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfiguration` Documentation

The `deviceconfiguration` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceconfiguration"
```


### Client Initialization

```go
client := deviceconfiguration.NewDeviceConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceConfigurationClient.AssignDeviceManagementDeviceConfigurations`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

payload := deviceconfiguration.AssignDeviceManagementDeviceConfigurationsRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementDeviceConfigurations(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementDeviceConfigurationsComplete(ctx, id, payload)
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


read, err := client.CreateDeviceConfiguration(ctx, payload)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

payload := deviceconfiguration.CreateDeviceConfigurationAssignedAccessMultiModeProfileRequest{
	// ...
}


read, err := client.CreateDeviceConfigurationAssignedAccessMultiModeProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.CreateDeviceConfigurationHasPayloadLink`

```go
ctx := context.TODO()

payload := deviceconfiguration.CreateDeviceConfigurationHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateDeviceConfigurationHasPayloadLink(ctx, payload)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

payload := deviceconfiguration.CreateDeviceConfigurationWindowsPrivacyAccessControlRequest{
	// ...
}


read, err := client.CreateDeviceConfigurationWindowsPrivacyAccessControl(ctx, id, payload)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

read, err := client.DeleteDeviceConfiguration(ctx, id)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

read, err := client.GetDeviceConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.GetDeviceConfigurationCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceConfigurationCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.GetDeviceManagementDeviceConfigurationsTargetedUsersAndDevice`

```go
ctx := context.TODO()

payload := deviceconfiguration.GetDeviceManagementDeviceConfigurationsTargetedUsersAndDeviceRequest{
	// ...
}


read, err := client.GetDeviceManagementDeviceConfigurationsTargetedUsersAndDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceConfigurationClient.ListDeviceConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceConfigurations(ctx)` can be used to do batched pagination
items, err := client.ListDeviceConfigurationsComplete(ctx)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

payload := deviceconfiguration.DeviceConfiguration{
	// ...
}


read, err := client.UpdateDeviceConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
