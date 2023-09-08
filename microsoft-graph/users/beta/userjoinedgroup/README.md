
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userjoinedgroup` Documentation

The `userjoinedgroup` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userjoinedgroup"
```


### Client Initialization

```go
client := userjoinedgroup.NewUserJoinedGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedGroupClient.CreateUserByIdJoinedGroupEvaluateDynamicMembership`

```go
ctx := context.TODO()
id := userjoinedgroup.NewUserID("userIdValue")

payload := userjoinedgroup.CreateUserByIdJoinedGroupEvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.CreateUserByIdJoinedGroupEvaluateDynamicMembership(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedGroupClient.GetUserByIdJoinedGroupsByIds`

```go
ctx := context.TODO()
id := userjoinedgroup.NewUserID("userIdValue")

// alternatively `client.GetUserByIdJoinedGroupsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetUserByIdJoinedGroupsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedGroupClient.GetUserByIdJoinedGroupsUserOwnedObject`

```go
ctx := context.TODO()
id := userjoinedgroup.NewUserID("userIdValue")

payload := userjoinedgroup.GetUserByIdJoinedGroupsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserByIdJoinedGroupsUserOwnedObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedGroupClient.ListUserByIdJoinedGroups`

```go
ctx := context.TODO()
id := userjoinedgroup.NewUserID("userIdValue")

// alternatively `client.ListUserByIdJoinedGroups(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedGroupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedGroupClient.ValidateUserByIdJoinedGroupsProperty`

```go
ctx := context.TODO()
id := userjoinedgroup.NewUserID("userIdValue")

payload := userjoinedgroup.ValidateUserByIdJoinedGroupsPropertyRequest{
	// ...
}


read, err := client.ValidateUserByIdJoinedGroupsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
