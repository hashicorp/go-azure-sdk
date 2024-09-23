
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectionpage` Documentation

The `siteonenotesectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesectionpage"
```


### Client Initialization

```go
client := siteonenotesectionpage.NewSiteOnenoteSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteSectionPageClient.CopySiteOnenoteSectionPageToSection`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionIdPageID("groupId", "siteId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectionpage.CopySiteOnenoteSectionPageToSectionRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionPageToSection(ctx, id, payload, siteonenotesectionpage.DefaultCopySiteOnenoteSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.CreateSiteOnenoteSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

payload := siteonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateSiteOnenoteSectionPage(ctx, id, payload, siteonenotesectionpage.DefaultCreateSiteOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.CreateSiteOnenoteSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionIdPageID("groupId", "siteId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectionpage.CreateSiteOnenoteSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateSiteOnenoteSectionPageOnenotePatchContent(ctx, id, payload, siteonenotesectionpage.DefaultCreateSiteOnenoteSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.DeleteSiteOnenoteSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionIdPageID("groupId", "siteId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteSiteOnenoteSectionPage(ctx, id, siteonenotesectionpage.DefaultDeleteSiteOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.GetSiteOnenoteSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionIdPageID("groupId", "siteId", "onenoteSectionId", "onenotePageId")

read, err := client.GetSiteOnenoteSectionPage(ctx, id, siteonenotesectionpage.DefaultGetSiteOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.GetSiteOnenoteSectionPagesCount`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

read, err := client.GetSiteOnenoteSectionPagesCount(ctx, id, siteonenotesectionpage.DefaultGetSiteOnenoteSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionPageClient.ListSiteOnenoteSectionPages`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

// alternatively `client.ListSiteOnenoteSectionPages(ctx, id, siteonenotesectionpage.DefaultListSiteOnenoteSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteSectionPagesComplete(ctx, id, siteonenotesectionpage.DefaultListSiteOnenoteSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteSectionPageClient.UpdateSiteOnenoteSectionPage`

```go
ctx := context.TODO()
id := siteonenotesectionpage.NewGroupIdSiteIdOnenoteSectionIdPageID("groupId", "siteId", "onenoteSectionId", "onenotePageId")

payload := siteonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateSiteOnenoteSectionPage(ctx, id, payload, siteonenotesectionpage.DefaultUpdateSiteOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
