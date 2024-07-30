
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/me` Documentation

The `me` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/me"
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


### Example Usage: `MeClient.CreateChangePassword`

```go
ctx := context.TODO()

payload := me.CreateChangePasswordRequest{
	// ...
}


read, err := client.CreateChangePassword(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateDeletePasswordSingleSignOnCredential`

```go
ctx := context.TODO()

payload := me.CreateDeletePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.CreateDeletePasswordSingleSignOnCredential(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateExportPersonalData`

```go
ctx := context.TODO()

payload := me.CreateExportPersonalDataRequest{
	// ...
}


read, err := client.CreateExportPersonalData(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateFindMeetingTime`

```go
ctx := context.TODO()

payload := me.CreateFindMeetingTimeRequest{
	// ...
}


read, err := client.CreateFindMeetingTime(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateInvalidateAllRefreshToken`

```go
ctx := context.TODO()


read, err := client.CreateInvalidateAllRefreshToken(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateReprocessLicenseAssignment`

```go
ctx := context.TODO()


read, err := client.CreateReprocessLicenseAssignment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateRetryServiceProvisioning`

```go
ctx := context.TODO()


read, err := client.CreateRetryServiceProvisioning(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateRevokeSignInSession`

```go
ctx := context.TODO()


read, err := client.CreateRevokeSignInSession(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateSendMail`

```go
ctx := context.TODO()

payload := me.CreateSendMailRequest{
	// ...
}


read, err := client.CreateSendMail(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateTranslateExchangeId`

```go
ctx := context.TODO()

payload := me.CreateTranslateExchangeIdRequest{
	// ...
}


read, err := client.CreateTranslateExchangeId(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateUnblockManagedApp`

```go
ctx := context.TODO()


read, err := client.CreateUnblockManagedApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateWipeAndBlockManagedApp`

```go
ctx := context.TODO()


read, err := client.CreateWipeAndBlockManagedApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateWipeManagedAppRegistrationByDeviceTag`

```go
ctx := context.TODO()

payload := me.CreateWipeManagedAppRegistrationByDeviceTagRequest{
	// ...
}


read, err := client.CreateWipeManagedAppRegistrationByDeviceTag(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateWipeManagedAppRegistrationsByAzureAdDeviceId`

```go
ctx := context.TODO()

payload := me.CreateWipeManagedAppRegistrationsByAzureAdDeviceIdRequest{
	// ...
}


read, err := client.CreateWipeManagedAppRegistrationsByAzureAdDeviceId(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CreateWipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()

payload := me.CreateWipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.CreateWipeManagedAppRegistrationsByDeviceTag(ctx, payload)
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


### Example Usage: `MeClient.GetMePasswordSingleSignOnCredential`

```go
ctx := context.TODO()


read, err := client.GetMePasswordSingleSignOnCredential(ctx)
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
