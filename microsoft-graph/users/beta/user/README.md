
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/user` Documentation

The `user` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/user"
```


### Client Initialization

```go
client := user.NewUserClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserClient.AssignUserByIdLicense`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.AssignUserByIdLicenseRequest{
	// ...
}


read, err := client.AssignUserByIdLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckUserByIdMemberGroup`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckUserByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckUserByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckUserByIdMemberObject`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckUserByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckUserByIdMemberObject(ctx, id, payload)
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


### Example Usage: `UserClient.CreateUserByIdChangePassword`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdChangePasswordRequest{
	// ...
}


read, err := client.CreateUserByIdChangePassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdExportPersonalData`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdExportPersonalDataRequest{
	// ...
}


read, err := client.CreateUserByIdExportPersonalData(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdFindMeetingTime`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdFindMeetingTimeRequest{
	// ...
}


read, err := client.CreateUserByIdFindMeetingTime(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdInvalidateAllRefreshToken`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdInvalidateAllRefreshToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdReprocessLicenseAssignment`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdReprocessLicenseAssignment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdRetryServiceProvisioning`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdRetryServiceProvisioning(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdRevokeSignInSession`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdRevokeSignInSession(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdSendMail`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdSendMailRequest{
	// ...
}


read, err := client.CreateUserByIdSendMail(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdTranslateExchangeId`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdTranslateExchangeIdRequest{
	// ...
}


read, err := client.CreateUserByIdTranslateExchangeId(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdUnblockManagedApp`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdUnblockManagedApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdWipeAndBlockManagedApp`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateUserByIdWipeAndBlockManagedApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdWipeManagedAppRegistrationByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdWipeManagedAppRegistrationByDeviceTagRequest{
	// ...
}


read, err := client.CreateUserByIdWipeManagedAppRegistrationByDeviceTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdWipeManagedAppRegistrationsByAzureAdDeviceId`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdWipeManagedAppRegistrationsByAzureAdDeviceIdRequest{
	// ...
}


read, err := client.CreateUserByIdWipeManagedAppRegistrationsByAzureAdDeviceId(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUserByIdWipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateUserByIdWipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.CreateUserByIdWipeManagedAppRegistrationsByDeviceTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.DeleteUserById`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.DeleteUserById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserById`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.GetUserById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserByIdMailTip`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserByIdMailTipRequest{
	// ...
}


read, err := client.GetUserByIdMailTip(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserByIdMemberGroup`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserByIdMemberGroupRequest{
	// ...
}


read, err := client.GetUserByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserByIdMemberObject`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserByIdMemberObjectRequest{
	// ...
}


read, err := client.GetUserByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserCount`

```go
ctx := context.TODO()


read, err := client.GetUserCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUsersByIds`

```go
ctx := context.TODO()


// alternatively `client.GetUsersByIds(ctx)` can be used to do batched pagination
items, err := client.GetUsersByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetUsersUserOwnedObject`

```go
ctx := context.TODO()

payload := user.GetUsersUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUsersUserOwnedObject(ctx, payload)
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


// alternatively `client.ListUsers(ctx)` can be used to do batched pagination
items, err := client.ListUsersComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.RemoveUserByIdAllDevicesFromManagement`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RemoveUserByIdAllDevicesFromManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RestoreUserById`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RestoreUserById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.UpdateUserById`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.User{
	// ...
}


read, err := client.UpdateUserById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ValidateUsersPassword`

```go
ctx := context.TODO()

payload := user.ValidateUsersPasswordRequest{
	// ...
}


read, err := client.ValidateUsersPassword(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ValidateUsersProperty`

```go
ctx := context.TODO()

payload := user.ValidateUsersPropertyRequest{
	// ...
}


read, err := client.ValidateUsersProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
