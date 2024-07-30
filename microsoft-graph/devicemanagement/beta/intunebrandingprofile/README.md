
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intunebrandingprofile` Documentation

The `intunebrandingprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intunebrandingprofile"
```


### Client Initialization

```go
client := intunebrandingprofile.NewIntuneBrandingProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntuneBrandingProfileClient.AssignDeviceManagementIntuneBrandingProfile`

```go
ctx := context.TODO()
id := intunebrandingprofile.NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue")

payload := intunebrandingprofile.AssignDeviceManagementIntuneBrandingProfileRequest{
	// ...
}


read, err := client.AssignDeviceManagementIntuneBrandingProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntuneBrandingProfileClient.CreateIntuneBrandingProfile`

```go
ctx := context.TODO()

payload := intunebrandingprofile.IntuneBrandingProfile{
	// ...
}


read, err := client.CreateIntuneBrandingProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntuneBrandingProfileClient.DeleteIntuneBrandingProfile`

```go
ctx := context.TODO()
id := intunebrandingprofile.NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue")

read, err := client.DeleteIntuneBrandingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntuneBrandingProfileClient.GetIntuneBrandingProfile`

```go
ctx := context.TODO()
id := intunebrandingprofile.NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue")

read, err := client.GetIntuneBrandingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntuneBrandingProfileClient.GetIntuneBrandingProfileCount`

```go
ctx := context.TODO()


read, err := client.GetIntuneBrandingProfileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntuneBrandingProfileClient.ListIntuneBrandingProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListIntuneBrandingProfiles(ctx)` can be used to do batched pagination
items, err := client.ListIntuneBrandingProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntuneBrandingProfileClient.UpdateIntuneBrandingProfile`

```go
ctx := context.TODO()
id := intunebrandingprofile.NewDeviceManagementIntuneBrandingProfileID("intuneBrandingProfileIdValue")

payload := intunebrandingprofile.IntuneBrandingProfile{
	// ...
}


read, err := client.UpdateIntuneBrandingProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
