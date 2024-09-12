
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


### Example Usage: `ComanagedDeviceClient.BypassComanagedDeviceActivationLock`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.BypassComanagedDeviceActivationLock(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ChangeComanagedDeviceAssignment`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.ChangeComanagedDeviceAssignmentRequest{
	// ...
}


read, err := client.ChangeComanagedDeviceAssignment(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CleanComanagedDeviceWindowsDeviceRequest{
	// ...
}


read, err := client.CleanComanagedDeviceWindowsDevice(ctx, id, payload)
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


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceBulkSetCloudPCReviewStatus`

```go
ctx := context.TODO()

payload := comanageddevice.CreateComanagedDeviceBulkSetCloudPCReviewStatusRequest{
	// ...
}


read, err := client.CreateComanagedDeviceBulkSetCloudPCReviewStatus(ctx, payload)
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


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceInitiateDeviceAttestation`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.CreateComanagedDeviceInitiateDeviceAttestation(ctx, id)
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


### Example Usage: `ComanagedDeviceClient.CreateComanagedDeviceLogCollectionRequest`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.CreateComanagedDeviceLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateComanagedDeviceLogCollectionRequest(ctx, id, payload)
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


### Example Usage: `ComanagedDeviceClient.DeleteComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.DeleteComanagedDeviceUserFromSharedAppleDeviceRequest{
	// ...
}


read, err := client.DeleteComanagedDeviceUserFromSharedAppleDevice(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.DisableComanagedDevice(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.DisableComanagedDeviceLostMode(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.EnableComanagedDeviceLostModeRequest{
	// ...
}


read, err := client.EnableComanagedDeviceLostMode(ctx, id, payload)
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


### Example Usage: `ComanagedDeviceClient.LocateComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.LocateComanagedDevice(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.LogoutComanagedDeviceSharedAppleDeviceActiveUser(ctx, id)
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


read, err := client.MoveComanagedDevicesToOU(ctx, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.PauseComanagedDeviceConfigurationRefreshRequest{
	// ...
}


read, err := client.PauseComanagedDeviceConfigurationRefresh(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RebootComanagedDeviceNow(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RecoverComanagedDevicePasscode(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RemoteLockComanagedDevice(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RemoveComanagedDeviceFirmwareConfigurationInterfaceManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ReprovisionComanagedDeviceCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.ReprovisionComanagedDeviceCloudPC(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RequestComanagedDeviceRemoteAssistance(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.ResetComanagedDevicePasscode(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.ResizeComanagedDeviceCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.ResizeComanagedDeviceCloudPCRequest{
	// ...
}


read, err := client.ResizeComanagedDeviceCloudPC(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RestoreComanagedDeviceCloudPC`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.RestoreComanagedDeviceCloudPCRequest{
	// ...
}


read, err := client.RestoreComanagedDeviceCloudPC(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RetireComanagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.RevokeComanagedDeviceAppleVppLicense`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.RevokeComanagedDeviceAppleVppLicense(ctx, id)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.SendComanagedDeviceCustomNotificationToCompanyPortalRequest{
	// ...
}


read, err := client.SendComanagedDeviceCustomNotificationToCompanyPortal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SetComanagedDeviceCloudPCReviewStatus`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.SetComanagedDeviceCloudPCReviewStatusRequest{
	// ...
}


read, err := client.SetComanagedDeviceCloudPCReviewStatus(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.SetComanagedDeviceNameRequest{
	// ...
}


read, err := client.SetComanagedDeviceName(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.ShutDownComanagedDevice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComanagedDeviceClient.SyncComanagedDevice`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

read, err := client.SyncComanagedDevice(ctx, id)
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


### Example Usage: `ComanagedDeviceClient.UpdateComanagedDeviceWindowsDeviceAccount`

```go
ctx := context.TODO()
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.UpdateComanagedDeviceWindowsDeviceAccountRequest{
	// ...
}


read, err := client.UpdateComanagedDeviceWindowsDeviceAccount(ctx, id, payload)
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
id := comanageddevice.NewDeviceManagementComanagedDeviceID("managedDeviceIdValue")

payload := comanageddevice.WipeComanagedDeviceRequest{
	// ...
}


read, err := client.WipeComanagedDevice(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
