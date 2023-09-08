
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebooksectiongroupsectionpage` Documentation

The `meonenotenotebooksectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebooksectiongroupsectionpage"
```


### Client Initialization

```go
client := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectiongroupsectionpage.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectiongroupsectionpage.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.DeleteMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.GetMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.GetMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.ListMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteNotebookSectionGroupSectionPageClient.UpdateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageID("notebookIdValue", "sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotenotebooksectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateMeOnenoteNotebookByIdSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
