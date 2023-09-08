
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesection` Documentation

The `groupsiteonenotesection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteonenotesection"
```


### Client Initialization

```go
client := groupsiteonenotesection.NewGroupSiteOnenoteSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteOnenoteSectionClient.CreateGroupByIdSiteByIdOnenoteSection`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteonenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesection.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesection.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdOnenoteSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.DeleteGroupByIdSiteByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdSiteByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.GetGroupByIdSiteByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.GetGroupByIdSiteByIdOnenoteSectionCount`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdOnenoteSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.ListGroupByIdSiteByIdOnenoteSections`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdOnenoteSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdOnenoteSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteOnenoteSectionClient.UpdateGroupByIdSiteByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := groupsiteonenotesection.NewGroupSiteOnenoteSectionID("groupIdValue", "siteIdValue", "onenoteSectionIdValue")

payload := groupsiteonenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdOnenoteSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
