
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user` Documentation

The `user` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user"
```


### Client Initialization

```go
client := user.NewUserClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserClient.AssignLicense`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ChangePassword`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.ChangePasswordRequest{
	// ...
}


read, err := client.ChangePassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, user.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, user.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, user.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, user.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.CreateConvertExternalToInternalMemberUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateConvertExternalToInternalMemberUserRequest{
	// ...
}


read, err := client.CreateConvertExternalToInternalMemberUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateInvalidateAllRefreshToken`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateInvalidateAllRefreshToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUnblockManagedApp`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUnblockManagedApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUser`

```go
ctx := context.TODO()

payload := user.User{
	// ...
}


read, err := client.CreateUser(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.DeletePasswordSingleSignOnCredential`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.DeletePasswordSingleSignOnCredentialRequest{
	// ...
}


read, err := client.DeletePasswordSingleSignOnCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.DeleteUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.DeleteUser(ctx, id, user.DefaultDeleteUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ExportPersonalData`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.ExportPersonalDataRequest{
	// ...
}


read, err := client.ExportPersonalData(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.FindMeetingTime`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.FindMeetingTimeRequest{
	// ...
}


read, err := client.FindMeetingTime(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetByIds`

```go
ctx := context.TODO()

payload := user.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, user.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, user.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, user.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetMailTips`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetMailTipsRequest{
	// ...
}


// alternatively `client.GetMailTips(ctx, id, payload, user.DefaultGetMailTipsOperationOptions())` can be used to do batched pagination
items, err := client.GetMailTipsComplete(ctx, id, payload, user.DefaultGetMailTipsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetMemberGroups`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, user.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, user.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetMemberObjects`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, user.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, user.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetOwnedObject`

```go
ctx := context.TODO()

payload := user.GetOwnedObjectRequest{
	// ...
}


read, err := client.GetOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetPasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

// alternatively `client.GetPasswordSingleSignOnCredentials(ctx, id, user.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())` can be used to do batched pagination
items, err := client.GetPasswordSingleSignOnCredentialsComplete(ctx, id, user.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.GetUser(ctx, id, user.DefaultGetUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ListUsers`

```go
ctx := context.TODO()


// alternatively `client.ListUsers(ctx, user.DefaultListUsersOperationOptions())` can be used to do batched pagination
items, err := client.ListUsersComplete(ctx, user.DefaultListUsersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.RemoveAllDevicesFromManagement`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RemoveAllDevicesFromManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ReprocessLicenseAssignment`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.ReprocessLicenseAssignment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.Restore`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RetryServiceProvisioning`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RetryServiceProvisioning(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RevokeSignInSession`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RevokeSignInSession(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.SendMail`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.SendMailRequest{
	// ...
}


read, err := client.SendMail(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.TranslateExchangeIds`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.TranslateExchangeIdsRequest{
	// ...
}


// alternatively `client.TranslateExchangeIds(ctx, id, payload, user.DefaultTranslateExchangeIdsOperationOptions())` can be used to do batched pagination
items, err := client.TranslateExchangeIdsComplete(ctx, id, payload, user.DefaultTranslateExchangeIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.UpdateUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.User{
	// ...
}


read, err := client.UpdateUser(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ValidatePassword`

```go
ctx := context.TODO()

payload := user.ValidatePasswordRequest{
	// ...
}


read, err := client.ValidatePassword(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ValidateProperty`

```go
ctx := context.TODO()

payload := user.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeAndBlockManagedApp`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.WipeAndBlockManagedApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.WipeManagedAppRegistrationByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationByDeviceTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationsByAzureAdDeviceId`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.WipeManagedAppRegistrationsByAzureAdDeviceIdRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByAzureAdDeviceId(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.WipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByDeviceTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
