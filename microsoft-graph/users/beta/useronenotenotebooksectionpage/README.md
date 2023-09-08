
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebooksectionpage` Documentation

The `useronenotenotebooksectionpage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebooksectionpage"
```


### Client Initialization

```go
client := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionByIdPage`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectionpage.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectionpage.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.DeleteUserByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteUserByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.GetUserByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.GetUserByIdOnenoteNotebookByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.ListUserByIdOnenoteNotebookByIdSectionByIdPages`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

// alternatively `client.ListUserByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteNotebookSectionPageClient.UpdateUserByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageID("userIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateUserByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
