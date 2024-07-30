
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkenrollmentprofile` Documentation

The `androidforworkenrollmentprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkenrollmentprofile"
```


### Client Initialization

```go
client := androidforworkenrollmentprofile.NewAndroidForWorkEnrollmentProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.CreateAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()

payload := androidforworkenrollmentprofile.AndroidForWorkEnrollmentProfile{
	// ...
}


read, err := client.CreateAndroidForWorkEnrollmentProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.CreateAndroidForWorkEnrollmentProfileCreateToken`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileIdValue")

payload := androidforworkenrollmentprofile.CreateAndroidForWorkEnrollmentProfileCreateTokenRequest{
	// ...
}


read, err := client.CreateAndroidForWorkEnrollmentProfileCreateToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.CreateAndroidForWorkEnrollmentProfileRevokeToken`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileIdValue")

read, err := client.CreateAndroidForWorkEnrollmentProfileRevokeToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.DeleteAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileIdValue")

read, err := client.DeleteAndroidForWorkEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.GetAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileIdValue")

read, err := client.GetAndroidForWorkEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.GetAndroidForWorkEnrollmentProfileCount`

```go
ctx := context.TODO()


read, err := client.GetAndroidForWorkEnrollmentProfileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.ListAndroidForWorkEnrollmentProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListAndroidForWorkEnrollmentProfiles(ctx)` can be used to do batched pagination
items, err := client.ListAndroidForWorkEnrollmentProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.UpdateAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileIdValue")

payload := androidforworkenrollmentprofile.AndroidForWorkEnrollmentProfile{
	// ...
}


read, err := client.UpdateAndroidForWorkEnrollmentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
