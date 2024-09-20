
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotenotebooksectiongroupsectionpage` Documentation

The `siteonenotenotebooksectiongroupsectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := siteonenotenotebooksectiongroupsectionpage.NewSiteOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.CopySiteOnenoteNotebookSectionGroupSectionPageToSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectiongroupsectionpage.CopySiteOnenoteNotebookSectionGroupSectionPageToSectionRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionGroupSectionPageToSection(ctx, id, payload, siteonenotenotebooksectiongroupsectionpage.DefaultCopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.CreateSiteOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSectionGroupSectionPage(ctx, id, payload, siteonenotenotebooksectiongroupsectionpage.DefaultCreateSiteOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.CreateSiteOnenoteNotebookSectionGroupSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectiongroupsectionpage.CreateSiteOnenoteNotebookSectionGroupSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSectionGroupSectionPageOnenotePatchContent(ctx, id, payload, siteonenotenotebooksectiongroupsectionpage.DefaultCreateSiteOnenoteNotebookSectionGroupSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.DeleteSiteOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteSiteOnenoteNotebookSectionGroupSectionPage(ctx, id, siteonenotenotebooksectiongroupsectionpage.DefaultDeleteSiteOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.GetSiteOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.GetSiteOnenoteNotebookSectionGroupSectionPage(ctx, id, siteonenotenotebooksectiongroupsectionpage.DefaultGetSiteOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.GetSiteOnenoteNotebookSectionGroupSectionPagesCount`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetSiteOnenoteNotebookSectionGroupSectionPagesCount(ctx, id, siteonenotenotebooksectiongroupsectionpage.DefaultGetSiteOnenoteNotebookSectionGroupSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.ListSiteOnenoteNotebookSectionGroupSectionPages`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

// alternatively `client.ListSiteOnenoteNotebookSectionGroupSectionPages(ctx, id, siteonenotenotebooksectiongroupsectionpage.DefaultListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteNotebookSectionGroupSectionPagesComplete(ctx, id, siteonenotenotebooksectiongroupsectionpage.DefaultListSiteOnenoteNotebookSectionGroupSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionPageClient.UpdateSiteOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateSiteOnenoteNotebookSectionGroupSectionPage(ctx, id, payload, siteonenotenotebooksectiongroupsectionpage.DefaultUpdateSiteOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
