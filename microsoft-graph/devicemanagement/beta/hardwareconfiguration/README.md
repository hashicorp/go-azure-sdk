
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfiguration` Documentation

The `hardwareconfiguration` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/hardwareconfiguration"
```


### Client Initialization

```go
client := hardwareconfiguration.NewHardwareConfigurationClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HardwareConfigurationClient.AssignHardwareConfigurations`

```go
ctx := context.TODO()
id := hardwareconfiguration.NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId")

payload := hardwareconfiguration.AssignHardwareConfigurationsRequest{
	// ...
}


// alternatively `client.AssignHardwareConfigurations(ctx, id, payload, hardwareconfiguration.DefaultAssignHardwareConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.AssignHardwareConfigurationsComplete(ctx, id, payload, hardwareconfiguration.DefaultAssignHardwareConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HardwareConfigurationClient.CreateHardwareConfiguration`

```go
ctx := context.TODO()

payload := hardwareconfiguration.HardwareConfiguration{
	// ...
}


read, err := client.CreateHardwareConfiguration(ctx, payload, hardwareconfiguration.DefaultCreateHardwareConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HardwareConfigurationClient.DeleteHardwareConfiguration`

```go
ctx := context.TODO()
id := hardwareconfiguration.NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId")

read, err := client.DeleteHardwareConfiguration(ctx, id, hardwareconfiguration.DefaultDeleteHardwareConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HardwareConfigurationClient.GetHardwareConfiguration`

```go
ctx := context.TODO()
id := hardwareconfiguration.NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId")

read, err := client.GetHardwareConfiguration(ctx, id, hardwareconfiguration.DefaultGetHardwareConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HardwareConfigurationClient.GetHardwareConfigurationsCount`

```go
ctx := context.TODO()


read, err := client.GetHardwareConfigurationsCount(ctx, hardwareconfiguration.DefaultGetHardwareConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HardwareConfigurationClient.ListHardwareConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListHardwareConfigurations(ctx, hardwareconfiguration.DefaultListHardwareConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListHardwareConfigurationsComplete(ctx, hardwareconfiguration.DefaultListHardwareConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HardwareConfigurationClient.UpdateHardwareConfiguration`

```go
ctx := context.TODO()
id := hardwareconfiguration.NewDeviceManagementHardwareConfigurationID("hardwareConfigurationId")

payload := hardwareconfiguration.HardwareConfiguration{
	// ...
}


read, err := client.UpdateHardwareConfiguration(ctx, id, payload, hardwareconfiguration.DefaultUpdateHardwareConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
