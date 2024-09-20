
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectiongroupsectionpage` Documentation

The `siteonenotesectiongroupsectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := siteonenotesectiongroupsectionpage.NewSiteOnenoteSectionGroupSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.CopySiteOnenoteSectionGroupSectionPageToSection`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionIdPageID("groupId", "siteId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectiongroupsectionpage.CopySiteOnenoteSectionGroupSectionPageToSectionRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionGroupSectionPageToSection(ctx, id, payload, siteonenotesectiongroupsectionpage.DefaultCopySiteOnenoteSectionGroupSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.CreateSiteOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateSiteOnenoteSectionGroupSectionPage(ctx, id, payload, siteonenotesectiongroupsectionpage.DefaultCreateSiteOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.CreateSiteOnenoteSectionGroupSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionIdPageID("groupId", "siteId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectiongroupsectionpage.CreateSiteOnenoteSectionGroupSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateSiteOnenoteSectionGroupSectionPageOnenotePatchContent(ctx, id, payload, siteonenotesectiongroupsectionpage.DefaultCreateSiteOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.DeleteSiteOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionIdPageID("groupId", "siteId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteSiteOnenoteSectionGroupSectionPage(ctx, id, siteonenotesectiongroupsectionpage.DefaultDeleteSiteOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.GetSiteOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionIdPageID("groupId", "siteId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.GetSiteOnenoteSectionGroupSectionPage(ctx, id, siteonenotesectiongroupsectionpage.DefaultGetSiteOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.GetSiteOnenoteSectionGroupSectionPagesCount`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetSiteOnenoteSectionGroupSectionPagesCount(ctx, id, siteonenotesectiongroupsectionpage.DefaultGetSiteOnenoteSectionGroupSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.ListSiteOnenoteSectionGroupSectionPages`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

// alternatively `client.ListSiteOnenoteSectionGroupSectionPages(ctx, id, siteonenotesectiongroupsectionpage.DefaultListSiteOnenoteSectionGroupSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteSectionGroupSectionPagesComplete(ctx, id, siteonenotesectiongroupsectionpage.DefaultListSiteOnenoteSectionGroupSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionPageClient.UpdateSiteOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsectionpage.NewGroupIdSiteIdOnenoteSectionGroupIdSectionIdPageID("groupId", "siteId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateSiteOnenoteSectionGroupSectionPage(ctx, id, payload, siteonenotesectiongroupsectionpage.DefaultUpdateSiteOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
