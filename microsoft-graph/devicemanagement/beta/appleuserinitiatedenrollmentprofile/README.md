
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofile` Documentation

The `appleuserinitiatedenrollmentprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofile"
```


### Client Initialization

```go
client := appleuserinitiatedenrollmentprofile.NewAppleUserInitiatedEnrollmentProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.CreateAppleUserInitiatedEnrollmentProfile`

```go
ctx := context.TODO()

payload := appleuserinitiatedenrollmentprofile.AppleUserInitiatedEnrollmentProfile{
	// ...
}


read, err := client.CreateAppleUserInitiatedEnrollmentProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.DeleteAppleUserInitiatedEnrollmentProfile`

```go
ctx := context.TODO()
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileIdValue")

read, err := client.DeleteAppleUserInitiatedEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.GetAppleUserInitiatedEnrollmentProfile`

```go
ctx := context.TODO()
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileIdValue")

read, err := client.GetAppleUserInitiatedEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.GetAppleUserInitiatedEnrollmentProfileCount`

```go
ctx := context.TODO()


read, err := client.GetAppleUserInitiatedEnrollmentProfileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.ListAppleUserInitiatedEnrollmentProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListAppleUserInitiatedEnrollmentProfiles(ctx)` can be used to do batched pagination
items, err := client.ListAppleUserInitiatedEnrollmentProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.SetDeviceManagementAppleUserInitiatedEnrollmentProfilePriority`

```go
ctx := context.TODO()
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileIdValue")

payload := appleuserinitiatedenrollmentprofile.SetDeviceManagementAppleUserInitiatedEnrollmentProfilePriorityRequest{
	// ...
}


read, err := client.SetDeviceManagementAppleUserInitiatedEnrollmentProfilePriority(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.UpdateAppleUserInitiatedEnrollmentProfile`

```go
ctx := context.TODO()
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileIdValue")

payload := appleuserinitiatedenrollmentprofile.AppleUserInitiatedEnrollmentProfile{
	// ...
}


read, err := client.UpdateAppleUserInitiatedEnrollmentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
