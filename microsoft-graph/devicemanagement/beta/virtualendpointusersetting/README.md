
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointusersetting` Documentation

The `virtualendpointusersetting` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointusersetting"
```


### Client Initialization

```go
client := virtualendpointusersetting.NewVirtualEndpointUserSettingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointUserSettingClient.AssignVirtualEndpointUserSetting`

```go
ctx := context.TODO()
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingIdValue")

payload := virtualendpointusersetting.AssignVirtualEndpointUserSettingRequest{
	// ...
}


read, err := client.AssignVirtualEndpointUserSetting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointUserSettingClient.CreateVirtualEndpointUserSetting`

```go
ctx := context.TODO()

payload := virtualendpointusersetting.CloudPCUserSetting{
	// ...
}


read, err := client.CreateVirtualEndpointUserSetting(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointUserSettingClient.DeleteVirtualEndpointUserSetting`

```go
ctx := context.TODO()
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingIdValue")

read, err := client.DeleteVirtualEndpointUserSetting(ctx, id, virtualendpointusersetting.DefaultDeleteVirtualEndpointUserSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointUserSettingClient.GetVirtualEndpointUserSetting`

```go
ctx := context.TODO()
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingIdValue")

read, err := client.GetVirtualEndpointUserSetting(ctx, id, virtualendpointusersetting.DefaultGetVirtualEndpointUserSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointUserSettingClient.GetVirtualEndpointUserSettingsCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointUserSettingsCount(ctx, virtualendpointusersetting.DefaultGetVirtualEndpointUserSettingsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointUserSettingClient.ListVirtualEndpointUserSettings`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointUserSettings(ctx, virtualendpointusersetting.DefaultListVirtualEndpointUserSettingsOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointUserSettingsComplete(ctx, virtualendpointusersetting.DefaultListVirtualEndpointUserSettingsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointUserSettingClient.UpdateVirtualEndpointUserSetting`

```go
ctx := context.TODO()
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingIdValue")

payload := virtualendpointusersetting.CloudPCUserSetting{
	// ...
}


read, err := client.UpdateVirtualEndpointUserSetting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
