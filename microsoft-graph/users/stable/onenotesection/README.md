
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onenotesection` Documentation

The `onenotesection` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onenotesection"
```


### Client Initialization

```go
client := onenotesection.NewOnenoteSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteSectionClient.CopyOnenoteSectionToNotebook`

```go
ctx := context.TODO()
id := onenotesection.NewUserIdOnenoteSectionID("userId", "onenoteSectionId")

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
id := onenotesection.NewUserIdOnenoteSectionID("userId", "onenoteSectionId")

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
id := onenotesection.NewUserID("userId")

payload := onenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateOnenoteSection(ctx, id, payload, onenotesection.DefaultCreateOnenoteSectionOperationOptions())
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
id := onenotesection.NewUserIdOnenoteSectionID("userId", "onenoteSectionId")

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
id := onenotesection.NewUserIdOnenoteSectionID("userId", "onenoteSectionId")

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
id := onenotesection.NewUserID("userId")

read, err := client.GetOnenoteSectionsCount(ctx, id, onenotesection.DefaultGetOnenoteSectionsCountOperationOptions())
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
id := onenotesection.NewUserID("userId")

// alternatively `client.ListOnenoteSections(ctx, id, onenotesection.DefaultListOnenoteSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteSectionsComplete(ctx, id, onenotesection.DefaultListOnenoteSectionsOperationOptions())
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
id := onenotesection.NewUserIdOnenoteSectionID("userId", "onenoteSectionId")

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
