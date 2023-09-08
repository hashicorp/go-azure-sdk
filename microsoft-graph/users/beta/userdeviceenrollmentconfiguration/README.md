
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userdeviceenrollmentconfiguration` Documentation

The `userdeviceenrollmentconfiguration` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userdeviceenrollmentconfiguration"
```


### Client Initialization

```go
client := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.AssignUserByIdDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

payload := userdeviceenrollmentconfiguration.AssignUserByIdDeviceEnrollmentConfigurationByIdRequest{
	// ...
}


read, err := client.AssignUserByIdDeviceEnrollmentConfigurationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.CreateUserByIdDeviceEnrollmentConfiguration`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserID("userIdValue")

payload := userdeviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.CreateUserByIdDeviceEnrollmentConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.CreateUserByIdDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserID("userIdValue")

payload := userdeviceenrollmentconfiguration.CreateUserByIdDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfigurationRequest{
	// ...
}


read, err := client.CreateUserByIdDeviceEnrollmentConfigurationCreateEnrollmentNotificationConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.CreateUserByIdDeviceEnrollmentConfigurationHasPayloadLink`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserID("userIdValue")

payload := userdeviceenrollmentconfiguration.CreateUserByIdDeviceEnrollmentConfigurationHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateUserByIdDeviceEnrollmentConfigurationHasPayloadLink(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.DeleteUserByIdDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

read, err := client.DeleteUserByIdDeviceEnrollmentConfigurationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.GetUserByIdDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

read, err := client.GetUserByIdDeviceEnrollmentConfigurationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.GetUserByIdDeviceEnrollmentConfigurationCount`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserID("userIdValue")

read, err := client.GetUserByIdDeviceEnrollmentConfigurationCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.ListUserByIdDeviceEnrollmentConfigurations`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserID("userIdValue")

// alternatively `client.ListUserByIdDeviceEnrollmentConfigurations(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdDeviceEnrollmentConfigurationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.SetUserByIdDeviceEnrollmentConfigurationByIdPriority`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

payload := userdeviceenrollmentconfiguration.SetUserByIdDeviceEnrollmentConfigurationByIdPriorityRequest{
	// ...
}


read, err := client.SetUserByIdDeviceEnrollmentConfigurationByIdPriority(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceEnrollmentConfigurationClient.UpdateUserByIdDeviceEnrollmentConfigurationById`

```go
ctx := context.TODO()
id := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationID("userIdValue", "deviceEnrollmentConfigurationIdValue")

payload := userdeviceenrollmentconfiguration.DeviceEnrollmentConfiguration{
	// ...
}


read, err := client.UpdateUserByIdDeviceEnrollmentConfigurationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
