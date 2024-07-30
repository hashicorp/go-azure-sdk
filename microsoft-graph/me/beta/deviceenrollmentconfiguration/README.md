
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/deviceenrollmentconfiguration` Documentation

The `deviceenrollmentconfiguration` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/deviceenrollmentconfiguration"
```


### Client Initialization

```go
client := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceEnrollmentConfigurationClient.AssignMeDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.AssignMeDeviceEnrollmentConfigurationRequest{
	// ...
}


read, err := client.AssignMeDeviceEnrollmentConfiguration(ctx, id, payload)
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


### Example Usage: `DeviceEnrollmentConfigurationClient.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration`

```go
ctx := context.TODO()

payload := deviceenrollmentconfiguration.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfigurationRequest{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration(ctx, payload)
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

payload := deviceenrollmentconfiguration.CreateDeviceEnrollmentConfigurationHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateDeviceEnrollmentConfigurationHasPayloadLink(ctx, payload)
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
id := deviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

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
id := deviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

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


read, err := client.GetDeviceEnrollmentConfigurationCount(ctx)
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


// alternatively `client.ListDeviceEnrollmentConfigurations(ctx)` can be used to do batched pagination
items, err := client.ListDeviceEnrollmentConfigurationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceEnrollmentConfigurationClient.SetMeDeviceEnrollmentConfigurationPriority`

```go
ctx := context.TODO()
id := deviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := deviceenrollmentconfiguration.SetMeDeviceEnrollmentConfigurationPriorityRequest{
	// ...
}


read, err := client.SetMeDeviceEnrollmentConfigurationPriority(ctx, id, payload)
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
id := deviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

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
