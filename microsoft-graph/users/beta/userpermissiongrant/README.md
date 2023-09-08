
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpermissiongrant` Documentation

The `userpermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpermissiongrant"
```


### Client Initialization

```go
client := userpermissiongrant.NewUserPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPermissionGrantClient.CheckUserByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userpermissiongrant.CheckUserByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckUserByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.CheckUserByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userpermissiongrant.CheckUserByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckUserByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.CreateUserByIdPermissionGrant`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

payload := userpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateUserByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.DeleteUserByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteUserByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetUserByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userpermissiongrant.GetUserByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetUserByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userpermissiongrant.GetUserByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetUserByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

read, err := client.GetUserByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

// alternatively `client.GetUserByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPermissionGrantClient.GetUserByIdPermissionGrantsUserOwnedObject`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

payload := userpermissiongrant.GetUserByIdPermissionGrantsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserByIdPermissionGrantsUserOwnedObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.ListUserByIdPermissionGrants`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

// alternatively `client.ListUserByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserPermissionGrantClient.RestoreUserByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreUserByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.UpdateUserByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserPermissionGrantID("userIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateUserByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPermissionGrantClient.ValidateUserByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := userpermissiongrant.NewUserID("userIdValue")

payload := userpermissiongrant.ValidateUserByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateUserByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
