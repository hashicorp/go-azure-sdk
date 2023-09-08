
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectiongroupsection` Documentation

The `grouponenotenotebooksectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/grouponenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSection`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupID("groupIdValue", "notebookIdValue", "sectionGroupIdValue")

payload := grouponenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksectiongroupsection.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksectiongroupsection.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.DeleteGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupID("groupIdValue", "notebookIdValue", "sectionGroupIdValue")

read, err := client.GetGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSections`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupID("groupIdValue", "notebookIdValue", "sectionGroupIdValue")

// alternatively `client.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteNotebookSectionGroupSectionClient.UpdateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionID("groupIdValue", "notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteNotebookByIdSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
