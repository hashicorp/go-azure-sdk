
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/medeviceenrollmentconfiguration` Documentation

The `medeviceenrollmentconfiguration` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/medeviceenrollmentconfiguration"
```


### Client Initialization

```go
client := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.AssignMeDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := medeviceenrollmentconfiguration.AssignMeDeviceEnrollmentConfigurationByIdRequest{
	// ...
}


read, err := client.AssignMeDeviceEnrollmentConfigurationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.CreateMeDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()

payload := medeviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.CreateMeDeviceEnrollmentConfiguration(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.CreateMeDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration`

```go
ctx := context.TODO()

payload := medeviceenrollmentconfiguration.CreateMeDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfigurationRequest{
	// ...
}


read, err := client.CreateMeDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.CreateMeDeviceEnrollmentConfigurationHasPayloadLink`

```go
ctx := context.TODO()

payload := medeviceenrollmentconfiguration.CreateMeDeviceEnrollmentConfigurationHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateMeDeviceEnrollmentConfigurationHasPayloadLink(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.DeleteMeDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

read, err := client.DeleteMeDeviceEnrollmentConfigurationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.GetMeDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

read, err := client.GetMeDeviceEnrollmentConfigurationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.GetMeDeviceEnrollmentConfigurationCount`

```go
ctx := context.TODO()


read, err := client.GetMeDeviceEnrollmentConfigurationCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.ListMeDeviceEnrollmentConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListMeDeviceEnrollmentConfigurations(ctx)` can be used to do batched pagination
items, err := client.ListMeDeviceEnrollmentConfigurationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.SetMeDeviceEnrollmentConfigurationByIdPriority`

```go
ctx := context.TODO()
id := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := medeviceenrollmentconfiguration.SetMeDeviceEnrollmentConfigurationByIdPriorityRequest{
	// ...
}


read, err := client.SetMeDeviceEnrollmentConfigurationByIdPriority(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceEnrollmentConfigurationClient.UpdateMeDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationID("deviceEnrollmentConfigurationIdValue")

payload := medeviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.UpdateMeDeviceEnrollmentConfigurationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
