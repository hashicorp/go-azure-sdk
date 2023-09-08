
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotesectiongroupsectionpage` Documentation

The `groupsiteonenotesectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupsiteonenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectiongroupsectionpage.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectiongroupsectionpage.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.DeleteGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionPageClient.UpdateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
