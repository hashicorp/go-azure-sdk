
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userchatpermissiongrant` Documentation

The `userchatpermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userchatpermissiongrant"
```


### Client Initialization

```go
client := userchatpermissiongrant.NewUserChatPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserChatPermissionGrantClient.CheckUserByIdChatByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userchatpermissiongrant.CheckUserByIdChatByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckUserByIdChatByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.CheckUserByIdChatByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userchatpermissiongrant.CheckUserByIdChatByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckUserByIdChatByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.CreateUserByIdChatByIdPermissionGrant`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateUserByIdChatByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.DeleteUserByIdChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteUserByIdChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetUserByIdChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userchatpermissiongrant.GetUserByIdChatByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetUserByIdChatByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userchatpermissiongrant.GetUserByIdChatByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetUserByIdChatByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

read, err := client.GetUserByIdChatByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantsAvailableExtensionProperties`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.GetUserByIdChatByIdPermissionGrantsAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdChatByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatPermissionGrantClient.GetUserByIdChatByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.GetUserByIdChatByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdChatByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatPermissionGrantClient.ListUserByIdChatByIdPermissionGrants`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

// alternatively `client.ListUserByIdChatByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdChatByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserChatPermissionGrantClient.RestoreUserByIdChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreUserByIdChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.UpdateUserByIdChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatPermissionGrantID("userIdValue", "chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userchatpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateUserByIdChatByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserChatPermissionGrantClient.ValidateUserByIdChatByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := userchatpermissiongrant.NewUserChatID("userIdValue", "chatIdValue")

payload := userchatpermissiongrant.ValidateUserByIdChatByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateUserByIdChatByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
