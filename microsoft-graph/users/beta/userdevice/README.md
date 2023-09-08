
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userdevice` Documentation

The `userdevice` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userdevice"
```


### Client Initialization

```go
client := userdevice.NewUserDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserDeviceClient.CheckUserByIdDeviceByIdMemberGroup`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

payload := userdevice.CheckUserByIdDeviceByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckUserByIdDeviceByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.CheckUserByIdDeviceByIdMemberObject`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

payload := userdevice.CheckUserByIdDeviceByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckUserByIdDeviceByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.CreateUserByIdDevice`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

payload := userdevice.Device{
	// ...
}


read, err := client.CreateUserByIdDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.DeleteUserByIdDeviceById`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

read, err := client.DeleteUserByIdDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDeviceById`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

read, err := client.GetUserByIdDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDeviceByIdMemberGroup`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

payload := userdevice.GetUserByIdDeviceByIdMemberGroupRequest{
	// ...
}


read, err := client.GetUserByIdDeviceByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDeviceByIdMemberObject`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

payload := userdevice.GetUserByIdDeviceByIdMemberObjectRequest{
	// ...
}


read, err := client.GetUserByIdDeviceByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDeviceCount`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

read, err := client.GetUserByIdDeviceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDevicesByIds`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

// alternatively `client.GetUserByIdDevicesByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdDevicesByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserDeviceClient.GetUserByIdDevicesUserOwnedObject`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

payload := userdevice.GetUserByIdDevicesUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserByIdDevicesUserOwnedObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.ListUserByIdDevices`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

// alternatively `client.ListUserByIdDevices(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdDevicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserDeviceClient.RestoreUserByIdDeviceById`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

read, err := client.RestoreUserByIdDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.UpdateUserByIdDeviceById`

```go
ctx := context.TODO()
id := userdevice.NewUserDeviceID("userIdValue", "deviceIdValue")

payload := userdevice.Device{
	// ...
}


read, err := client.UpdateUserByIdDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserDeviceClient.ValidateUserByIdDevicesProperty`

```go
ctx := context.TODO()
id := userdevice.NewUserID("userIdValue")

payload := userdevice.ValidateUserByIdDevicesPropertyRequest{
	// ...
}


read, err := client.ValidateUserByIdDevicesProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
