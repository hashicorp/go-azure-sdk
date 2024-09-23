
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/manageddevice` Documentation

The `manageddevice` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/manageddevice"
```


### Client Initialization

```go
client := manageddevice.NewManagedDeviceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceClient.BypassManagedDeviceActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.BypassManagedDeviceActivationLock(ctx, id, manageddevice.DefaultBypassManagedDeviceActivationLockOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.CleanManagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanManagedDeviceWindowsDevice(ctx, id, payload, manageddevice.DefaultCleanManagedDeviceWindowsDeviceOperationOptions())
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


read, err := client.CreateManagedDevice(ctx, payload, manageddevice.DefaultCreateManagedDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.CreateManagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateManagedDeviceWindowsDefenderScan(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceWindowsDefenderScanOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.CreateManagedDeviceWindowsDefenderUpdateSignature(ctx, id, manageddevice.DefaultCreateManagedDeviceWindowsDefenderUpdateSignatureOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.DeleteManagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteManagedDeviceUserFromSharedAppleDevice(ctx, id, payload, manageddevice.DefaultDeleteManagedDeviceUserFromSharedAppleDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.DisableManagedDeviceLostMode(ctx, id, manageddevice.DefaultDisableManagedDeviceLostModeOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

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


### Example Usage: `ManagedDeviceClient.LocateManagedDeviceDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.LocateManagedDeviceDevice(ctx, id, manageddevice.DefaultLocateManagedDeviceDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.LogoutManagedDeviceSharedAppleDeviceActiveUser(ctx, id, manageddevice.DefaultLogoutManagedDeviceSharedAppleDeviceActiveUserOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.RebootManagedDeviceNow(ctx, id, manageddevice.DefaultRebootManagedDeviceNowOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.RecoverManagedDevicePasscode(ctx, id, manageddevice.DefaultRecoverManagedDevicePasscodeOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.RemoteLockManagedDevice(ctx, id, manageddevice.DefaultRemoteLockManagedDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.RequestManagedDeviceRemoteAssistance(ctx, id, manageddevice.DefaultRequestManagedDeviceRemoteAssistanceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.ResetManagedDevicePasscode(ctx, id, manageddevice.DefaultResetManagedDevicePasscodeOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.RetireManagedDevice(ctx, id, manageddevice.DefaultRetireManagedDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.ShutDownManagedDevice(ctx, id, manageddevice.DefaultShutDownManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SyncManagedDeviceDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

read, err := client.SyncManagedDeviceDevice(ctx, id, manageddevice.DefaultSyncManagedDeviceDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateManagedDevice(ctx, id, payload, manageddevice.DefaultUpdateManagedDeviceOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.UpdateManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateManagedDeviceWindowsDeviceAccount(ctx, id, payload, manageddevice.DefaultUpdateManagedDeviceWindowsDeviceAccountOperationOptions())
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceId")

payload := manageddevice.WipeManagedDeviceRequest{
	// ...
}


read, err := client.WipeManagedDevice(ctx, id, payload, manageddevice.DefaultWipeManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
