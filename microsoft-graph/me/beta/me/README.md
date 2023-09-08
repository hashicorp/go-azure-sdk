
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/me` Documentation

The `me` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/me"
```


### Client Initialization

```go
client := me.NewMeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeClient.AssignMeLicense`

```go
ctx := context.TODO()

payload := me.AssignMeLicenseRequest{
	// ...
}


read, err := client.AssignMeLicense(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CheckMeMemberGroup`

```go
ctx := context.TODO()

payload := me.CheckMeMemberGroupRequest{
	// ...
}


read, err := client.CheckMeMemberGroup(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CheckMeMemberObject`

```go
ctx := context.TODO()

payload := me.CheckMeMemberObjectRequest{
	// ...
}


read, err := client.CheckMeMemberObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeChangePassword`

```go
ctx := context.TODO()

payload := me.CreateMeChangePasswordRequest{
	// ...
}


read, err := client.CreateMeChangePassword(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeExportPersonalData`

```go
ctx := context.TODO()

payload := me.CreateMeExportPersonalDataRequest{
	// ...
}


read, err := client.CreateMeExportPersonalData(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeFindMeetingTime`

```go
ctx := context.TODO()

payload := me.CreateMeFindMeetingTimeRequest{
	// ...
}


read, err := client.CreateMeFindMeetingTime(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeInvalidateAllRefreshToken`

```go
ctx := context.TODO()


read, err := client.CreateMeInvalidateAllRefreshToken(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeReprocessLicenseAssignment`

```go
ctx := context.TODO()


read, err := client.CreateMeReprocessLicenseAssignment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeRetryServiceProvisioning`

```go
ctx := context.TODO()


read, err := client.CreateMeRetryServiceProvisioning(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeRevokeSignInSession`

```go
ctx := context.TODO()


read, err := client.CreateMeRevokeSignInSession(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeSendMail`

```go
ctx := context.TODO()

payload := me.CreateMeSendMailRequest{
	// ...
}


read, err := client.CreateMeSendMail(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeTranslateExchangeId`

```go
ctx := context.TODO()

payload := me.CreateMeTranslateExchangeIdRequest{
	// ...
}


read, err := client.CreateMeTranslateExchangeId(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeUnblockManagedApp`

```go
ctx := context.TODO()


read, err := client.CreateMeUnblockManagedApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeWipeAndBlockManagedApp`

```go
ctx := context.TODO()


read, err := client.CreateMeWipeAndBlockManagedApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeWipeManagedAppRegistrationByDeviceTag`

```go
ctx := context.TODO()

payload := me.CreateMeWipeManagedAppRegistrationByDeviceTagRequest{
	// ...
}


read, err := client.CreateMeWipeManagedAppRegistrationByDeviceTag(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeWipeManagedAppRegistrationsByAzureAdDeviceId`

```go
ctx := context.TODO()

payload := me.CreateMeWipeManagedAppRegistrationsByAzureAdDeviceIdRequest{
	// ...
}


read, err := client.CreateMeWipeManagedAppRegistrationsByAzureAdDeviceId(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateMeWipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()

payload := me.CreateMeWipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.CreateMeWipeManagedAppRegistrationsByDeviceTag(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMe`

```go
ctx := context.TODO()


read, err := client.GetMe(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMeMailTip`

```go
ctx := context.TODO()

payload := me.GetMeMailTipRequest{
	// ...
}


read, err := client.GetMeMailTip(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMeMemberGroup`

```go
ctx := context.TODO()

payload := me.GetMeMemberGroupRequest{
	// ...
}


read, err := client.GetMeMemberGroup(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMeMemberObject`

```go
ctx := context.TODO()

payload := me.GetMeMemberObjectRequest{
	// ...
}


read, err := client.GetMeMemberObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RemoveMeAllDevicesFromManagement`

```go
ctx := context.TODO()


read, err := client.RemoveMeAllDevicesFromManagement(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RestoreMe`

```go
ctx := context.TODO()


read, err := client.RestoreMe(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.UpdateMe`

```go
ctx := context.TODO()

payload := me.User{
	// ...
}


read, err := client.UpdateMe(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
