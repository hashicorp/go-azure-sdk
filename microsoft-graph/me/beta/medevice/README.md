
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/medevice` Documentation

The `medevice` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/medevice"
```


### Client Initialization

```go
client := medevice.NewMeDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeDeviceClient.CheckMeDeviceByIdMemberGroup`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

payload := medevice.CheckMeDeviceByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckMeDeviceByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.CheckMeDeviceByIdMemberObject`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

payload := medevice.CheckMeDeviceByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckMeDeviceByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.CreateMeDevice`

```go
ctx := context.TODO()

payload := medevice.Device{
	// ...
}


read, err := client.CreateMeDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.DeleteMeDeviceById`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

read, err := client.DeleteMeDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.GetMeDeviceById`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

read, err := client.GetMeDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.GetMeDeviceByIdMemberGroup`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

payload := medevice.GetMeDeviceByIdMemberGroupRequest{
	// ...
}


read, err := client.GetMeDeviceByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.GetMeDeviceByIdMemberObject`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

payload := medevice.GetMeDeviceByIdMemberObjectRequest{
	// ...
}


read, err := client.GetMeDeviceByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.GetMeDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetMeDeviceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.GetMeDevicesByIds`

```go
ctx := context.TODO()


// alternatively `client.GetMeDevicesByIds(ctx)` can be used to do batched pagination
items, err := client.GetMeDevicesByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeDeviceClient.GetMeDevicesUserOwnedObject`

```go
ctx := context.TODO()

payload := medevice.GetMeDevicesUserOwnedObjectRequest{
	// ...
}


read, err := client.GetMeDevicesUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.ListMeDevices`

```go
ctx := context.TODO()


// alternatively `client.ListMeDevices(ctx)` can be used to do batched pagination
items, err := client.ListMeDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeDeviceClient.RestoreMeDeviceById`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

read, err := client.RestoreMeDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.UpdateMeDeviceById`

```go
ctx := context.TODO()
id := medevice.NewMeDeviceID("deviceIdValue")

payload := medevice.Device{
	// ...
}


read, err := client.UpdateMeDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeDeviceClient.ValidateMeDevicesProperty`

```go
ctx := context.TODO()

payload := medevice.ValidateMeDevicesPropertyRequest{
	// ...
}


read, err := client.ValidateMeDevicesProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
