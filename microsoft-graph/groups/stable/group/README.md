
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/group` Documentation

The `group` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/group"
```


### Client Initialization

```go
client := group.NewGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupClient.AddGroupFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.AddGroupFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.AssignGroupLicense`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.AssignGroupLicenseRequest{
	// ...
}


read, err := client.AssignGroupLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGroupGrantedPermissionsForApps`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

// alternatively `client.CheckGroupGrantedPermissionsForApps(ctx, id)` can be used to do batched pagination
items, err := client.CheckGroupGrantedPermissionsForAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckGroupMemberGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckGroupMemberGroupRequest{
	// ...
}


read, err := client.CheckGroupMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGroupMemberObject`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckGroupMemberObjectRequest{
	// ...
}


read, err := client.CheckGroupMemberObject(ctx, id, payload)
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


### Example Usage: `GroupClient.CreateResetUnseenCount`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateResetUnseenCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateRetryServiceProvisioning`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateRetryServiceProvisioning(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateSubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateSubscribeByMail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateUnsubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.CreateUnsubscribeByMail(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.DeleteGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.DeleteGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetCount`

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


### Example Usage: `GroupClient.GetGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.GetGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupMemberGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetGroupMemberGroupRequest{
	// ...
}


read, err := client.GetGroupMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupMemberObject`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetGroupMemberObjectRequest{
	// ...
}


read, err := client.GetGroupMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroupsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := group.GetGroupsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetGroupsAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetGroupsAvailableExtensionPropertiesComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetGroupsByIds`

```go
ctx := context.TODO()

payload := group.GetGroupsByIdsRequest{
	// ...
}


// alternatively `client.GetGroupsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetGroupsByIdsComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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


### Example Usage: `GroupClient.RemoveGroupFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RemoveGroupFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RenewGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RenewGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RestoreGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RestoreGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.UpdateGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.Group{
	// ...
}


read, err := client.UpdateGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ValidateGroupProperty`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.ValidateGroupPropertyRequest{
	// ...
}


read, err := client.ValidateGroupProperty(ctx, id, payload)
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
