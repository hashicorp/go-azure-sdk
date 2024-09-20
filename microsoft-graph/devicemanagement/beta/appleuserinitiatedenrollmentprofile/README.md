
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofile` Documentation

The `appleuserinitiatedenrollmentprofile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/appleuserinitiatedenrollmentprofile"
```


### Client Initialization

```go
client := appleuserinitiatedenrollmentprofile.NewAppleUserInitiatedEnrollmentProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.CreateAppleUserInitiatedEnrollmentProfile`

```go
ctx := context.TODO()

payload := appleuserinitiatedenrollmentprofile.AppleUserInitiatedEnrollmentProfile{
	// ...
}


read, err := client.CreateAppleUserInitiatedEnrollmentProfile(ctx, payload, appleuserinitiatedenrollmentprofile.DefaultCreateAppleUserInitiatedEnrollmentProfileOperationOptions())
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
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId")

read, err := client.DeleteAppleUserInitiatedEnrollmentProfile(ctx, id, appleuserinitiatedenrollmentprofile.DefaultDeleteAppleUserInitiatedEnrollmentProfileOperationOptions())
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
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId")

read, err := client.GetAppleUserInitiatedEnrollmentProfile(ctx, id, appleuserinitiatedenrollmentprofile.DefaultGetAppleUserInitiatedEnrollmentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.GetAppleUserInitiatedEnrollmentProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetAppleUserInitiatedEnrollmentProfilesCount(ctx, appleuserinitiatedenrollmentprofile.DefaultGetAppleUserInitiatedEnrollmentProfilesCountOperationOptions())
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


// alternatively `client.ListAppleUserInitiatedEnrollmentProfiles(ctx, appleuserinitiatedenrollmentprofile.DefaultListAppleUserInitiatedEnrollmentProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListAppleUserInitiatedEnrollmentProfilesComplete(ctx, appleuserinitiatedenrollmentprofile.DefaultListAppleUserInitiatedEnrollmentProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AppleUserInitiatedEnrollmentProfileClient.SetAppleUserInitiatedEnrollmentProfilePriority`

```go
ctx := context.TODO()
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId")

payload := appleuserinitiatedenrollmentprofile.SetAppleUserInitiatedEnrollmentProfilePriorityRequest{
	// ...
}


read, err := client.SetAppleUserInitiatedEnrollmentProfilePriority(ctx, id, payload, appleuserinitiatedenrollmentprofile.DefaultSetAppleUserInitiatedEnrollmentProfilePriorityOperationOptions())
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
id := appleuserinitiatedenrollmentprofile.NewDeviceManagementAppleUserInitiatedEnrollmentProfileID("appleUserInitiatedEnrollmentProfileId")

payload := appleuserinitiatedenrollmentprofile.AppleUserInitiatedEnrollmentProfile{
	// ...
}


read, err := client.UpdateAppleUserInitiatedEnrollmentProfile(ctx, id, payload, appleuserinitiatedenrollmentprofile.DefaultUpdateAppleUserInitiatedEnrollmentProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
