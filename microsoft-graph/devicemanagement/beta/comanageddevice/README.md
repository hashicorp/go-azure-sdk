
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevice` Documentation

The `comanageddevice` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevice"
```


### Client Initialization

```go
client := comanageddevice.NewComanagedDeviceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComanagedDeviceClient.BypassComanagedDeviceActivationLock`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.BypassComanagedDeviceActivationLock(ctx, id, comanageddevice.DefaultBypassComanagedDeviceActivationLockOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ChangeComanagedDeviceAssignments`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.ChangeComanagedDeviceAssignmentsRequest{
	// ...
}


read, err := client.ChangeComanagedDeviceAssignments(ctx, id, payload, comanageddevice.DefaultChangeComanagedDeviceAssignmentsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CleanComanagedDeviceWindowsDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CleanComanagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanComanagedDeviceWindowsDevice(ctx, id, payload, comanageddevice.DefaultCleanComanagedDeviceWindowsDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDevice`

```go
ctx := context.TODO()

payload := comanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateComanagedDevice(ctx, payload, comanageddevice.DefaultCreateComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceActivateDeviceEsim`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateComanagedDeviceActivateDeviceEsim(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceActivateDeviceEsimOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDeprovision`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceDeprovisionRequest{
	// ...
}


read, err := client.CreateComanagedDeviceDeprovision(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceDeprovisionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDownloadAppDiagnostic`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceDownloadAppDiagnosticRequest{
	// ...
}


read, err := client.CreateComanagedDeviceDownloadAppDiagnostic(ctx, payload, comanageddevice.DefaultCreateComanagedDeviceDownloadAppDiagnosticOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDownloadPowerliftAppDiagnostic`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceDownloadPowerliftAppDiagnosticRequest{
	// ...
}


read, err := client.CreateComanagedDeviceDownloadPowerliftAppDiagnostic(ctx, payload, comanageddevice.DefaultCreateComanagedDeviceDownloadPowerliftAppDiagnosticOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceEnrollNowAction`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceEnrollNowAction(ctx, id, comanageddevice.DefaultCreateComanagedDeviceEnrollNowActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceExecuteAction`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceExecuteActionRequest{
	// ...
}


read, err := client.CreateComanagedDeviceExecuteAction(ctx, payload, comanageddevice.DefaultCreateComanagedDeviceExecuteActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceInitiateDeviceAttestation`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceInitiateDeviceAttestation(ctx, id, comanageddevice.DefaultCreateComanagedDeviceInitiateDeviceAttestationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceInitiateMobileDeviceManagementKeyRecovery`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceInitiateMobileDeviceManagementKeyRecovery(ctx, id, comanageddevice.DefaultCreateComanagedDeviceInitiateMobileDeviceManagementKeyRecoveryOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceInitiateOnDemandProactiveRemediation`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateComanagedDeviceInitiateOnDemandProactiveRemediation(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceInitiateOnDemandProactiveRemediationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateComanagedDeviceLogCollectionRequest(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceLogCollectionRequestOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceOverrideComplianceState`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateComanagedDeviceOverrideComplianceState(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceOverrideComplianceStateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDevicePlayLostModeSound`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDevicePlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateComanagedDevicePlayLostModeSound(ctx, id, payload, comanageddevice.DefaultCreateComanagedDevicePlayLostModeSoundOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceReenable`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceReenable(ctx, id, comanageddevice.DefaultCreateComanagedDeviceReenableOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRotateBitLockerKey`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceRotateBitLockerKey(ctx, id, comanageddevice.DefaultCreateComanagedDeviceRotateBitLockerKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRotateFileVaultKey`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceRotateFileVaultKey(ctx, id, comanageddevice.DefaultCreateComanagedDeviceRotateFileVaultKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRotateLocalAdminPassword`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceRotateLocalAdminPassword(ctx, id, comanageddevice.DefaultCreateComanagedDeviceRotateLocalAdminPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceTriggerConfigurationManagerAction`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateComanagedDeviceTriggerConfigurationManagerAction(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceTriggerConfigurationManagerActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceWindowsDefenderScan`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.CreateComanagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateComanagedDeviceWindowsDefenderScan(ctx, id, payload, comanageddevice.DefaultCreateComanagedDeviceWindowsDefenderScanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.CreateComanagedDeviceWindowsDefenderUpdateSignature(ctx, id, comanageddevice.DefaultCreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.DeleteComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.DeleteComanagedDevice(ctx, id, comanageddevice.DefaultDeleteComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.DeleteComanagedDeviceUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.DeleteComanagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteComanagedDeviceUserFromSharedAppleDevice(ctx, id, payload, comanageddevice.DefaultDeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.DisableComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.DisableComanagedDevice(ctx, id, comanageddevice.DefaultDisableComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.DisableComanagedDeviceLostMode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.DisableComanagedDeviceLostMode(ctx, id, comanageddevice.DefaultDisableComanagedDeviceLostModeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.EnableComanagedDeviceLostMode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.EnableComanagedDeviceLostModeRequest{
	// ...
}


read, err := client.EnableComanagedDeviceLostMode(ctx, id, payload, comanageddevice.DefaultEnableComanagedDeviceLostModeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.GetComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.GetComanagedDevice(ctx, id, comanageddevice.DefaultGetComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.GetComanagedDevicesCount`

```go
ctx := context.TODO()


read, err := client.GetComanagedDevicesCount(ctx, comanageddevice.DefaultGetComanagedDevicesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ListComanagedDevices`

```go
ctx := context.TODO()


// alternatively `client.ListComanagedDevices(ctx, comanageddevice.DefaultListComanagedDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListComanagedDevicesComplete(ctx, comanageddevice.DefaultListComanagedDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComanagedDeviceClient.LocateComanagedDeviceDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.LocateComanagedDeviceDevice(ctx, id, comanageddevice.DefaultLocateComanagedDeviceDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.LogoutComanagedDeviceSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.LogoutComanagedDeviceSharedAppleDeviceActiveUser(ctx, id, comanageddevice.DefaultLogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.MoveComanagedDevicesToOU`

```go
ctx := context.TODO()

payload := comanageddevice.MoveComanagedDevicesToOURequest{
	// ...
}


read, err := client.MoveComanagedDevicesToOU(ctx, payload, comanageddevice.DefaultMoveComanagedDevicesToOUOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.PauseComanagedDeviceConfigurationRefresh`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.PauseComanagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseComanagedDeviceConfigurationRefresh(ctx, id, payload, comanageddevice.DefaultPauseComanagedDeviceConfigurationRefreshOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RebootComanagedDeviceNow`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RebootComanagedDeviceNow(ctx, id, comanageddevice.DefaultRebootComanagedDeviceNowOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RecoverComanagedDevicePasscode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RecoverComanagedDevicePasscode(ctx, id, comanageddevice.DefaultRecoverComanagedDevicePasscodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RemoteLockComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RemoteLockComanagedDevice(ctx, id, comanageddevice.DefaultRemoteLockComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RemoveComanagedDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RemoveComanagedDeviceFirmwareConfigurationInterfaceManagement(ctx, id, comanageddevice.DefaultRemoveComanagedDeviceFirmwareConfigurationInterfaceManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RequestComanagedDeviceRemoteAssistance`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RequestComanagedDeviceRemoteAssistance(ctx, id, comanageddevice.DefaultRequestComanagedDeviceRemoteAssistanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ResetComanagedDevicePasscode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.ResetComanagedDevicePasscode(ctx, id, comanageddevice.DefaultResetComanagedDevicePasscodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RetireComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RetireComanagedDevice(ctx, id, comanageddevice.DefaultRetireComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RevokeComanagedDeviceAppleVppLicenses`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.RevokeComanagedDeviceAppleVppLicenses(ctx, id, comanageddevice.DefaultRevokeComanagedDeviceAppleVppLicensesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SendComanagedDeviceCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.SendComanagedDeviceCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.SendComanagedDeviceCustomNotificationToCompanyPortal(ctx, id, payload, comanageddevice.DefaultSendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SetComanagedDeviceName`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.SetComanagedDeviceNameRequest{
	// ...
}


read, err := client.SetComanagedDeviceName(ctx, id, payload, comanageddevice.DefaultSetComanagedDeviceNameOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ShutDownComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.ShutDownComanagedDevice(ctx, id, comanageddevice.DefaultShutDownComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SyncComanagedDeviceDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

read, err := client.SyncComanagedDeviceDevice(ctx, id, comanageddevice.DefaultSyncComanagedDeviceDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.UpdateComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateComanagedDevice(ctx, id, payload, comanageddevice.DefaultUpdateComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.UpdateComanagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.UpdateComanagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateComanagedDeviceWindowsDeviceAccount(ctx, id, payload, comanageddevice.DefaultUpdateComanagedDeviceWindowsDeviceAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.WipeComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceId")

payload := comanageddevice.WipeComanagedDeviceRequest{
	// ...
}


read, err := client.WipeComanagedDevice(ctx, id, payload, comanageddevice.DefaultWipeComanagedDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
