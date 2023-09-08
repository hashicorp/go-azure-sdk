
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesectionpage` Documentation

The `groupsiteonenotesectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesectionpage"
```


### Client Initialization

```go
client := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionByIdPage`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageID("groupIdValue", "siteIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectionpage.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageID("groupIdValue", "siteIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectionpage.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.DeleteGroupByIdSiteByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageID("groupIdValue", "siteIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.GetGroupByIdSiteByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageID("groupIdValue", "siteIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.GetGroupByIdSiteByIdOnenoteSectionByIdPageCount`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.ListGroupByIdSiteByIdOnenoteSectionByIdPages`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteSectionPageClient.UpdateGroupByIdSiteByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageID("groupIdValue", "siteIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := groupsiteonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
