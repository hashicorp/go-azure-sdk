
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetag` Documentation

The `rolescopetag` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/rolescopetag"
```


### Client Initialization

```go
client := rolescopetag.NewRoleScopeTagClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleScopeTagClient.AssignDeviceManagementRoleScopeTags`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue")

payload := rolescopetag.AssignDeviceManagementRoleScopeTagsRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementRoleScopeTags(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementRoleScopeTagsComplete(ctx, id, payload)
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


read, err := client.CreateRoleScopeTag(ctx, payload)
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
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue")

read, err := client.DeleteRoleScopeTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.GetDeviceManagementRoleScopeTagsRoleScopeTagsByIds`

```go
ctx := context.TODO()

payload := rolescopetag.GetDeviceManagementRoleScopeTagsRoleScopeTagsByIdsRequest{
	// ...
}


// alternatively `client.GetDeviceManagementRoleScopeTagsRoleScopeTagsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetDeviceManagementRoleScopeTagsRoleScopeTagsByIdsComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleScopeTagClient.GetRoleScopeTag`

```go
ctx := context.TODO()
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue")

read, err := client.GetRoleScopeTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleScopeTagClient.GetRoleScopeTagCount`

```go
ctx := context.TODO()


read, err := client.GetRoleScopeTagCount(ctx)
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


// alternatively `client.ListRoleScopeTags(ctx)` can be used to do batched pagination
items, err := client.ListRoleScopeTagsComplete(ctx)
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
id := rolescopetag.NewDeviceManagementRoleScopeTagID("roleScopeTagIdValue")

payload := rolescopetag.RoleScopeTag{
	// ...
}


read, err := client.UpdateRoleScopeTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
