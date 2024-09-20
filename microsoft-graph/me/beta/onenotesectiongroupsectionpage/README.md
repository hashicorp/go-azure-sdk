
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onenotesectiongroupsectionpage` Documentation

The `onenotesectiongroupsectionpage` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onenotesectiongroupsectionpage"
```


### Client Initialization

```go
client := onenotesectiongroupsectionpage.NewOnenoteSectionGroupSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.CopyOnenoteSectionGroupSectionPageToSection`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionIdPageID("sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectiongroupsectionpage.CopyOnenoteSectionGroupSectionPageToSectionRequest{
	// ...
}


read, err := client.CopyOnenoteSectionGroupSectionPageToSection(ctx, id, payload, onenotesectiongroupsectionpage.DefaultCopyOnenoteSectionGroupSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.CreateOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionID("sectionGroupId", "onenoteSectionId")

payload := onenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.CreateOnenoteSectionGroupSectionPage(ctx, id, payload, onenotesectiongroupsectionpage.DefaultCreateOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.CreateOnenoteSectionGroupSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionIdPageID("sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectiongroupsectionpage.CreateOnenoteSectionGroupSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateOnenoteSectionGroupSectionPageOnenotePatchContent(ctx, id, payload, onenotesectiongroupsectionpage.DefaultCreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.DeleteOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionIdPageID("sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteOnenoteSectionGroupSectionPage(ctx, id, onenotesectiongroupsectionpage.DefaultDeleteOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.GetOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionIdPageID("sectionGroupId", "onenoteSectionId", "onenotePageId")

read, err := client.GetOnenoteSectionGroupSectionPage(ctx, id, onenotesectiongroupsectionpage.DefaultGetOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.GetOnenoteSectionGroupSectionPagesCount`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionID("sectionGroupId", "onenoteSectionId")

read, err := client.GetOnenoteSectionGroupSectionPagesCount(ctx, id, onenotesectiongroupsectionpage.DefaultGetOnenoteSectionGroupSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.ListOnenoteSectionGroupSectionPages`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionID("sectionGroupId", "onenoteSectionId")

// alternatively `client.ListOnenoteSectionGroupSectionPages(ctx, id, onenotesectiongroupsectionpage.DefaultListOnenoteSectionGroupSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteSectionGroupSectionPagesComplete(ctx, id, onenotesectiongroupsectionpage.DefaultListOnenoteSectionGroupSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteSectionGroupSectionPageClient.UpdateOnenoteSectionGroupSectionPage`

```go
ctx := context.TODO()
id := onenotesectiongroupsectionpage.NewMeOnenoteSectionGroupIdSectionIdPageID("sectionGroupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectiongroupsectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateOnenoteSectionGroupSectionPage(ctx, id, payload, onenotesectiongroupsectionpage.DefaultUpdateOnenoteSectionGroupSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
