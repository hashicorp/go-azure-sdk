
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


### Example Usage: `ManagedDeviceClient.BypassManagedDeviceActivationLock`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.BypassManagedDeviceActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.ChangeManagedDeviceAssignment`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.ChangeManagedDeviceAssignmentRequest{
	// ...
}


read, err := client.ChangeManagedDeviceAssignment(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CleanManagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanManagedDeviceWindowsDevice(ctx, id, payload)
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


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceBulkSetCloudPCReviewStatus`

```go
ctx := context.TODO()

payload := manageddevice.CreateManagedDeviceBulkSetCloudPCReviewStatusRequest{
	// ...
}


read, err := client.CreateManagedDeviceBulkSetCloudPCReviewStatus(ctx, payload)
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


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceInitiateDeviceAttestation`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.CreateManagedDeviceInitiateDeviceAttestation(ctx, id)
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


### Example Usage: `ManagedDeviceClient.CreateManagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.CreateManagedDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateManagedDeviceLogCollectionRequest(ctx, id, payload)
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


### Example Usage: `ManagedDeviceClient.DeleteManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.DeleteManagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteManagedDeviceUserFromSharedAppleDevice(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.DisableManagedDevice(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.DisableManagedDeviceLostMode(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.EnableManagedDeviceLostModeRequest{
	// ...
}


read, err := client.EnableManagedDeviceLostMode(ctx, id, payload)
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


read, err := client.GetManagedDevicesCount(ctx, manageddevice.DefaultGetManagedDevicesCountOperationOptions())
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


// alternatively `client.ListManagedDevices(ctx, manageddevice.DefaultListManagedDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListManagedDevicesComplete(ctx, manageddevice.DefaultListManagedDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedDeviceClient.LocateManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.LocateManagedDevice(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.LogoutManagedDeviceSharedAppleDeviceActiveUser(ctx, id)
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

payload := manageddevice.MoveManagedDevicesToOURequest{
	// ...
}


read, err := client.MoveManagedDevicesToOU(ctx, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.PauseManagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseManagedDeviceConfigurationRefresh(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RebootManagedDeviceNow(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RecoverManagedDevicePasscode(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RemoteLockManagedDevice(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RemoveManagedDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.ReprovisionManagedDeviceCloudPC(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RequestManagedDeviceRemoteAssistance(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.ResetManagedDevicePasscode(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.ResizeManagedDeviceCloudPCRequest{
	// ...
}


read, err := client.ResizeManagedDeviceCloudPC(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.RestoreManagedDeviceCloudPCRequest{
	// ...
}


read, err := client.RestoreManagedDeviceCloudPC(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RetireManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.RevokeManagedDeviceAppleVppLicense`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.RevokeManagedDeviceAppleVppLicense(ctx, id)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.SendManagedDeviceCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.SendManagedDeviceCustomNotificationToCompanyPortal(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.SetManagedDeviceCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetManagedDeviceCloudPCReviewStatus(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.SetManagedDeviceNameRequest{
	// ...
}


read, err := client.SetManagedDeviceName(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.ShutDownManagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedDeviceClient.SyncManagedDevice`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

read, err := client.SyncManagedDevice(ctx, id)
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


### Example Usage: `ManagedDeviceClient.UpdateManagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.UpdateManagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateManagedDeviceWindowsDeviceAccount(ctx, id, payload)
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
id := manageddevice.NewDeviceManagementManagedDeviceID("managedDeviceIdValue")

payload := manageddevice.WipeManagedDeviceRequest{
	// ...
}


read, err := client.WipeManagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
