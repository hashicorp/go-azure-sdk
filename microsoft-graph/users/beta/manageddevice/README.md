
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevice` Documentation

The `manageddevice` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevice"
```


### Client Initialization

```go
client := manageddevice.NewManagedDeviceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceClient.BypassManagedDeviceActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.BypassManagedDeviceActivationLock(ctx, id, manageddevice.DefaultBypassManagedDeviceActivationLockOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ChangeManagedDeviceAssignments`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.ChangeManagedDeviceAssignmentsRequest{
	// ...
}


read, err := client.ChangeManagedDeviceAssignments(ctx, id, payload, manageddevice.DefaultChangeManagedDeviceAssignmentsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CleanManagedDeviceWindowsDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CleanManagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanManagedDeviceWindowsDevice(ctx, id, payload, manageddevice.DefaultCleanManagedDeviceWindowsDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateManagedDevice(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceActivateDeviceEsim`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateManagedDeviceActivateDeviceEsim(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceActivateDeviceEsimOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBulkReprovisionCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.CreateManagedDeviceBulkReprovisionCloudPCRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkReprovisionCloudPC(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceBulkReprovisionCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBulkRestoreCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.CreateManagedDeviceBulkRestoreCloudPCRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkRestoreCloudPC(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceBulkRestoreCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBulkSetCloudPCReviewStatus`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.CreateManagedDeviceBulkSetCloudPCReviewStatusRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkSetCloudPCReviewStatus(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceBulkSetCloudPCReviewStatusOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDeprovision`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceDeprovisionRequest{
	// ...
}


read, err := client.CreateManagedDeviceDeprovision(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceDeprovisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDownloadAppDiagnostic`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.CreateManagedDeviceDownloadAppDiagnosticRequest{
	// ...
}


read, err := client.CreateManagedDeviceDownloadAppDiagnostic(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceDownloadAppDiagnosticOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceEnrollNowAction`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceEnrollNowAction(ctx, id, manageddevice.DefaultCreateManagedDeviceEnrollNowActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceExecuteAction`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.CreateManagedDeviceExecuteActionRequest{
	// ...
}


read, err := client.CreateManagedDeviceExecuteAction(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceExecuteActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceInitiateDeviceAttestation`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceInitiateDeviceAttestation(ctx, id, manageddevice.DefaultCreateManagedDeviceInitiateDeviceAttestationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceInitiateMobileDeviceManagementKeyRecovery`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceInitiateMobileDeviceManagementKeyRecovery(ctx, id, manageddevice.DefaultCreateManagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceInitiateOnDemandProactiveRemediation`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateManagedDeviceInitiateOnDemandProactiveRemediation(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceInitiateOnDemandProactiveRemediationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateManagedDeviceLogCollectionRequest(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceOverrideComplianceState`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateManagedDeviceOverrideComplianceState(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceOverrideComplianceStateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDevicePlayLostModeSound`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDevicePlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateManagedDevicePlayLostModeSound(ctx, id, payload, manageddevice.DefaultCreateManagedDevicePlayLostModeSoundOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceReenable`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceReenable(ctx, id, manageddevice.DefaultCreateManagedDeviceReenableOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRotateBitLockerKey`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceRotateBitLockerKey(ctx, id, manageddevice.DefaultCreateManagedDeviceRotateBitLockerKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRotateFileVaultKey`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceRotateFileVaultKey(ctx, id, manageddevice.DefaultCreateManagedDeviceRotateFileVaultKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRotateLocalAdminPassword`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceRotateLocalAdminPassword(ctx, id, manageddevice.DefaultCreateManagedDeviceRotateLocalAdminPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceTriggerConfigurationManagerAction`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateManagedDeviceTriggerConfigurationManagerAction(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceTriggerConfigurationManagerActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWindowsDefenderScan`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.CreateManagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateManagedDeviceWindowsDefenderScan(ctx, id, payload, manageddevice.DefaultCreateManagedDeviceWindowsDefenderScanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.CreateManagedDeviceWindowsDefenderUpdateSignature(ctx, id, manageddevice.DefaultCreateManagedDeviceWindowsDefenderUpdateSignatureOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DeleteManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.DeleteManagedDevice(ctx, id, manageddevice.DefaultDeleteManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DeleteManagedDeviceUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.DeleteManagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteManagedDeviceUserFromSharedAppleDevice(ctx, id, payload, manageddevice.DefaultDeleteManagedDeviceUserFromSharedAppleDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DisableManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.DisableManagedDevice(ctx, id, manageddevice.DefaultDisableManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.DisableManagedDeviceLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.DisableManagedDeviceLostMode(ctx, id, manageddevice.DefaultDisableManagedDeviceLostModeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.EnableManagedDeviceLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.EnableManagedDeviceLostModeRequest{
	// ...
}


read, err := client.EnableManagedDeviceLostMode(ctx, id, payload, manageddevice.DefaultEnableManagedDeviceLostModeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.GetManagedDevice(ctx, id, manageddevice.DefaultGetManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDevicesCount`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

read, err := client.GetManagedDevicesCount(ctx, id, manageddevice.DefaultGetManagedDevicesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ListManagedDevices`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

// alternatively `client.ListManagedDevices(ctx, id, manageddevice.DefaultListManagedDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListManagedDevicesComplete(ctx, id, manageddevice.DefaultListManagedDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceClient.LocateManagedDeviceDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.LocateManagedDeviceDevice(ctx, id, manageddevice.DefaultLocateManagedDeviceDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.LogoutManagedDeviceSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.LogoutManagedDeviceSharedAppleDeviceActiveUser(ctx, id, manageddevice.DefaultLogoutManagedDeviceSharedAppleDeviceActiveUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.MoveManagedDevicesToOU`

```go
ctx := context.TODO()
id := manageddevice.NewUserID("userId")

payload := manageddevice.MoveManagedDevicesToOURequest{
	// ...
}


read, err := client.MoveManagedDevicesToOU(ctx, id, payload, manageddevice.DefaultMoveManagedDevicesToOUOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.PauseManagedDeviceConfigurationRefresh`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.PauseManagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseManagedDeviceConfigurationRefresh(ctx, id, payload, manageddevice.DefaultPauseManagedDeviceConfigurationRefreshOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RebootManagedDeviceNow`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RebootManagedDeviceNow(ctx, id, manageddevice.DefaultRebootManagedDeviceNowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RecoverManagedDevicePasscode`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RecoverManagedDevicePasscode(ctx, id, manageddevice.DefaultRecoverManagedDevicePasscodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RemoteLockManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RemoteLockManagedDevice(ctx, id, manageddevice.DefaultRemoteLockManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RemoveManagedDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RemoveManagedDeviceFirmwareConfigurationInterfaceManagement(ctx, id, manageddevice.DefaultRemoveManagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ReprovisionManagedDeviceCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.ReprovisionManagedDeviceCloudPC(ctx, id, manageddevice.DefaultReprovisionManagedDeviceCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RequestManagedDeviceRemoteAssistance`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RequestManagedDeviceRemoteAssistance(ctx, id, manageddevice.DefaultRequestManagedDeviceRemoteAssistanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ResetManagedDevicePasscode`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.ResetManagedDevicePasscode(ctx, id, manageddevice.DefaultResetManagedDevicePasscodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ResizeManagedDeviceCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.ResizeManagedDeviceCloudPCRequest{
	// ...
}


read, err := client.ResizeManagedDeviceCloudPC(ctx, id, payload, manageddevice.DefaultResizeManagedDeviceCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RestoreManagedDeviceCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.RestoreManagedDeviceCloudPCRequest{
	// ...
}


read, err := client.RestoreManagedDeviceCloudPC(ctx, id, payload, manageddevice.DefaultRestoreManagedDeviceCloudPCOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RetireManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RetireManagedDevice(ctx, id, manageddevice.DefaultRetireManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RevokeManagedDeviceAppleVppLicenses`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.RevokeManagedDeviceAppleVppLicenses(ctx, id, manageddevice.DefaultRevokeManagedDeviceAppleVppLicensesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SendManagedDeviceCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.SendManagedDeviceCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.SendManagedDeviceCustomNotificationToCompanyPortal(ctx, id, payload, manageddevice.DefaultSendManagedDeviceCustomNotificationToCompanyPortalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SetManagedDeviceCloudPCReviewStatus`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.SetManagedDeviceCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetManagedDeviceCloudPCReviewStatus(ctx, id, payload, manageddevice.DefaultSetManagedDeviceCloudPCReviewStatusOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SetManagedDeviceName`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.SetManagedDeviceNameRequest{
	// ...
}


read, err := client.SetManagedDeviceName(ctx, id, payload, manageddevice.DefaultSetManagedDeviceNameOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ShutDownManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.ShutDownManagedDevice(ctx, id, manageddevice.DefaultShutDownManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SyncManagedDeviceDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

read, err := client.SyncManagedDeviceDevice(ctx, id, manageddevice.DefaultSyncManagedDeviceDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.UpdateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateManagedDevice(ctx, id, payload, manageddevice.DefaultUpdateManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.UpdateManagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.UpdateManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateManagedDeviceWindowsDeviceAccount(ctx, id, payload, manageddevice.DefaultUpdateManagedDeviceWindowsDeviceAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.WipeManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewUserIdManagedDeviceID("userId", "managedDeviceId")

payload := manageddevice.WipeManagedDeviceRequest{
	// ...
}


read, err := client.WipeManagedDevice(ctx, id, payload, manageddevice.DefaultWipeManagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
