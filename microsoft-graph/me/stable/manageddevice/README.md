
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/manageddevice` Documentation

The `manageddevice` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/manageddevice"
```


### Client Initialization

```go
client := manageddevice.NewManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
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


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBypassActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceCleanWindowsDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateManagedDeviceCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateManagedDeviceDeleteUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDisableLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceDisableLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLocateDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRebootNow`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRecoverPasscode`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRecoverPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRemoteLock`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRequestRemoteAssistance`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceResetPasscode`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRetire`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceShutDown`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceSyncDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceSyncDevice(ctx, id)
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

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
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWipe`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceWipeRequest{
	// ...
}


read, err := client.CreateManagedDeviceWipe(ctx, id, payload)
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteManagedDevice(ctx, id)
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
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.GetManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetManagedDeviceCount(ctx)
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


// alternatively `client.ListManagedDevices(ctx)` can be used to do batched pagination
items, err := client.ListManagedDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceClient.UpdateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

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


### Example Usage: `ManagedDeviceClient.UpdateMeManagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := manageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.UpdateMeManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateMeManagedDeviceWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
