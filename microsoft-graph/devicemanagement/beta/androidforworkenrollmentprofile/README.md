
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkenrollmentprofile` Documentation

The `androidforworkenrollmentprofile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/androidforworkenrollmentprofile"
```


### Client Initialization

```go
client := androidforworkenrollmentprofile.NewAndroidForWorkEnrollmentProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.CreateAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()

payload := androidforworkenrollmentprofile.AndroidForWorkEnrollmentProfile{
	// ...
}


read, err := client.CreateAndroidForWorkEnrollmentProfile(ctx, payload, androidforworkenrollmentprofile.DefaultCreateAndroidForWorkEnrollmentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.CreateAndroidForWorkEnrollmentProfileToken`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

payload := androidforworkenrollmentprofile.CreateAndroidForWorkEnrollmentProfileTokenRequest{
	// ...
}


read, err := client.CreateAndroidForWorkEnrollmentProfileToken(ctx, id, payload, androidforworkenrollmentprofile.DefaultCreateAndroidForWorkEnrollmentProfileTokenOperationOptions())
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
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

read, err := client.DeleteAndroidForWorkEnrollmentProfile(ctx, id, androidforworkenrollmentprofile.DefaultDeleteAndroidForWorkEnrollmentProfileOperationOptions())
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
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

read, err := client.GetAndroidForWorkEnrollmentProfile(ctx, id, androidforworkenrollmentprofile.DefaultGetAndroidForWorkEnrollmentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.GetAndroidForWorkEnrollmentProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetAndroidForWorkEnrollmentProfilesCount(ctx, androidforworkenrollmentprofile.DefaultGetAndroidForWorkEnrollmentProfilesCountOperationOptions())
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


// alternatively `client.ListAndroidForWorkEnrollmentProfiles(ctx, androidforworkenrollmentprofile.DefaultListAndroidForWorkEnrollmentProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListAndroidForWorkEnrollmentProfilesComplete(ctx, androidforworkenrollmentprofile.DefaultListAndroidForWorkEnrollmentProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.RevokeAndroidForWorkEnrollmentProfileToken`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

read, err := client.RevokeAndroidForWorkEnrollmentProfileToken(ctx, id, androidforworkenrollmentprofile.DefaultRevokeAndroidForWorkEnrollmentProfileTokenOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AndroidForWorkEnrollmentProfileClient.UpdateAndroidForWorkEnrollmentProfile`

```go
ctx := context.TODO()
id := androidforworkenrollmentprofile.NewDeviceManagementAndroidForWorkEnrollmentProfileID("androidForWorkEnrollmentProfileId")

payload := androidforworkenrollmentprofile.AndroidForWorkEnrollmentProfile{
	// ...
}


read, err := client.UpdateAndroidForWorkEnrollmentProfile(ctx, id, payload, androidforworkenrollmentprofile.DefaultUpdateAndroidForWorkEnrollmentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
