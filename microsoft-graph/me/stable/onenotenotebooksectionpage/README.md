
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksectionpage` Documentation

The `onenotenotebooksectionpage` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksectionpage"
```


### Client Initialization

```go
client := onenotenotebooksectionpage.NewOnenoteNotebookSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteNotebookSectionPageClient.CopyOnenoteNotebookSectionPageToSection`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionIdPageID("notebookId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectionpage.CopyOnenoteNotebookSectionPageToSectionRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionPageToSection(ctx, id, payload, onenotenotebooksectionpage.DefaultCopyOnenoteNotebookSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.CreateOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

payload := onenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateOnenoteNotebookSectionPage(ctx, id, payload, onenotenotebooksectionpage.DefaultCreateOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.CreateOnenoteNotebookSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionIdPageID("notebookId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectionpage.CreateOnenoteNotebookSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateOnenoteNotebookSectionPageOnenotePatchContent(ctx, id, payload, onenotenotebooksectionpage.DefaultCreateOnenoteNotebookSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.DeleteOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionIdPageID("notebookId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteOnenoteNotebookSectionPage(ctx, id, onenotenotebooksectionpage.DefaultDeleteOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.GetOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionIdPageID("notebookId", "onenoteSectionId", "onenotePageId")

read, err := client.GetOnenoteNotebookSectionPage(ctx, id, onenotenotebooksectionpage.DefaultGetOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.GetOnenoteNotebookSectionPagesCount`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

read, err := client.GetOnenoteNotebookSectionPagesCount(ctx, id, onenotenotebooksectionpage.DefaultGetOnenoteNotebookSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.ListOnenoteNotebookSectionPages`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

// alternatively `client.ListOnenoteNotebookSectionPages(ctx, id, onenotenotebooksectionpage.DefaultListOnenoteNotebookSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteNotebookSectionPagesComplete(ctx, id, onenotenotebooksectionpage.DefaultListOnenoteNotebookSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteNotebookSectionPageClient.UpdateOnenoteNotebookSectionPage`

```go
ctx := context.TODO()
id := onenotenotebooksectionpage.NewMeOnenoteNotebookIdSectionIdPageID("notebookId", "onenoteSectionId", "onenotePageId")

payload := onenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateOnenoteNotebookSectionPage(ctx, id, payload, onenotenotebooksectionpage.DefaultUpdateOnenoteNotebookSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
