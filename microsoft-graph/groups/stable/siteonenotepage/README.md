
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotepage` Documentation

The `siteonenotepage` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteonenotepage"
```


### Client Initialization

```go
client := siteonenotepage.NewSiteOnenotePageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenotePageClient.CopySiteOnenotePageToSection`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteIdOnenotePageID("groupId", "siteId", "onenotePageId")

payload := siteonenotepage.CopySiteOnenotePageToSectionRequest{
	// ...
}


read, err := client.CopySiteOnenotePageToSection(ctx, id, payload, siteonenotepage.DefaultCopySiteOnenotePageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.CreateSiteOnenotePage`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteID("groupId", "siteId")

payload := siteonenotepage.OnenotePage{
	// ...
}


read, err := client.CreateSiteOnenotePage(ctx, id, payload, siteonenotepage.DefaultCreateSiteOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.CreateSiteOnenotePageOnenotePatchContent`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteIdOnenotePageID("groupId", "siteId", "onenotePageId")

payload := siteonenotepage.CreateSiteOnenotePageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateSiteOnenotePageOnenotePatchContent(ctx, id, payload, siteonenotepage.DefaultCreateSiteOnenotePageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.DeleteSiteOnenotePage`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteIdOnenotePageID("groupId", "siteId", "onenotePageId")

read, err := client.DeleteSiteOnenotePage(ctx, id, siteonenotepage.DefaultDeleteSiteOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.GetSiteOnenotePage`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteIdOnenotePageID("groupId", "siteId", "onenotePageId")

read, err := client.GetSiteOnenotePage(ctx, id, siteonenotepage.DefaultGetSiteOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.GetSiteOnenotePagesCount`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteOnenotePagesCount(ctx, id, siteonenotepage.DefaultGetSiteOnenotePagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenotePageClient.ListSiteOnenotePages`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteOnenotePages(ctx, id, siteonenotepage.DefaultListSiteOnenotePagesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenotePagesComplete(ctx, id, siteonenotepage.DefaultListSiteOnenotePagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenotePageClient.UpdateSiteOnenotePage`

```go
ctx := context.TODO()
id := siteonenotepage.NewGroupIdSiteIdOnenotePageID("groupId", "siteId", "onenotePageId")

payload := siteonenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateSiteOnenotePage(ctx, id, payload, siteonenotepage.DefaultUpdateSiteOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
