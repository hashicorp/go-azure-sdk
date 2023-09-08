
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectiongroupsectionpage` Documentation

The `grouponenotesectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectiongroupsectionpage.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectiongroupsectionpage.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.DeleteGroupByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.GetGroupByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.GetGroupByIdOnenoteSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.ListGroupByIdOnenoteSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdOnenoteSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionPageClient.UpdateGroupByIdOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
