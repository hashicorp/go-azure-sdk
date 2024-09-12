
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceenrollmentconfiguration` Documentation

The `deviceenrollmentconfiguration` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/deviceenrollmentconfiguration"
```


### Client Initialization

```go
client := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceEnrollmentConfigurationClient.AssignDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewDeviceManagementDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.AssignDeviceEnrollmentConfigurationRequest{
	// ...
}


read, err := client.AssignDeviceEnrollmentConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.CreateDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()

payload := deviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfiguration(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.DeleteDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewDeviceManagementDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

read, err := client.DeleteDeviceEnrollmentConfiguration(ctx, id, deviceenrollmentconfiguration.DefaultDeleteDeviceEnrollmentConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.GetDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewDeviceManagementDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

read, err := client.GetDeviceEnrollmentConfiguration(ctx, id, deviceenrollmentconfiguration.DefaultGetDeviceEnrollmentConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.GetDeviceEnrollmentConfigurationsCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceEnrollmentConfigurationsCount(ctx, deviceenrollmentconfiguration.DefaultGetDeviceEnrollmentConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.ListDeviceEnrollmentConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceEnrollmentConfigurations(ctx, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceEnrollmentConfigurationsComplete(ctx, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.SetDeviceEnrollmentConfigurationPriority`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewDeviceManagementDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.SetDeviceEnrollmentConfigurationPriorityRequest{
	// ...
}


read, err := client.SetDeviceEnrollmentConfigurationPriority(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.UpdateDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewDeviceManagementDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.UpdateDeviceEnrollmentConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
