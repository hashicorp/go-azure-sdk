
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectiongroupsection` Documentation

The `groupsiteonenotenotebooksectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSection`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue")

payload := groupsiteonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksectiongroupsection.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksectiongroupsection.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSections`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionGroupSectionClient.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
