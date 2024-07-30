
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingenrollmentprofile` Documentation

The `deponboardingsettingenrollmentprofile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingenrollmentprofile"
```


### Client Initialization

```go
client := deponboardingsettingenrollmentprofile.NewDepOnboardingSettingEnrollmentProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.CreateDepOnboardingSettingEnrollmentProfile`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

payload := deponboardingsettingenrollmentprofile.EnrollmentProfile{
	// ...
}


read, err := client.CreateDepOnboardingSettingEnrollmentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.DeleteDepOnboardingSettingEnrollmentProfile`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

read, err := client.DeleteDepOnboardingSettingEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.GetDepOnboardingSettingEnrollmentProfile`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

read, err := client.GetDepOnboardingSettingEnrollmentProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.GetDepOnboardingSettingEnrollmentProfileCount`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

read, err := client.GetDepOnboardingSettingEnrollmentProfileCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.ListDepOnboardingSettingEnrollmentProfiles`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

// alternatively `client.ListDepOnboardingSettingEnrollmentProfiles(ctx, id)` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingEnrollmentProfilesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.SetDeviceManagementDepOnboardingSettingEnrollmentProfileDefaultProfile`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

read, err := client.SetDeviceManagementDepOnboardingSettingEnrollmentProfileDefaultProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.UpdateDepOnboardingSettingEnrollmentProfile`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

payload := deponboardingsettingenrollmentprofile.EnrollmentProfile{
	// ...
}


read, err := client.UpdateDepOnboardingSettingEnrollmentProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingEnrollmentProfileClient.UpdateDeviceManagementDepOnboardingSettingEnrollmentProfileDeviceProfileAssignment`

```go
ctx := context.TODO()
id := deponboardingsettingenrollmentprofile.NewDeviceManagementDepOnboardingSettingIdEnrollmentProfileID("depOnboardingSettingIdValue", "enrollmentProfileIdValue")

payload := deponboardingsettingenrollmentprofile.UpdateDeviceManagementDepOnboardingSettingEnrollmentProfileDeviceProfileAssignmentRequest{
	// ...
}


read, err := client.UpdateDeviceManagementDepOnboardingSettingEnrollmentProfileDeviceProfileAssignment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
