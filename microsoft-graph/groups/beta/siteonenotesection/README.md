
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesection` Documentation

The `siteonenotesection` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteonenotesection"
```


### Client Initialization

```go
client := siteonenotesection.NewSiteOnenoteSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteOnenoteSectionClient.CopySiteOnenoteSectionToNotebook`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

payload := siteonenotesection.CopySiteOnenoteSectionToNotebookRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionToNotebook(ctx, id, payload, siteonenotesection.DefaultCopySiteOnenoteSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.CopySiteOnenoteSectionToSectionGroup`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

payload := siteonenotesection.CopySiteOnenoteSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopySiteOnenoteSectionToSectionGroup(ctx, id, payload, siteonenotesection.DefaultCopySiteOnenoteSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.CreateSiteOnenoteSection`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteID("groupId", "siteId")

payload := siteonenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateSiteOnenoteSection(ctx, id, payload, siteonenotesection.DefaultCreateSiteOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.DeleteSiteOnenoteSection`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

read, err := client.DeleteSiteOnenoteSection(ctx, id, siteonenotesection.DefaultDeleteSiteOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.GetSiteOnenoteSection`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

read, err := client.GetSiteOnenoteSection(ctx, id, siteonenotesection.DefaultGetSiteOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.GetSiteOnenoteSectionsCount`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteID("groupId", "siteId")

read, err := client.GetSiteOnenoteSectionsCount(ctx, id, siteonenotesection.DefaultGetSiteOnenoteSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteOnenoteSectionClient.ListSiteOnenoteSections`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteID("groupId", "siteId")

// alternatively `client.ListSiteOnenoteSections(ctx, id, siteonenotesection.DefaultListSiteOnenoteSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteOnenoteSectionsComplete(ctx, id, siteonenotesection.DefaultListSiteOnenoteSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteOnenoteSectionClient.UpdateSiteOnenoteSection`

```go
ctx := context.TODO()
id := siteonenotesection.NewGroupIdSiteIdOnenoteSectionID("groupId", "siteId", "onenoteSectionId")

payload := siteonenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateSiteOnenoteSection(ctx, id, payload, siteonenotesection.DefaultUpdateSiteOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
