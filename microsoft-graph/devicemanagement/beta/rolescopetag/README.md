
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetag` Documentation

The `rolescopetag` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetag"
```


### Client Initialization

```go
client := rolescopetag.NewRoleScopeTagClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleScopeTagClient.AssignRoleScopeTags`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagId")

payload := rolescopetag.AssignRoleScopeTagsRequest{
	// ...
}


// alternatively `client.AssignRoleScopeTags(ctx, id, payload, rolescopetag.DefaultAssignRoleScopeTagsOperationOptions())` can be used to do batched pagination
items, err := client.AssignRoleScopeTagsComplete(ctx, id, payload, rolescopetag.DefaultAssignRoleScopeTagsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleScopeTagClient.CreateRoleScopeTag`

```go
ctx := context.TODO()

payload := rolescopetag.RoleScopeTag{
	// ...
}


read, err := client.CreateRoleScopeTag(ctx, payload, rolescopetag.DefaultCreateRoleScopeTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.DeleteRoleScopeTag`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagId")

read, err := client.DeleteRoleScopeTag(ctx, id, rolescopetag.DefaultDeleteRoleScopeTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.GetRoleScopeTag`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagId")

read, err := client.GetRoleScopeTag(ctx, id, rolescopetag.DefaultGetRoleScopeTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.GetRoleScopeTagsByIds`

```go
ctx := context.TODO()

payload := rolescopetag.GetRoleScopeTagsByIdsRequest{
	// ...
}


// alternatively `client.GetRoleScopeTagsByIds(ctx, payload, rolescopetag.DefaultGetRoleScopeTagsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetRoleScopeTagsByIdsComplete(ctx, payload, rolescopetag.DefaultGetRoleScopeTagsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleScopeTagClient.GetRoleScopeTagsCount`

```go
ctx := context.TODO()


read, err := client.GetRoleScopeTagsCount(ctx, rolescopetag.DefaultGetRoleScopeTagsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.ListRoleScopeTags`

```go
ctx := context.TODO()


// alternatively `client.ListRoleScopeTags(ctx, rolescopetag.DefaultListRoleScopeTagsOperationOptions())` can be used to do batched pagination
items, err := client.ListRoleScopeTagsComplete(ctx, rolescopetag.DefaultListRoleScopeTagsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleScopeTagClient.UpdateRoleScopeTag`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagId")

payload := rolescopetag.RoleScopeTag{
	// ...
}


read, err := client.UpdateRoleScopeTag(ctx, id, payload, rolescopetag.DefaultUpdateRoleScopeTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
