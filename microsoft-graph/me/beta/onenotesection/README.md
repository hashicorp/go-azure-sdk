
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onenotesection` Documentation

The `onenotesection` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onenotesection"
```


### Client Initialization

```go
client := onenotesection.NewOnenoteSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteSectionClient.CopyOnenoteSectionToNotebook`

```go
ctx := context.TODO()
id := onenotesection.NewMeOnenoteSectionID("onenoteSectionId")

payload := onenotesection.CopyOnenoteSectionToNotebookRequest{
	// ...
}


read, err := client.CopyOnenoteSectionToNotebook(ctx, id, payload, onenotesection.DefaultCopyOnenoteSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.CopyOnenoteSectionToSectionGroup`

```go
ctx := context.TODO()
id := onenotesection.NewMeOnenoteSectionID("onenoteSectionId")

payload := onenotesection.CopyOnenoteSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopyOnenoteSectionToSectionGroup(ctx, id, payload, onenotesection.DefaultCopyOnenoteSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.CreateOnenoteSection`

```go
ctx := context.TODO()

payload := onenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateOnenoteSection(ctx, payload, onenotesection.DefaultCreateOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.DeleteOnenoteSection`

```go
ctx := context.TODO()
id := onenotesection.NewMeOnenoteSectionID("onenoteSectionId")

read, err := client.DeleteOnenoteSection(ctx, id, onenotesection.DefaultDeleteOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.GetOnenoteSection`

```go
ctx := context.TODO()
id := onenotesection.NewMeOnenoteSectionID("onenoteSectionId")

read, err := client.GetOnenoteSection(ctx, id, onenotesection.DefaultGetOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.GetOnenoteSectionsCount`

```go
ctx := context.TODO()


read, err := client.GetOnenoteSectionsCount(ctx, onenotesection.DefaultGetOnenoteSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteSectionClient.ListOnenoteSections`

```go
ctx := context.TODO()


// alternatively `client.ListOnenoteSections(ctx, onenotesection.DefaultListOnenoteSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteSectionsComplete(ctx, onenotesection.DefaultListOnenoteSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteSectionClient.UpdateOnenoteSection`

```go
ctx := context.TODO()
id := onenotesection.NewMeOnenoteSectionID("onenoteSectionId")

payload := onenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateOnenoteSection(ctx, id, payload, onenotesection.DefaultUpdateOnenoteSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
