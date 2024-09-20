
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotenotebooksection` Documentation

The `siteonenotenotebooksection` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotenotebooksection"
```


### Client Initialization

```go
client := siteonenotenotebooksection.NewSiteOnenoteNotebookSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteNotebookSectionClient.CopySiteOnenoteNotebookSectionToNotebook`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

payload := siteonenotenotebooksection.CopySiteOnenoteNotebookSectionToNotebookRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionToNotebook(ctx, id, payload, siteonenotenotebooksection.DefaultCopySiteOnenoteNotebookSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.CopySiteOnenoteNotebookSectionToSectionGroup`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

payload := siteonenotenotebooksection.CopySiteOnenoteNotebookSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionToSectionGroup(ctx, id, payload, siteonenotenotebooksection.DefaultCopySiteOnenoteNotebookSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.CreateSiteOnenoteNotebookSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

payload := siteonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSection(ctx, id, payload, siteonenotenotebooksection.DefaultCreateSiteOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.DeleteSiteOnenoteNotebookSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

read, err := client.DeleteSiteOnenoteNotebookSection(ctx, id, siteonenotenotebooksection.DefaultDeleteSiteOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.GetSiteOnenoteNotebookSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

read, err := client.GetSiteOnenoteNotebookSection(ctx, id, siteonenotenotebooksection.DefaultGetSiteOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.GetSiteOnenoteNotebookSectionsCount`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

read, err := client.GetSiteOnenoteNotebookSectionsCount(ctx, id, siteonenotenotebooksection.DefaultGetSiteOnenoteNotebookSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.ListSiteOnenoteNotebookSections`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookID("groupId", "siteId", "notebookId")

// alternatively `client.ListSiteOnenoteNotebookSections(ctx, id, siteonenotenotebooksection.DefaultListSiteOnenoteNotebookSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteNotebookSectionsComplete(ctx, id, siteonenotenotebooksection.DefaultListSiteOnenoteNotebookSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteNotebookSectionClient.UpdateSiteOnenoteNotebookSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksection.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

payload := siteonenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateSiteOnenoteNotebookSection(ctx, id, payload, siteonenotenotebooksection.DefaultUpdateSiteOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
