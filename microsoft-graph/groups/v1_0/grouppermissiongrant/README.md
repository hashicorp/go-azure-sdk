
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouppermissiongrant` Documentation

The `grouppermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouppermissiongrant"
```


### Client Initialization

```go
client := grouppermissiongrant.NewGroupPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPermissionGrantClient.CheckGroupByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := grouppermissiongrant.CheckGroupByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckGroupByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.CheckGroupByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := grouppermissiongrant.CheckGroupByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckGroupByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.CreateGroupByIdPermissionGrant`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

payload := grouppermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateGroupByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.DeleteGroupByIdPermissionGrantById`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteGroupByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantById`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetGroupByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := grouppermissiongrant.GetGroupByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetGroupByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := grouppermissiongrant.GetGroupByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetGroupByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantsAvailableExtensionProperties`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

// alternatively `client.GetGroupByIdPermissionGrantsAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetGroupByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPermissionGrantClient.GetGroupByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

// alternatively `client.GetGroupByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetGroupByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPermissionGrantClient.ListGroupByIdPermissionGrants`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPermissionGrantClient.RestoreGroupByIdPermissionGrantById`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreGroupByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.UpdateGroupByIdPermissionGrantById`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupPermissionGrantID("groupIdValue", "resourceSpecificPermissionGrantIdValue")

payload := grouppermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateGroupByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPermissionGrantClient.ValidateGroupByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := grouppermissiongrant.NewGroupID("groupIdValue")

payload := grouppermissiongrant.ValidateGroupByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateGroupByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
