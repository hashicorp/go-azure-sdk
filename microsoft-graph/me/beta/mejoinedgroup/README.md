
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mejoinedgroup` Documentation

The `mejoinedgroup` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mejoinedgroup"
```


### Client Initialization

```go
client := mejoinedgroup.NewMeJoinedGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedGroupClient.CreateMeJoinedGroupEvaluateDynamicMembership`

```go
ctx := context.TODO()

payload := mejoinedgroup.CreateMeJoinedGroupEvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.CreateMeJoinedGroupEvaluateDynamicMembership(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedGroupClient.GetMeJoinedGroupsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetMeJoinedGroupsByIds(ctx)` can be used to do batched pagination
items, err := client.GetMeJoinedGroupsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedGroupClient.GetMeJoinedGroupsUserOwnedObject`

```go
ctx := context.TODO()

payload := mejoinedgroup.GetMeJoinedGroupsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetMeJoinedGroupsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedGroupClient.ListMeJoinedGroups`

```go
ctx := context.TODO()


// alternatively `client.ListMeJoinedGroups(ctx)` can be used to do batched pagination
items, err := client.ListMeJoinedGroupsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedGroupClient.ValidateMeJoinedGroupsProperty`

```go
ctx := context.TODO()

payload := mejoinedgroup.ValidateMeJoinedGroupsPropertyRequest{
	// ...
}


read, err := client.ValidateMeJoinedGroupsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
