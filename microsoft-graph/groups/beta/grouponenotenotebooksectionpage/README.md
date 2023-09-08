
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectionpage` Documentation

The `grouponenotenotebooksectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectionpage"
```


### Client Initialization

```go
client := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionByIdPage`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectionpage.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectionpage.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.DeleteGroupByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.GetGroupByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.GetGroupByIdOnenoteNotebookByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.ListGroupByIdOnenoteNotebookByIdSectionByIdPages`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteNotebookSectionPageClient.UpdateGroupByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageID("groupIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
