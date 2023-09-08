
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/group` Documentation

The `group` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/group"
```


### Client Initialization

```go
client := group.NewGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupClient.AddGroupByIdFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.AddGroupByIdFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.AssignGroupByIdLicense`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.AssignGroupByIdLicenseRequest{
	// ...
}


read, err := client.AssignGroupByIdLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGroupByIdGrantedPermissionsForApps`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

// alternatively `client.CheckGroupByIdGrantedPermissionsForApps(ctx, id)` can be used to do batched pagination
items, err := client.CheckGroupByIdGrantedPermissionsForAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckGroupByIdMemberGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckGroupByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckGroupByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGroupByIdMemberObject`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckGroupByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckGroupByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroup`

```go
ctx := context.TODO()

payload := group.Group{
	// ...
}


read, err := client.CreateGroup(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupByIdEvaluateDynamicMembership`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CreateGroupByIdEvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.CreateGroupByIdEvaluateDynamicMembership(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupByIdResetUnseenCount`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdResetUnseenCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupByIdRetryServiceProvisioning`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdRetryServiceProvisioning(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupByIdSubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdSubscribeByMail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupByIdUnsubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdUnsubscribeByMail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroupEvaluateDynamicMembership`

```go
ctx := context.TODO()

payload := group.CreateGroupEvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.CreateGroupEvaluateDynamicMembership(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.DeleteGroupById`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.DeleteGroupById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupById`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.GetGroupById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupByIdMemberGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetGroupByIdMemberGroupRequest{
	// ...
}


read, err := client.GetGroupByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupByIdMemberObject`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetGroupByIdMemberObjectRequest{
	// ...
}


read, err := client.GetGroupByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupCount`

```go
ctx := context.TODO()


read, err := client.GetGroupCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetGroupsByIds(ctx)` can be used to do batched pagination
items, err := client.GetGroupsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetGroupsUserOwnedObject`

```go
ctx := context.TODO()

payload := group.GetGroupsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetGroupsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ListGroups`

```go
ctx := context.TODO()


// alternatively `client.ListGroups(ctx)` can be used to do batched pagination
items, err := client.ListGroupsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.RemoveGroupByIdFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RemoveGroupByIdFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RenewGroupById`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RenewGroupById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RestoreGroupById`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RestoreGroupById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.UpdateGroupById`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.Group{
	// ...
}


read, err := client.UpdateGroupById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ValidateGroupByIdProperty`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.ValidateGroupByIdPropertyRequest{
	// ...
}


read, err := client.ValidateGroupByIdProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ValidateGroupsProperty`

```go
ctx := context.TODO()

payload := group.ValidateGroupsPropertyRequest{
	// ...
}


read, err := client.ValidateGroupsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
