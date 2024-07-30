
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevice` Documentation

The `manageddevice` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/manageddevice"
```


### Client Initialization

```go
client := manageddevice.NewManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedDeviceClient.CreateManagedDevice`

```go
ctx := context.TODO()

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateManagedDevice(ctx, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateManagedDeviceActivateDeviceEsim(ctx, id, payload)
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

payload := manageddevice.CreateManagedDeviceBulkReprovisionCloudPCRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkReprovisionCloudPC(ctx, payload)
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

payload := manageddevice.CreateManagedDeviceBulkRestoreCloudPCRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkRestoreCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBulkSetCloudPCReviewStatu`

```go
ctx := context.TODO()

payload := manageddevice.CreateManagedDeviceBulkSetCloudPCReviewStatuRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkSetCloudPCReviewStatu(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBypassActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceCleanWindowsDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateManagedDeviceCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceCreateDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceCreateDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateManagedDeviceCreateDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateManagedDeviceDeleteUserFromSharedAppleDevice(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceDeprovisionRequest{
	// ...
}


read, err := client.CreateManagedDeviceDeprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDisable`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceDisable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceDisableLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceDisableLostMode(ctx, id)
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

payload := manageddevice.CreateManagedDeviceDownloadAppDiagnosticRequest{
	// ...
}


read, err := client.CreateManagedDeviceDownloadAppDiagnostic(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceEnableLostMode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceEnableLostModeRequest{
	// ...
}


read, err := client.CreateManagedDeviceEnableLostMode(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceEnrollNowAction(ctx, id)
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

payload := manageddevice.CreateManagedDeviceExecuteActionRequest{
	// ...
}


read, err := client.CreateManagedDeviceExecuteAction(ctx, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceInitiateMobileDeviceManagementKeyRecovery(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateManagedDeviceInitiateOnDemandProactiveRemediation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLocateDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceMoveDevicesToOU`

```go
ctx := context.TODO()

payload := manageddevice.CreateManagedDeviceMoveDevicesToOURequest{
	// ...
}


read, err := client.CreateManagedDeviceMoveDevicesToOU(ctx, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateManagedDeviceOverrideComplianceState(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDevicePlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateManagedDevicePlayLostModeSound(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRebootNow`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRecoverPasscode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRecoverPasscode(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceReenable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRemoteLock`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceReprovisionCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceReprovisionCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRequestRemoteAssistance`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceResetPasscode`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceResizeCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceResizeCloudPCRequest{
	// ...
}


read, err := client.CreateManagedDeviceResizeCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRetire`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceRevokeAppleVppLicens`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRevokeAppleVppLicens(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRotateBitLockerKey(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRotateFileVaultKey(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceRotateLocalAdminPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceSendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceSendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.CreateManagedDeviceSendCustomNotificationToCompanyPortal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceShutDown`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceSyncDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceSyncDevice(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateManagedDeviceTriggerConfigurationManagerAction(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateManagedDeviceWindowsDefenderScan(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceWipe`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceWipeRequest{
	// ...
}


read, err := client.CreateManagedDeviceWipe(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteManagedDevice(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.GetManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.GetManagedDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetManagedDeviceCount(ctx)
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


// alternatively `client.ListManagedDevices(ctx)` can be used to do batched pagination
items, err := client.ListManagedDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceClient.PauseDeviceManagementManagedDeviceConfigurationRefresh`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.PauseDeviceManagementManagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseDeviceManagementManagedDeviceConfigurationRefresh(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RemoveDeviceManagementManagedDeviceDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RemoveDeviceManagementManagedDeviceDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RestoreDeviceManagementManagedDeviceCloudPC`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.RestoreDeviceManagementManagedDeviceCloudPCRequest{
	// ...
}


read, err := client.RestoreDeviceManagementManagedDeviceCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SetDeviceManagementManagedDeviceCloudPCReviewStatu`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.SetDeviceManagementManagedDeviceCloudPCReviewStatuRequest{
	// ...
}


read, err := client.SetDeviceManagementManagedDeviceCloudPCReviewStatu(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SetDeviceManagementManagedDeviceDeviceName`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.SetDeviceManagementManagedDeviceDeviceNameRequest{
	// ...
}


read, err := client.SetDeviceManagementManagedDeviceDeviceName(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.UpdateDeviceManagementManagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.UpdateDeviceManagementManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateDeviceManagementManagedDeviceWindowsDeviceAccount(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
