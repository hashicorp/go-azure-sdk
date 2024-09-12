
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/beta/directoryroletemplate` Documentation

The `directoryroletemplate` SDK allows for interaction with the Azure Resource Manager Service `directoryroletemplates` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/beta/directoryroletemplate"
```


### Client Initialization

```go
client := directoryroletemplate.NewDirectoryRoleTemplateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRoleTemplateClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryroletemplate.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryroletemplate.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryroletemplate.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryroletemplate.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.CreateDirectoryRoleTemplate`

```go
ctx := context.TODO()

payload := directoryroletemplate.DirectoryRoleTemplate{
	// ...
}


read, err := client.CreateDirectoryRoleTemplate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.DeleteDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

read, err := client.DeleteDirectoryRoleTemplate(ctx, id, directoryroletemplate.DefaultDeleteDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetByIds`

```go
ctx := context.TODO()

payload := directoryroletemplate.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, directoryroletemplate.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, directoryroletemplate.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, directoryroletemplate.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

read, err := client.GetDirectoryRoleTemplate(ctx, id, directoryroletemplate.DefaultGetDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryroletemplate.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryroletemplate.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryroletemplate.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryroletemplate.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetUserOwnedObject`

```go
ctx := context.TODO()

payload := directoryroletemplate.GetUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.ListDirectoryRoleTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryRoleTemplates(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryRoleTemplatesComplete(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.Restore`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.UpdateDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateIdValue")

payload := directoryroletemplate.DirectoryRoleTemplate{
	// ...
}


read, err := client.UpdateDirectoryRoleTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.ValidateProperty`

```go
ctx := context.TODO()

payload := directoryroletemplate.ValidatePropertyRequest{
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
