
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


### Example Usage: `DeviceEnrollmentConfigurationClient.AssignUserDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.AssignUserDeviceEnrollmentConfigurationRequest{
	// ...
}


read, err := client.AssignUserDeviceEnrollmentConfiguration(ctx, id, payload)
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


### Example Usage: `DeviceEnrollmentConfigurationClient.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

payload := deviceenrollmentconfiguration.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfigurationRequest{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.CreateDeviceEnrollmentConfigurationHasPayloadLink`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

payload := deviceenrollmentconfiguration.CreateDeviceEnrollmentConfigurationHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfigurationHasPayloadLink(ctx, id, payload)
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

read, err := client.DeleteDeviceEnrollmentConfiguration(ctx, id)
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

read, err := client.GetDeviceEnrollmentConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.GetDeviceEnrollmentConfigurationCount`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

read, err := client.GetDeviceEnrollmentConfigurationCount(ctx, id)
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
id := deviceenrollmentconfiguration.NewUserID("userIdValue")

// alternatively `client.ListDeviceEnrollmentConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.ListDeviceEnrollmentConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.SetUserDeviceEnrollmentConfigurationPriority`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewUserIdDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.SetUserDeviceEnrollmentConfigurationPriorityRequest{
	// ...
}


read, err := client.SetUserDeviceEnrollmentConfigurationPriority(ctx, id, payload)
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
