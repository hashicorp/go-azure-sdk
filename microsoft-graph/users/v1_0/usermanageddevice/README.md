
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermanageddevice` Documentation

The `usermanageddevice` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermanageddevice"
```


### Client Initialization

```go
client := usermanageddevice.NewUserManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateUserByIdManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdBypassActivationLock`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdCleanWindowsDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDisableLostMode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdDisableLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdLocateDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRebootNow`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRecoverPasscode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRecoverPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRemoteLock`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRequestRemoteAssistance`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdResetPasscode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRetire`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdShutDown`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdSyncDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdSyncDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWindowsDefenderScan`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdWindowsDefenderScan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWipe`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdWipeRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdWipe(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.DeleteUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.DeleteUserByIdManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.GetUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.GetUserByIdManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.GetUserByIdManagedDeviceCount`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

read, err := client.GetUserByIdManagedDeviceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.ListUserByIdManagedDevices`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

// alternatively `client.ListUserByIdManagedDevices(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdManagedDevicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserManagedDeviceClient.UpdateUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateUserByIdManagedDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccount`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
