
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebooksectionpage` Documentation

The `siteonenotenotebooksectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotenotebooksectionpage"
```


### Client Initialization

```go
client := siteonenotenotebooksectionpage.NewSiteOnenoteNotebookSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.CopySiteOnenoteNotebookSectionPageToSection`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID("groupId", "siteId", "notebookId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectionpage.CopySiteOnenoteNotebookSectionPageToSectionRequest{
	// ...
}


read, err := client.CopySiteOnenoteNotebookSectionPageToSection(ctx, id, payload, siteonenotenotebooksectionpage.DefaultCopySiteOnenoteNotebookSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.CreateSiteOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

payload := siteonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSectionPage(ctx, id, payload, siteonenotenotebooksectionpage.DefaultCreateSiteOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.CreateSiteOnenoteNotebookSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID("groupId", "siteId", "notebookId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectionpage.CreateSiteOnenoteNotebookSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateSiteOnenoteNotebookSectionPageOnenotePatchContent(ctx, id, payload, siteonenotenotebooksectionpage.DefaultCreateSiteOnenoteNotebookSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.DeleteSiteOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID("groupId", "siteId", "notebookId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteSiteOnenoteNotebookSectionPage(ctx, id, siteonenotenotebooksectionpage.DefaultDeleteSiteOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.GetSiteOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID("groupId", "siteId", "notebookId", "onenoteSectionId", "onenotePageId")

read, err := client.GetSiteOnenoteNotebookSectionPage(ctx, id, siteonenotenotebooksectionpage.DefaultGetSiteOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.GetSiteOnenoteNotebookSectionPagesCount`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

read, err := client.GetSiteOnenoteNotebookSectionPagesCount(ctx, id, siteonenotenotebooksectionpage.DefaultGetSiteOnenoteNotebookSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.ListSiteOnenoteNotebookSectionPages`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionID("groupId", "siteId", "notebookId", "onenoteSectionId")

// alternatively `client.ListSiteOnenoteNotebookSectionPages(ctx, id, siteonenotenotebooksectionpage.DefaultListSiteOnenoteNotebookSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteNotebookSectionPagesComplete(ctx, id, siteonenotenotebooksectionpage.DefaultListSiteOnenoteNotebookSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteNotebookSectionPageClient.UpdateSiteOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := siteonenotenotebooksectionpage.NewGroupIdSiteIdOnenoteNotebookIdSectionIdPageID("groupId", "siteId", "notebookId", "onenoteSectionId", "onenotePageId")

payload := siteonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateSiteOnenoteNotebookSectionPage(ctx, id, payload, siteonenotenotebooksectionpage.DefaultUpdateSiteOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
