
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user` Documentation

The `user` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
```


### Client Initialization

```go
client := user.NewUserClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserClient.AssignUserLicense`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.AssignUserLicenseRequest{
	// ...
}


read, err := client.AssignUserLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckUserMemberGroup`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckUserMemberGroupRequest{
	// ...
}


read, err := client.CheckUserMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckUserMemberObject`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CheckUserMemberObjectRequest{
	// ...
}


read, err := client.CheckUserMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateChangePassword`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateChangePasswordRequest{
	// ...
}


read, err := client.CreateChangePassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateExportPersonalData`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateExportPersonalDataRequest{
	// ...
}


read, err := client.CreateExportPersonalData(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateFindMeetingTime`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateFindMeetingTimeRequest{
	// ...
}


read, err := client.CreateFindMeetingTime(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateReprocessLicenseAssignment`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateReprocessLicenseAssignment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateRetryServiceProvisioning`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateRetryServiceProvisioning(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateRevokeSignInSession`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.CreateRevokeSignInSession(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateSendMail`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateSendMailRequest{
	// ...
}


read, err := client.CreateSendMail(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateTranslateExchangeId`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateTranslateExchangeIdRequest{
	// ...
}


read, err := client.CreateTranslateExchangeId(ctx, id, payload)
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


### Example Usage: `UserClient.CreateWipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.CreateWipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.CreateWipeManagedAppRegistrationsByDeviceTag(ctx, id, payload)
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

read, err := client.DeleteUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.GetUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserMailTip`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserMailTipRequest{
	// ...
}


read, err := client.GetUserMailTip(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserMemberGroup`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserMemberGroupRequest{
	// ...
}


read, err := client.GetUserMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUserMemberObject`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

payload := user.GetUserMemberObjectRequest{
	// ...
}


read, err := client.GetUserMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetUsersAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := user.GetUsersAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetUsersAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetUsersAvailableExtensionPropertiesComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetUsersByIds`

```go
ctx := context.TODO()

payload := user.GetUsersByIdsRequest{
	// ...
}


// alternatively `client.GetUsersByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetUsersByIdsComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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


### Example Usage: `UserClient.RemoveUserAllDevicesFromManagement`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RemoveUserAllDevicesFromManagement(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RestoreUser`

```go
ctx := context.TODO()
id := user.NewUserID("userIdValue")

read, err := client.RestoreUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
