
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteampermissiongrant` Documentation

The `userjoinedteampermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteampermissiongrant"
```


### Client Initialization

```go
client := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userjoinedteampermissiongrant.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userjoinedteampermissiongrant.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckUserByIdJoinedTeamByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.CreateUserByIdJoinedTeamByIdPermissionGrant`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.DeleteUserByIdJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userjoinedteampermissiongrant.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userjoinedteampermissiongrant.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetUserByIdJoinedTeamByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantsAvailableExtensionProperties`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.GetUserByIdJoinedTeamByIdPermissionGrantsAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdJoinedTeamByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.GetUserByIdJoinedTeamByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.GetUserByIdJoinedTeamByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdJoinedTeamByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.ListUserByIdJoinedTeamByIdPermissionGrants`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListUserByIdJoinedTeamByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.RestoreUserByIdJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreUserByIdJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.UpdateUserByIdJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantID("userIdValue", "teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := userjoinedteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamPermissionGrantClient.ValidateUserByIdJoinedTeamByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := userjoinedteampermissiongrant.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteampermissiongrant.ValidateUserByIdJoinedTeamByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateUserByIdJoinedTeamByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
