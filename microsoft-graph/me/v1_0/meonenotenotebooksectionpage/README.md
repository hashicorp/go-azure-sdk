
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotenotebooksectionpage` Documentation

The `meonenotenotebooksectionpage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotenotebooksectionpage"
```


### Client Initialization

```go
client := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.CreateMeOnenoteNotebookByIdSectionByIdPage`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.CreateMeOnenoteNotebookByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageID("notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectionpage.CreateMeOnenoteNotebookByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.CreateMeOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageID("notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectionpage.CreateMeOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.DeleteMeOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageID("notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteMeOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.GetMeOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageID("notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.GetMeOnenoteNotebookByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.ListMeOnenoteNotebookByIdSectionByIdPages`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionID("notebookIdValue", "onenoteSectionIdValue")

// alternatively `client.ListMeOnenoteNotebookByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteNotebookByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteNotebookSectionPageClient.UpdateMeOnenoteNotebookByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageID("notebookIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateMeOnenoteNotebookByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
