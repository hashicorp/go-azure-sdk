
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/onenotenotebooksectiongroupsectionpage` Documentation

The `onenotenotebooksectiongroupsectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/onenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := onenotenotebooksectiongroupsectionpage.NewOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.CopyOnenoteNotebookSectionGroupSectionPageToSection`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectiongroupsectionpage.CopyOnenoteNotebookSectionGroupSectionPageToSectionRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionGroupSectionPageToSection(ctx, id, payload, onenotenotebooksectiongroupsectionpage.DefaultCopyOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.CreateOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId")

payload := onenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateOnenoteNotebookSectionGroupSectionPage(ctx, id, payload, onenotenotebooksectiongroupsectionpage.DefaultCreateOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.CreateOnenoteNotebookSectionGroupSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectiongroupsectionpage.CreateOnenoteNotebookSectionGroupSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateOnenoteNotebookSectionGroupSectionPageOnenotePatchContent(ctx, id, payload, onenotenotebooksectiongroupsectionpage.DefaultCreateOnenoteNotebookSectionGroupSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.DeleteOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteOnenoteNotebookSectionGroupSectionPage(ctx, id, onenotenotebooksectiongroupsectionpage.DefaultDeleteOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.GetOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.GetOnenoteNotebookSectionGroupSectionPage(ctx, id, onenotenotebooksectiongroupsectionpage.DefaultGetOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.GetOnenoteNotebookSectionGroupSectionPagesCount`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId")

read, err := client.GetOnenoteNotebookSectionGroupSectionPagesCount(ctx, id, onenotenotebooksectiongroupsectionpage.DefaultGetOnenoteNotebookSectionGroupSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.ListOnenoteNotebookSectionGroupSectionPages`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId")

// alternatively `client.ListOnenoteNotebookSectionGroupSectionPages(ctx, id, onenotenotebooksectiongroupsectionpage.DefaultListOnenoteNotebookSectionGroupSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteNotebookSectionGroupSectionPagesComplete(ctx, id, onenotenotebooksectiongroupsectionpage.DefaultListOnenoteNotebookSectionGroupSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteNotebookSectionGroupSectionPageClient.UpdateOnenoteNotebookSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectiongroupsectionpage.NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID("groupId", "notebookId", "sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateOnenoteNotebookSectionGroupSectionPage(ctx, id, payload, onenotenotebooksectiongroupsectionpage.DefaultUpdateOnenoteNotebookSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
