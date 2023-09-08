
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectionpage` Documentation

The `groupsiteonenotenotebooksectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectionpage"
```


### Client Initialization

```go
client := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPage`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectionpage.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectionpage.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPages`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionPageClient.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
