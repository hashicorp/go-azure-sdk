
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


### Example Usage: `GroupClient.AddFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.AddFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.AssignLicense`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGrantedPermissionsForApps`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

// alternatively `client.CheckGrantedPermissionsForApps(ctx, id, group.DefaultCheckGrantedPermissionsForAppsOperationOptions())` can be used to do batched pagination
items, err := client.CheckGrantedPermissionsForAppsComplete(ctx, id, group.DefaultCheckGrantedPermissionsForAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, group.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, group.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, group.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, group.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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

read, err := client.DeleteGroup(ctx, id, group.DefaultDeleteGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := group.GetAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetAvailableExtensionProperties(ctx, payload, group.DefaultGetAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.GetAvailableExtensionPropertiesComplete(ctx, payload, group.DefaultGetAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetByIds`

```go
ctx := context.TODO()

payload := group.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, group.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, group.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, group.DefaultGetCountOperationOptions())
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

read, err := client.GetGroup(ctx, id, group.DefaultGetGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetMemberGroups`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, group.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, group.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetMemberObjects`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

payload := group.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, group.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, group.DefaultGetMemberObjectsOperationOptions())
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


// alternatively `client.ListGroups(ctx, group.DefaultListGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupsComplete(ctx, group.DefaultListGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.RemoveFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RemoveFavorite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.Renew`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.Renew(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ResetUnseenCount`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.ResetUnseenCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.Restore`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.Restore(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RetryServiceProvisioning`

```go
ctx := context.TODO()
id := group.NewGroupID("groupIdValue")

read, err := client.RetryServiceProvisioning(ctx, id)
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


### Example Usage: `GroupClient.ValidateProperty`

```go
ctx := context.TODO()

payload := group.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
