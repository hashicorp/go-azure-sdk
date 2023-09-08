
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectiongroupsection` Documentation

The `grouponenotesectiongroupsection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectiongroupsection"
```


### Client Initialization

```go
client := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.CreateGroupByIdOnenoteSectionGroupByIdSection`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupID("groupIdValue", "sectionGroupIdValue")

payload := grouponenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotesectiongroupsection.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotesectiongroupsection.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionGroupByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.DeleteGroupByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.GetGroupByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteSectionGroupByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.GetGroupByIdOnenoteSectionGroupByIdSectionCount`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupID("groupIdValue", "sectionGroupIdValue")

read, err := client.GetGroupByIdOnenoteSectionGroupByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.ListGroupByIdOnenoteSectionGroupByIdSections`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupID("groupIdValue", "sectionGroupIdValue")

// alternatively `client.ListGroupByIdOnenoteSectionGroupByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteSectionGroupByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteSectionGroupSectionClient.UpdateGroupByIdOnenoteSectionGroupByIdSectionById`

```go
ctx := context.TODO()
id := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionID("groupIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := grouponenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteSectionGroupByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
