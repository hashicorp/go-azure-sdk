
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteampermissiongrant` Documentation

The `groupteampermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteampermissiongrant"
```


### Client Initialization

```go
client := groupteampermissiongrant.NewGroupTeamPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamPermissionGrantClient.CheckGroupByIdTeamPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := groupteampermissiongrant.CheckGroupByIdTeamPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckGroupByIdTeamPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.CheckGroupByIdTeamPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := groupteampermissiongrant.CheckGroupByIdTeamPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckGroupByIdTeamPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.CreateGroupByIdTeamPermissionGrant`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

payload := groupteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateGroupByIdTeamPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.DeleteGroupByIdTeamPermissionGrantById`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteGroupByIdTeamPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantById`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetGroupByIdTeamPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := groupteampermissiongrant.GetGroupByIdTeamPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetGroupByIdTeamPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := groupteampermissiongrant.GetGroupByIdTeamPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetGroupByIdTeamPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantCount`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantsByIds`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

// alternatively `client.GetGroupByIdTeamPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetGroupByIdTeamPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamPermissionGrantClient.GetGroupByIdTeamPermissionGrantsUserOwnedObject`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

payload := groupteampermissiongrant.GetGroupByIdTeamPermissionGrantsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetGroupByIdTeamPermissionGrantsUserOwnedObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.ListGroupByIdTeamPermissionGrants`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamPermissionGrantClient.RestoreGroupByIdTeamPermissionGrantById`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreGroupByIdTeamPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.UpdateGroupByIdTeamPermissionGrantById`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupTeamPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := groupteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateGroupByIdTeamPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPermissionGrantClient.ValidateGroupByIdTeamPermissionGrantsProperty`

```go
ctx := context.TODO()
id := groupteampermissiongrant.NewGroupID("groupIdValue")

payload := groupteampermissiongrant.ValidateGroupByIdTeamPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateGroupByIdTeamPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
