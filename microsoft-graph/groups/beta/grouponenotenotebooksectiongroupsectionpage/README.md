
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectiongroupsectionpage` Documentation

The `grouponenotenotebooksectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectiongroupsectionpage.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectiongroupsectionpage.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.DeleteGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionPageClient.UpdateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
