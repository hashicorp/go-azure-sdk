
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotenotebooksectiongroupsectionpage` Documentation

The `useronenotenotebooksectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectiongroupsectionpage.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectiongroupsectionpage.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.DeleteUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteNotebookSectionGroupSectionPageClient.UpdateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageID("userIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateUserByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
