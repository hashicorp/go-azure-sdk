
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectiongroupsectionpage` Documentation

The `meonenotesectiongroupsectionpage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.CreateMeOnenoteSectionGroupByIdSectionByIdPage`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

payload := meonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageID("sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectiongroupsectionpage.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageID("sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectiongroupsectionpage.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionGroupByIdSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.DeleteMeOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageID("sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteMeOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.GetMeOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageID("sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetMeOnenoteSectionGroupByIdSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.GetMeOnenoteSectionGroupByIdSectionByIdPageCount`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

read, err := client.GetMeOnenoteSectionGroupByIdSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.ListMeOnenoteSectionGroupByIdSectionByIdPages`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionID("sectionGroupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListMeOnenoteSectionGroupByIdSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteSectionGroupByIdSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteSectionGroupSectionPageClient.UpdateMeOnenoteSectionGroupByIdSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageID("sectionGroupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateMeOnenoteSectionGroupByIdSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
