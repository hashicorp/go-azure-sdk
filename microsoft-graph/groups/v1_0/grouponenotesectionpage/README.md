
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectionpage` Documentation

The `grouponenotesectionpage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesectionpage"
```


### Client Initialization

```go
client := grouponenotesectionpage.NewGroupOnenoteSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteSectionPageClient.CreateGroupByIdOnenoteSectionByIdPage`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

payload := grouponenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.CreateGroupByIdOnenoteSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionPageID("groupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectionpage.CreateGroupByIdOnenoteSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.CreateGroupByIdOnenoteSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionPageID("groupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectionpage.CreateGroupByIdOnenoteSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.DeleteGroupByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionPageID("groupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.GetGroupByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionPageID("groupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.GetGroupByIdOnenoteSectionByIdPageCount`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionPageClient.ListGroupByIdOnenoteSectionByIdPages`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

// alternatively `client.ListGroupByIdOnenoteSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteSectionPageClient.UpdateGroupByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := grouponenotesectionpage.NewGroupOnenoteSectionPageID("groupIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := grouponenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
