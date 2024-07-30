
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevice` Documentation

The `comanageddevice` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/comanageddevice"
```


### Client Initialization

```go
client := comanageddevice.NewComanagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDevice`

```go
ctx := context.TODO()

payload := comanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateComanagedDevice(ctx, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateComanagedDeviceActivateDeviceEsim(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceBulkReprovisionCloudPC`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceBulkReprovisionCloudPCRequest{
	// ...
}


read, err := client.CreateComanagedDeviceBulkReprovisionCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceBulkRestoreCloudPC`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceBulkRestoreCloudPCRequest{
	// ...
}


read, err := client.CreateComanagedDeviceBulkRestoreCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceBulkSetCloudPCReviewStatu`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceBulkSetCloudPCReviewStatuRequest{
	// ...
}


read, err := client.CreateComanagedDeviceBulkSetCloudPCReviewStatu(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceBypassActivationLock`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceCleanWindowsDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateComanagedDeviceCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceCreateDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceCreateDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateComanagedDeviceCreateDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateComanagedDeviceDeleteUserFromSharedAppleDevice(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceDeprovisionRequest{
	// ...
}


read, err := client.CreateComanagedDeviceDeprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDisable`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceDisable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceDisableLostMode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceDisableLostMode(ctx, id)
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


read, err := client.CreateComanagedDeviceDownloadAppDiagnostic(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceEnableLostMode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceEnableLostModeRequest{
	// ...
}


read, err := client.CreateComanagedDeviceEnableLostMode(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceEnrollNowAction(ctx, id)
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


read, err := client.CreateComanagedDeviceExecuteAction(ctx, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceInitiateMobileDeviceManagementKeyRecovery(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateComanagedDeviceInitiateOnDemandProactiveRemediation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceLocateDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceMoveDevicesToOU`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceMoveDevicesToOURequest{
	// ...
}


read, err := client.CreateComanagedDeviceMoveDevicesToOU(ctx, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateComanagedDeviceOverrideComplianceState(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDevicePlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateComanagedDevicePlayLostModeSound(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRebootNow`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRecoverPasscode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRecoverPasscode(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceReenable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRemoteLock`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceReprovisionCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceReprovisionCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRequestRemoteAssistance`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceResetPasscode`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceResizeCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceResizeCloudPCRequest{
	// ...
}


read, err := client.CreateComanagedDeviceResizeCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRetire`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceRevokeAppleVppLicens`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRevokeAppleVppLicens(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRotateBitLockerKey(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRotateFileVaultKey(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceRotateLocalAdminPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceSendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceSendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.CreateComanagedDeviceSendCustomNotificationToCompanyPortal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceShutDown`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceSyncDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceSyncDevice(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateComanagedDeviceTriggerConfigurationManagerAction(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateComanagedDeviceWindowsDefenderScan(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceWipe`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceWipeRequest{
	// ...
}


read, err := client.CreateComanagedDeviceWipe(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteComanagedDevice(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.GetComanagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.GetComanagedDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetComanagedDeviceCount(ctx)
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


// alternatively `client.ListComanagedDevices(ctx)` can be used to do batched pagination
items, err := client.ListComanagedDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComanagedDeviceClient.PauseDeviceManagementComanagedDeviceConfigurationRefresh`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.PauseDeviceManagementComanagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseDeviceManagementComanagedDeviceConfigurationRefresh(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RemoveDeviceManagementComanagedDeviceDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RemoveDeviceManagementComanagedDeviceDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RestoreDeviceManagementComanagedDeviceCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.RestoreDeviceManagementComanagedDeviceCloudPCRequest{
	// ...
}


read, err := client.RestoreDeviceManagementComanagedDeviceCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SetDeviceManagementComanagedDeviceCloudPCReviewStatu`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.SetDeviceManagementComanagedDeviceCloudPCReviewStatuRequest{
	// ...
}


read, err := client.SetDeviceManagementComanagedDeviceCloudPCReviewStatu(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SetDeviceManagementComanagedDeviceDeviceName`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.SetDeviceManagementComanagedDeviceDeviceNameRequest{
	// ...
}


read, err := client.SetDeviceManagementComanagedDeviceDeviceName(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateComanagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.UpdateDeviceManagementComanagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.UpdateDeviceManagementComanagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateDeviceManagementComanagedDeviceWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
