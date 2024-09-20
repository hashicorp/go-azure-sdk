
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/virtualendpointusersetting` Documentation

The `virtualendpointusersetting` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/virtualendpointusersetting"
```


### Client Initialization

```go
client := virtualendpointusersetting.NewVirtualEndpointUserSettingClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointUserSettingClient.AssignVirtualEndpointUserSetting`

```go
ctx := context.TODO()
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId")

payload := virtualendpointusersetting.AssignVirtualEndpointUserSettingRequest{
	// ...
}


read, err := client.AssignVirtualEndpointUserSetting(ctx, id, payload, virtualendpointusersetting.DefaultAssignVirtualEndpointUserSettingOperationOptions())
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


read, err := client.CreateVirtualEndpointUserSetting(ctx, payload, virtualendpointusersetting.DefaultCreateVirtualEndpointUserSettingOperationOptions())
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
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId")

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
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId")

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
id := virtualendpointusersetting.NewDeviceManagementVirtualEndpointUserSettingID("cloudPCUserSettingId")

payload := virtualendpointusersetting.CloudPCUserSetting{
	// ...
}


read, err := client.UpdateVirtualEndpointUserSetting(ctx, id, payload, virtualendpointusersetting.DefaultUpdateVirtualEndpointUserSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
