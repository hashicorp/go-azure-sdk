
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermanageddevice` Documentation

The `usermanageddevice` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermanageddevice"
```


### Client Initialization

```go
client := usermanageddevice.NewUserManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateUserByIdManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceBulkReprovisionCloudPC`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceBulkReprovisionCloudPCRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceBulkReprovisionCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceBulkRestoreCloudPC`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceBulkRestoreCloudPCRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceBulkRestoreCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceBulkSetCloudPCReviewStatu`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceBulkSetCloudPCReviewStatuRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceBulkSetCloudPCReviewStatu(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdActivateDeviceEsim`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdActivateDeviceEsim(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdBypassActivationLock`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdCleanWindowsDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdCreateDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdCreateDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdCreateDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdDeleteUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDeprovision`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdDeprovisionRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdDeprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDisable`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdDisable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdDisableLostMode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdDisableLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdEnableLostMode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdEnableLostModeRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdEnableLostMode(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdEnrollNowAction`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdEnrollNowAction(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdInitiateMobileDeviceManagementKeyRecovery`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdInitiateMobileDeviceManagementKeyRecovery(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdInitiateOnDemandProactiveRemediation`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdInitiateOnDemandProactiveRemediation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdLocateDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdOverrideComplianceState`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdOverrideComplianceState(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdPlayLostModeSound`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdPlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdPlayLostModeSound(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRebootNow`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRecoverPasscode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRecoverPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdReenable`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdReenable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRemoteLock`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdReprovisionCloudPC`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdReprovisionCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRequestRemoteAssistance`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdResetPasscode`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdResizeCloudPC`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdResizeCloudPCRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdResizeCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRetire`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRevokeAppleVppLicens`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRevokeAppleVppLicens(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRotateBitLockerKey`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRotateBitLockerKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRotateFileVaultKey`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRotateFileVaultKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdRotateLocalAdminPassword`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdRotateLocalAdminPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdSendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdSendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdSendCustomNotificationToCompanyPortal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdShutDown`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdSyncDevice`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdSyncDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdTriggerConfigurationManagerAction`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdTriggerConfigurationManagerAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWindowsDefenderScan`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdWindowsDefenderScan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.CreateUserByIdManagedDeviceByIdWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceByIdWipe`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceByIdWipeRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceByIdWipe(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceDownloadAppDiagnostic`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceDownloadAppDiagnosticRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceDownloadAppDiagnostic(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceExecuteAction`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceExecuteActionRequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceExecuteAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.CreateUserByIdManagedDeviceMoveDevicesToOU`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

payload := usermanageddevice.CreateUserByIdManagedDeviceMoveDevicesToOURequest{
	// ...
}


read, err := client.CreateUserByIdManagedDeviceMoveDevicesToOU(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.DeleteUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.DeleteUserByIdManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.GetUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.GetUserByIdManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.GetUserByIdManagedDeviceCount`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

read, err := client.GetUserByIdManagedDeviceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.ListUserByIdManagedDevices`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserID("userIdValue")

// alternatively `client.ListUserByIdManagedDevices(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdManagedDevicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserManagedDeviceClient.RemoveUserByIdManagedDeviceByIdDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

read, err := client.RemoveUserByIdManagedDeviceByIdDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.RestoreUserByIdManagedDeviceByIdCloudPC`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.RestoreUserByIdManagedDeviceByIdCloudPCRequest{
	// ...
}


read, err := client.RestoreUserByIdManagedDeviceByIdCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.SetUserByIdManagedDeviceByIdCloudPCReviewStatu`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.SetUserByIdManagedDeviceByIdCloudPCReviewStatuRequest{
	// ...
}


read, err := client.SetUserByIdManagedDeviceByIdCloudPCReviewStatu(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.SetUserByIdManagedDeviceByIdDeviceName`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.SetUserByIdManagedDeviceByIdDeviceNameRequest{
	// ...
}


read, err := client.SetUserByIdManagedDeviceByIdDeviceName(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.UpdateUserByIdManagedDeviceById`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateUserByIdManagedDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserManagedDeviceClient.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccount`

```go
ctx := context.TODO()
id := usermanageddevice.NewUserManagedDeviceID("userIdValue", "managedDeviceIdValue")

payload := usermanageddevice.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateUserByIdManagedDeviceByIdWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
