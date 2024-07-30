
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androiddeviceownerenrollmentprofile` Documentation

The `androiddeviceownerenrollmentprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androiddeviceownerenrollmentprofile"
```


### Client Initialization

```go
client := androiddeviceownerenrollmentprofile.NewAndroidDeviceOwnerEnrollmentProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.CreateAndroidDeviceOwnerEnrollmentProfile`

```go
ctx := context.TODO()

payload := androiddeviceownerenrollmentprofile.AndroidDeviceOwnerEnrollmentProfile{
	// ...
}


read, err := client.CreateAndroidDeviceOwnerEnrollmentProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.CreateAndroidDeviceOwnerEnrollmentProfileCreateToken`

```go
ctx := context.TODO()
id := androiddeviceownerenrollmentprofile.NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileIdValue")

payload := androiddeviceownerenrollmentprofile.CreateAndroidDeviceOwnerEnrollmentProfileCreateTokenRequest{
	// ...
}


read, err := client.CreateAndroidDeviceOwnerEnrollmentProfileCreateToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.CreateAndroidDeviceOwnerEnrollmentProfileRevokeToken`

```go
ctx := context.TODO()
id := androiddeviceownerenrollmentprofile.NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileIdValue")

read, err := client.CreateAndroidDeviceOwnerEnrollmentProfileRevokeToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.DeleteAndroidDeviceOwnerEnrollmentProfile`

```go
ctx := context.TODO()
id := androiddeviceownerenrollmentprofile.NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileIdValue")

read, err := client.DeleteAndroidDeviceOwnerEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.GetAndroidDeviceOwnerEnrollmentProfile`

```go
ctx := context.TODO()
id := androiddeviceownerenrollmentprofile.NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileIdValue")

read, err := client.GetAndroidDeviceOwnerEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.GetAndroidDeviceOwnerEnrollmentProfileCount`

```go
ctx := context.TODO()


read, err := client.GetAndroidDeviceOwnerEnrollmentProfileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.ListAndroidDeviceOwnerEnrollmentProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListAndroidDeviceOwnerEnrollmentProfiles(ctx)` can be used to do batched pagination
items, err := client.ListAndroidDeviceOwnerEnrollmentProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AndroidDeviceOwnerEnrollmentProfileClient.UpdateAndroidDeviceOwnerEnrollmentProfile`

```go
ctx := context.TODO()
id := androiddeviceownerenrollmentprofile.NewDeviceManagementAndroidDeviceOwnerEnrollmentProfileID("androidDeviceOwnerEnrollmentProfileIdValue")

payload := androiddeviceownerenrollmentprofile.AndroidDeviceOwnerEnrollmentProfile{
	// ...
}


read, err := client.UpdateAndroidDeviceOwnerEnrollmentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
