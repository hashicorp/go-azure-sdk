
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevice` Documentation

The `manageddevice` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/manageddevice"
```


### Client Initialization

```go
client := manageddevice.NewManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceClient.BypassManagedDeviceActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.BypassManagedDeviceActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CleanManagedDeviceWindowsDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CleanManagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanManagedDeviceWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDevice`

```go
ctx := context.TODO()

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateManagedDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWindowsDefenderScan`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateManagedDeviceWindowsDefenderScan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DeleteManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteManagedDevice(ctx, id, manageddevice.DefaultDeleteManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DeleteManagedDeviceUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.DeleteManagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteManagedDeviceUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DisableManagedDeviceLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.DisableManagedDeviceLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.GetManagedDevice(ctx, id, manageddevice.DefaultGetManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDevicesCount`

```go
ctx := context.TODO()


read, err := client.GetManagedDevicesCount(ctx, manageddevice.DefaultGetManagedDevicesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ListManagedDevices`

```go
ctx := context.TODO()


// alternatively `client.ListManagedDevices(ctx, manageddevice.DefaultListManagedDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListManagedDevicesComplete(ctx, manageddevice.DefaultListManagedDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceClient.LocateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.LocateManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.LogoutManagedDeviceSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.LogoutManagedDeviceSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RebootManagedDeviceNow`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RebootManagedDeviceNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RecoverManagedDevicePasscode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RecoverManagedDevicePasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RemoteLockManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RemoteLockManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RequestManagedDeviceRemoteAssistance`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RequestManagedDeviceRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ResetManagedDevicePasscode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.ResetManagedDevicePasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RetireManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RetireManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ShutDownManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.ShutDownManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SyncManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.SyncManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.UpdateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.UpdateManagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.UpdateManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateManagedDeviceWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.WipeManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.WipeManagedDeviceRequest{
	// ...
}


read, err := client.WipeManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
