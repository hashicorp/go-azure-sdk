
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectionpage` Documentation

The `meonenotesectionpage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonenotesectionpage"
```


### Client Initialization

```go
client := meonenotesectionpage.NewMeOnenoteSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteSectionPageClient.CreateMeOnenoteSectionByIdPage`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionID("onenoteSectionIdValue")

payload := meonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateMeOnenoteSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.CreateMeOnenoteSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionPageID("onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectionpage.CreateMeOnenoteSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.CreateMeOnenoteSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionPageID("onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectionpage.CreateMeOnenoteSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.DeleteMeOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionPageID("onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteMeOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.GetMeOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionPageID("onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetMeOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.GetMeOnenoteSectionByIdPageCount`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionID("onenoteSectionIdValue")

read, err := client.GetMeOnenoteSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionPageClient.ListMeOnenoteSectionByIdPages`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionID("onenoteSectionIdValue")

// alternatively `client.ListMeOnenoteSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeOnenoteSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteSectionPageClient.UpdateMeOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := meonenotesectionpage.NewMeOnenoteSectionPageID("onenoteSectionIdValue", "onenotePageIdValue")

payload := meonenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateMeOnenoteSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
