
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memanageddevice` Documentation

The `memanageddevice` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memanageddevice"
```


### Client Initialization

```go
client := memanageddevice.NewMeManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDevice`

```go
ctx := context.TODO()

payload := memanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateMeManagedDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdBypassActivationLock`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdCleanWindowsDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDisableLostMode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdDisableLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdLocateDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRebootNow`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRecoverPasscode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRecoverPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRemoteLock`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRequestRemoteAssistance`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdResetPasscode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRetire`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdShutDown`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdSyncDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdSyncDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWindowsDefenderScan`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdWindowsDefenderScan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWipe`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdWipeRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdWipe(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.DeleteMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteMeManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.GetMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.GetMeManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.GetMeManagedDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetMeManagedDeviceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.ListMeManagedDevices`

```go
ctx := context.TODO()


// alternatively `client.ListMeManagedDevices(ctx)` can be used to do batched pagination
items, err := client.ListMeManagedDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeManagedDeviceClient.UpdateMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateMeManagedDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.UpdateMeManagedDeviceByIdWindowsDeviceAccount`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.UpdateMeManagedDeviceByIdWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateMeManagedDeviceByIdWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
