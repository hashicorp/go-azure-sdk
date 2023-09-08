
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memanageddevice` Documentation

The `memanageddevice` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memanageddevice"
```


### Client Initialization

```go
client := memanageddevice.NewMeManagedDeviceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDevice`

```go
ctx := context.TODO()

payload := memanageddevice.ManagedDevice{
	// ...
}


read, err := client.CreateMeManagedDevice(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceBulkReprovisionCloudPC`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceBulkReprovisionCloudPCRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceBulkReprovisionCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceBulkRestoreCloudPC`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceBulkRestoreCloudPCRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceBulkRestoreCloudPC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceBulkSetCloudPCReviewStatu`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceBulkSetCloudPCReviewStatuRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceBulkSetCloudPCReviewStatu(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdActivateDeviceEsim`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdActivateDeviceEsimRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdActivateDeviceEsim(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdBypassActivationLock`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdBypassActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdCleanWindowsDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdCleanWindowsDeviceRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdCleanWindowsDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdCreateDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdCreateDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdCreateDeviceLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdDeleteUserFromSharedAppleDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDeprovision`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdDeprovisionRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdDeprovision(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDisable`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdDisable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdDisableLostMode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdDisableLostMode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdEnableLostMode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdEnableLostModeRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdEnableLostMode(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdEnrollNowAction`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdEnrollNowAction(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdInitiateMobileDeviceManagementKeyRecovery`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdInitiateMobileDeviceManagementKeyRecovery(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdInitiateOnDemandProactiveRemediation`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdInitiateOnDemandProactiveRemediationRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdInitiateOnDemandProactiveRemediation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdLocateDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdLocateDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdLogoutSharedAppleDeviceActiveUser`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdLogoutSharedAppleDeviceActiveUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdOverrideComplianceState`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdOverrideComplianceStateRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdOverrideComplianceState(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdPlayLostModeSound`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdPlayLostModeSoundRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdPlayLostModeSound(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRebootNow`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRebootNow(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRecoverPasscode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRecoverPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdReenable`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdReenable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRemoteLock`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRemoteLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdReprovisionCloudPC`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdReprovisionCloudPC(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRequestRemoteAssistance`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRequestRemoteAssistance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdResetPasscode`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdResetPasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdResizeCloudPC`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdResizeCloudPCRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdResizeCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRetire`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRetire(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRevokeAppleVppLicens`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRevokeAppleVppLicens(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRotateBitLockerKey`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRotateBitLockerKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRotateFileVaultKey`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRotateFileVaultKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdRotateLocalAdminPassword`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdRotateLocalAdminPassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdSendCustomNotificationToCompanyPortal`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdSendCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdSendCustomNotificationToCompanyPortal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdShutDown`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdShutDown(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdSyncDevice`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdSyncDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdTriggerConfigurationManagerAction`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdTriggerConfigurationManagerActionRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdTriggerConfigurationManagerAction(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWindowsDefenderScan`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdWindowsDefenderScanRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdWindowsDefenderScan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWindowsDefenderUpdateSignature`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateMeManagedDeviceByIdWindowsDefenderUpdateSignature(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceByIdWipe`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.CreateMeManagedDeviceByIdWipeRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceByIdWipe(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceDownloadAppDiagnostic`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceDownloadAppDiagnosticRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceDownloadAppDiagnostic(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceExecuteAction`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceExecuteActionRequest{
	// ...
}


read, err := client.CreateMeManagedDeviceExecuteAction(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.CreateMeManagedDeviceMoveDevicesToOU`

```go
ctx := context.TODO()

payload := memanageddevice.CreateMeManagedDeviceMoveDevicesToOURequest{
	// ...
}


read, err := client.CreateMeManagedDeviceMoveDevicesToOU(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.DeleteMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.DeleteMeManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.GetMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.GetMeManagedDeviceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.GetMeManagedDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetMeManagedDeviceCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.ListMeManagedDevices`

```go
ctx := context.TODO()


// alternatively `client.ListMeManagedDevices(ctx)` can be used to do batched pagination
items, err := client.ListMeManagedDevicesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeManagedDeviceClient.RemoveMeManagedDeviceByIdDeviceFirmwareConfigurationInterfaceManagement`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

read, err := client.RemoveMeManagedDeviceByIdDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.RestoreMeManagedDeviceByIdCloudPC`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.RestoreMeManagedDeviceByIdCloudPCRequest{
	// ...
}


read, err := client.RestoreMeManagedDeviceByIdCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.SetMeManagedDeviceByIdCloudPCReviewStatu`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.SetMeManagedDeviceByIdCloudPCReviewStatuRequest{
	// ...
}


read, err := client.SetMeManagedDeviceByIdCloudPCReviewStatu(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.SetMeManagedDeviceByIdDeviceName`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.SetMeManagedDeviceByIdDeviceNameRequest{
	// ...
}


read, err := client.SetMeManagedDeviceByIdDeviceName(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.UpdateMeManagedDeviceById`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.ManagedDevice{
	// ...
}


read, err := client.UpdateMeManagedDeviceById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeManagedDeviceClient.UpdateMeManagedDeviceByIdWindowsDeviceAccount`

```go
ctx := context.TODO()
id := memanageddevice.NewMeManagedDeviceID("managedDeviceIdValue")

payload := memanageddevice.UpdateMeManagedDeviceByIdWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateMeManagedDeviceByIdWindowsDeviceAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
