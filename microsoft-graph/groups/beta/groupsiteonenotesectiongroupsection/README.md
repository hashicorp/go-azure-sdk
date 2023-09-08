
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesectiongroupsection` Documentation

The `groupsiteonenotesectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesectiongroupsection"
```


### Client Initialization

```go
client := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSection`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupID("groupIdValue", "siteIdValue", "sectionGroupIdValue")

payload := groupsiteonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesectiongroupsection.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesectiongroupsection.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.DeleteGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupID("groupIdValue", "siteIdValue", "sectionGroupIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSections`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupID("groupIdValue", "siteIdValue", "sectionGroupIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteSectionGroupSectionClient.UpdateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionID("groupIdValue", "siteIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
