
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfiguration` Documentation

The `deviceconfiguration` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceconfiguration"
```


### Client Initialization

```go
client := deviceconfiguration.NewDeviceConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceConfigurationClient.AssignDeviceConfigurations`

```go
ctx := context.TODO()
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

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


read, err := client.CreateDeviceConfiguration(ctx, payload)
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
id := deviceconfiguration.NewDeviceManagementDeviceConfigurationID("deviceConfigurationIdValue")

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
