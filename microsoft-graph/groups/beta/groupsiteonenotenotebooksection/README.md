
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksection` Documentation

The `groupsiteonenotenotebooksection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotenotebooksection"
```


### Client Initialization

```go
client := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSection`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

payload := groupsiteonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksection.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksection.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionCount`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteNotebookByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.ListGroupByIdSiteByIdOnenoteNotebookByIdSections`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookID("groupIdValue", "siteIdValue", "notebookIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteNotebookByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteNotebookByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteNotebookSectionClient.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionID("groupIdValue", "siteIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteNotebookByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
