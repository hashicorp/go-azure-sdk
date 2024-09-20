
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsetting` Documentation

The `deponboardingsetting` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsetting"
```


### Client Initialization

```go
client := deponboardingsetting.NewDepOnboardingSettingClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DepOnboardingSettingClient.CreateDepOnboardingSetting`

```go
ctx := context.TODO()

payload := deponboardingsetting.DepOnboardingSetting{
	// ...
}


read, err := client.CreateDepOnboardingSetting(ctx, payload, deponboardingsetting.DefaultCreateDepOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.CreateDepOnboardingSettingGenerateEncryptionPublicKey`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.CreateDepOnboardingSettingGenerateEncryptionPublicKey(ctx, id, deponboardingsetting.DefaultCreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.CreateDepOnboardingSettingUploadDepToken`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

payload := deponboardingsetting.CreateDepOnboardingSettingUploadDepTokenRequest{
	// ...
}


read, err := client.CreateDepOnboardingSettingUploadDepToken(ctx, id, payload, deponboardingsetting.DefaultCreateDepOnboardingSettingUploadDepTokenOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.DeleteDepOnboardingSetting`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.DeleteDepOnboardingSetting(ctx, id, deponboardingsetting.DefaultDeleteDepOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.GetDepOnboardingSetting`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.GetDepOnboardingSetting(ctx, id, deponboardingsetting.DefaultGetDepOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.GetDepOnboardingSettingsCount`

```go
ctx := context.TODO()


read, err := client.GetDepOnboardingSettingsCount(ctx, deponboardingsetting.DefaultGetDepOnboardingSettingsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.ListDepOnboardingSettings`

```go
ctx := context.TODO()


// alternatively `client.ListDepOnboardingSettings(ctx, deponboardingsetting.DefaultListDepOnboardingSettingsOperationOptions())` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingsComplete(ctx, deponboardingsetting.DefaultListDepOnboardingSettingsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DepOnboardingSettingClient.ShareDepOnboardingSettingForSchoolDataSyncService`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.ShareDepOnboardingSettingForSchoolDataSyncService(ctx, id, deponboardingsetting.DefaultShareDepOnboardingSettingForSchoolDataSyncServiceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.SyncDepOnboardingSettingWithAppleDeviceEnrollmentProgram`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.SyncDepOnboardingSettingWithAppleDeviceEnrollmentProgram(ctx, id, deponboardingsetting.DefaultSyncDepOnboardingSettingWithAppleDeviceEnrollmentProgramOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.UnshareDepOnboardingSettingForSchoolDataSyncService`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

read, err := client.UnshareDepOnboardingSettingForSchoolDataSyncService(ctx, id, deponboardingsetting.DefaultUnshareDepOnboardingSettingForSchoolDataSyncServiceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingClient.UpdateDepOnboardingSetting`

```go
ctx := context.TODO()
id := deponboardingsetting.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingId")

payload := deponboardingsetting.DepOnboardingSetting{
	// ...
}


read, err := client.UpdateDepOnboardingSetting(ctx, id, payload, deponboardingsetting.DefaultUpdateDepOnboardingSettingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
