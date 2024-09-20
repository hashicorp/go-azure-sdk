
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/onenotesectionpage` Documentation

The `onenotesectionpage` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/onenotesectionpage"
```


### Client Initialization

```go
client := onenotesectionpage.NewOnenoteSectionPageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteSectionPageClient.CopyOnenoteSectionPageToSection`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionIdPageID("groupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectionpage.CopyOnenoteSectionPageToSectionRequest{
	// ...
}


read, err := client.CopyOnenoteSectionPageToSection(ctx, id, payload, onenotesectionpage.DefaultCopyOnenoteSectionPageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.CreateOnenoteSectionPage`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionID("groupId", "onenoteSectionId")

payload := onenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateOnenoteSectionPage(ctx, id, payload, onenotesectionpage.DefaultCreateOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.CreateOnenoteSectionPageOnenotePatchContent`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionIdPageID("groupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectionpage.CreateOnenoteSectionPageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateOnenoteSectionPageOnenotePatchContent(ctx, id, payload, onenotesectionpage.DefaultCreateOnenoteSectionPageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.DeleteOnenoteSectionPage`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionIdPageID("groupId", "onenoteSectionId", "onenotePageId")

read, err := client.DeleteOnenoteSectionPage(ctx, id, onenotesectionpage.DefaultDeleteOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.GetOnenoteSectionPage`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionIdPageID("groupId", "onenoteSectionId", "onenotePageId")

read, err := client.GetOnenoteSectionPage(ctx, id, onenotesectionpage.DefaultGetOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.GetOnenoteSectionPagesCount`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionID("groupId", "onenoteSectionId")

read, err := client.GetOnenoteSectionPagesCount(ctx, id, onenotesectionpage.DefaultGetOnenoteSectionPagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionPageClient.ListOnenoteSectionPages`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionID("groupId", "onenoteSectionId")

// alternatively `client.ListOnenoteSectionPages(ctx, id, onenotesectionpage.DefaultListOnenoteSectionPagesOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteSectionPagesComplete(ctx, id, onenotesectionpage.DefaultListOnenoteSectionPagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteSectionPageClient.UpdateOnenoteSectionPage`

```go
ctx := context.TODO()
id := onenotesectionpage.NewGroupIdOnenoteSectionIdPageID("groupId", "onenoteSectionId", "onenotePageId")

payload := onenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateOnenoteSectionPage(ctx, id, payload, onenotesectionpage.DefaultUpdateOnenoteSectionPageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
