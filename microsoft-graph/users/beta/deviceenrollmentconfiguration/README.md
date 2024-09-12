
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceenrollmentconfiguration` Documentation

The `deviceenrollmentconfiguration` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceenrollmentconfiguration"
```


### Client Initialization

```go
client := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceEnrollmentConfigurationClient.AssignDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

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
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

payload := deviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.CreateDeviceEnrollmentConfigurationsEnrollmentNotificationConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

payload := deviceenrollmentconfiguration.CreateDeviceEnrollmentConfigurationsEnrollmentNotificationConfigurationRequest{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfigurationsEnrollmentNotificationConfiguration(ctx, id, payload)
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
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

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
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

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
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

read, err := client.GetDeviceEnrollmentConfigurationsCount(ctx, id, deviceenrollmentconfiguration.DefaultGetDeviceEnrollmentConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.ListDeviceEnrollmentConfigurationHasPayloadLinks`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

payload := deviceenrollmentconfiguration.ListDeviceEnrollmentConfigurationHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListDeviceEnrollmentConfigurationHasPayloadLinks(ctx, id, payload, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceEnrollmentConfigurationHasPayloadLinksComplete(ctx, id, payload, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.ListDeviceEnrollmentConfigurations`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

// alternatively `client.ListDeviceEnrollmentConfigurations(ctx, id, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceEnrollmentConfigurationsComplete(ctx, id, deviceenrollmentconfiguration.DefaultListDeviceEnrollmentConfigurationsOperationOptions())
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
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

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
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

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
