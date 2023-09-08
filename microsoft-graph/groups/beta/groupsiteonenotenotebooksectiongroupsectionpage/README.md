
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectiongroupsectionpage` Documentation

The `groupsiteonenotenotebooksectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectiongroupsectionpage.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectiongroupsectionpage.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionPageClient.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
