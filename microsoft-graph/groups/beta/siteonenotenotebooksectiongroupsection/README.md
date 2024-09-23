
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebooksectiongroupsection` Documentation

The `siteonenotenotebooksectiongroupsection` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebooksectiongroupsection"
```


### Client Initialization

```go
client := siteonenotenotebooksectiongroupsection.NewSiteOnenoteNotebookSectionGroupSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.CopySiteOnenoteNotebookSectionGroupSectionToNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotenotebooksectiongroupsection.CopySiteOnenoteNotebookSectionGroupSectionToNotebookRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionGroupSectionToNotebook(ctx, id, payload, siteonenotenotebooksectiongroupsection.DefaultCopySiteOnenoteNotebookSectionGroupSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.CopySiteOnenoteNotebookSectionGroupSectionToSectionGroup`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotenotebooksectiongroupsection.CopySiteOnenoteNotebookSectionGroupSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionGroupSectionToSectionGroup(ctx, id, payload, siteonenotenotebooksectiongroupsection.DefaultCopySiteOnenoteNotebookSectionGroupSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.CreateSiteOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupID("groupId", "siteId", "notebookId", "sectionGroupId")

payload := siteonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSectionGroupSection(ctx, id, payload, siteonenotenotebooksectiongroupsection.DefaultCreateSiteOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.DeleteSiteOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.DeleteSiteOnenoteNotebookSectionGroupSection(ctx, id, siteonenotenotebooksectiongroupsection.DefaultDeleteSiteOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.GetSiteOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetSiteOnenoteNotebookSectionGroupSection(ctx, id, siteonenotenotebooksectiongroupsection.DefaultGetSiteOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.GetSiteOnenoteNotebookSectionGroupSectionsCount`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupID("groupId", "siteId", "notebookId", "sectionGroupId")

read, err := client.GetSiteOnenoteNotebookSectionGroupSectionsCount(ctx, id, siteonenotenotebooksectiongroupsection.DefaultGetSiteOnenoteNotebookSectionGroupSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.ListSiteOnenoteNotebookSectionGroupSections`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupID("groupId", "siteId", "notebookId", "sectionGroupId")

// alternatively `client.ListSiteOnenoteNotebookSectionGroupSections(ctx, id, siteonenotenotebooksectiongroupsection.DefaultListSiteOnenoteNotebookSectionGroupSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteNotebookSectionGroupSectionsComplete(ctx, id, siteonenotenotebooksectiongroupsection.DefaultListSiteOnenoteNotebookSectionGroupSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteNotebookSectionGroupSectionClient.UpdateSiteOnenoteNotebookSectionGroupSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectiongroupsection.NewGroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "siteId", "notebookId", "sectionGroupId", "onenoteSectionId")

payload := siteonenotenotebooksectiongroupsection.OnenoteSection{
	// ...
}


read, err := client.UpdateSiteOnenoteNotebookSectionGroupSection(ctx, id, payload, siteonenotenotebooksectiongroupsection.DefaultUpdateSiteOnenoteNotebookSectionGroupSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
