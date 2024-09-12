
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole` Documentation

The `directoryrole` SDK allows for interaction with the Azure Resource Manager Service `directoryroles` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
```


### Client Initialization

```go
client := directoryrole.NewDirectoryRoleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRoleClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

payload := directoryrole.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryrole.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryrole.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

payload := directoryrole.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryrole.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryrole.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.CreateDirectoryRole`

```go
ctx := context.TODO()

payload := directoryrole.DirectoryRole{
	// ...
}


read, err := client.CreateDirectoryRole(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.DeleteDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

read, err := client.DeleteDirectoryRole(ctx, id, directoryrole.DefaultDeleteDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.GetAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := directoryrole.GetAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetAvailableExtensionProperties(ctx, payload, directoryrole.DefaultGetAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.GetAvailableExtensionPropertiesComplete(ctx, payload, directoryrole.DefaultGetAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.GetByIds`

```go
ctx := context.TODO()

payload := directoryrole.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, directoryrole.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, directoryrole.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, directoryrole.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.GetDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

read, err := client.GetDirectoryRole(ctx, id, directoryrole.DefaultGetDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

payload := directoryrole.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryrole.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryrole.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

payload := directoryrole.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryrole.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryrole.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.ListDirectoryRoles`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryRoles(ctx, directoryrole.DefaultListDirectoryRolesOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryRolesComplete(ctx, directoryrole.DefaultListDirectoryRolesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.Restore`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

read, err := client.Restore(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.UpdateDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleIdValue")

payload := directoryrole.DirectoryRole{
	// ...
}


read, err := client.UpdateDirectoryRole(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.ValidateProperty`

```go
ctx := context.TODO()

payload := directoryrole.ValidatePropertyRequest{
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
