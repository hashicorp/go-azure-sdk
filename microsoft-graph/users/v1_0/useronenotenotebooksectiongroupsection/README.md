
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotenotebooksectiongroupsection` Documentation

The `useronenotenotebooksectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSection`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupID("userIdValue", "notebookIdValue", "sectionGroupIdValue")

payload := useronenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksectiongroupsection.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksectiongroupsection.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.DeleteUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupID("userIdValue", "notebookIdValue", "sectionGroupIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSections`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupID("userIdValue", "notebookIdValue", "sectionGroupIdValue")

// alternatively `client.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionClient.UpdateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
