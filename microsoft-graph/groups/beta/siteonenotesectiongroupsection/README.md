
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectiongroupsection` Documentation

The `siteonenotesectiongroupsection` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectiongroupsection"
```


### Client Initialization

```go
client := siteonenotesectiongroupsection.NewSiteOnenoteSectionGroupSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.CopySiteOnenoteSectionGroupSectionToNotebook`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotesectiongroupsection.CopySiteOnenoteSectionGroupSectionToNotebookRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionGroupSectionToNotebook(ctx, id, payload, siteonenotesectiongroupsection.DefaultCopySiteOnenoteSectionGroupSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.CopySiteOnenoteSectionGroupSectionToSectionGroup`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotesectiongroupsection.CopySiteOnenoteSectionGroupSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionGroupSectionToSectionGroup(ctx, id, payload, siteonenotesectiongroupsection.DefaultCopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.CreateSiteOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupID("groupId", "siteId", "sectionGroupId")

payload := siteonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateSiteOnenoteSectionGroupSection(ctx, id, payload, siteonenotesectiongroupsection.DefaultCreateSiteOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.DeleteSiteOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

read, err := client.DeleteSiteOnenoteSectionGroupSection(ctx, id, siteonenotesectiongroupsection.DefaultDeleteSiteOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.GetSiteOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetSiteOnenoteSectionGroupSection(ctx, id, siteonenotesectiongroupsection.DefaultGetSiteOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.GetSiteOnenoteSectionGroupSectionsCount`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupID("groupId", "siteId", "sectionGroupId")

read, err := client.GetSiteOnenoteSectionGroupSectionsCount(ctx, id, siteonenotesectiongroupsection.DefaultGetSiteOnenoteSectionGroupSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.ListSiteOnenoteSectionGroupSections`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupID("groupId", "siteId", "sectionGroupId")

// alternatively `client.ListSiteOnenoteSectionGroupSections(ctx, id, siteonenotesectiongroupsection.DefaultListSiteOnenoteSectionGroupSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteSectionGroupSectionsComplete(ctx, id, siteonenotesectiongroupsection.DefaultListSiteOnenoteSectionGroupSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteSectionGroupSectionClient.UpdateSiteOnenoteSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotesectiongroupsection.NewGroupIdSiteIdOnenoteSectionGroupIdSectionID("groupId", "siteId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotesectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateSiteOnenoteSectionGroupSection(ctx, id, payload, siteonenotesectiongroupsection.DefaultUpdateSiteOnenoteSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
