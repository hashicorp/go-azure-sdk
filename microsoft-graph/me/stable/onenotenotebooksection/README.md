
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksection` Documentation

The `onenotenotebooksection` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotenotebooksection"
```


### Client Initialization

```go
client := onenotenotebooksection.NewOnenoteNotebookSectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteNotebookSectionClient.CopyOnenoteNotebookSectionToNotebook`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

payload := onenotenotebooksection.CopyOnenoteNotebookSectionToNotebookRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionToNotebook(ctx, id, payload, onenotenotebooksection.DefaultCopyOnenoteNotebookSectionToNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.CopyOnenoteNotebookSectionToSectionGroup`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

payload := onenotenotebooksection.CopyOnenoteNotebookSectionToSectionGroupRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookSectionToSectionGroup(ctx, id, payload, onenotenotebooksection.DefaultCopyOnenoteNotebookSectionToSectionGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.CreateOnenoteNotebookSection`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookID("notebookId")

payload := onenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateOnenoteNotebookSection(ctx, id, payload, onenotenotebooksection.DefaultCreateOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.DeleteOnenoteNotebookSection`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

read, err := client.DeleteOnenoteNotebookSection(ctx, id, onenotenotebooksection.DefaultDeleteOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.GetOnenoteNotebookSection`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

read, err := client.GetOnenoteNotebookSection(ctx, id, onenotenotebooksection.DefaultGetOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.GetOnenoteNotebookSectionsCount`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookID("notebookId")

read, err := client.GetOnenoteNotebookSectionsCount(ctx, id, onenotenotebooksection.DefaultGetOnenoteNotebookSectionsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookSectionClient.ListOnenoteNotebookSections`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookID("notebookId")

// alternatively `client.ListOnenoteNotebookSections(ctx, id, onenotenotebooksection.DefaultListOnenoteNotebookSectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteNotebookSectionsComplete(ctx, id, onenotenotebooksection.DefaultListOnenoteNotebookSectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteNotebookSectionClient.UpdateOnenoteNotebookSection`

```go
ctx := context.TODO()
id := onenotenotebooksection.NewMeOnenoteNotebookIdSectionID("notebookId", "onenoteSectionId")

payload := onenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateOnenoteNotebookSection(ctx, id, payload, onenotenotebooksection.DefaultUpdateOnenoteNotebookSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
