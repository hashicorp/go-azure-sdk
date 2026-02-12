
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/notebooks` Documentation

The `notebooks` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/notebooks"
```


### Client Initialization

```go
client := notebooks.NewNotebooksClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `NotebooksClient.NotebookCreateOrUpdateNotebook`

```go
ctx := context.TODO()
id := notebooks.NewNotebookID("notebookName")

payload := notebooks.NotebookResource{
	// ...
}


if err := client.NotebookCreateOrUpdateNotebookThenPoll(ctx, id, payload, notebooks.DefaultNotebookCreateOrUpdateNotebookOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NotebooksClient.NotebookDeleteNotebook`

```go
ctx := context.TODO()
id := notebooks.NewNotebookID("notebookName")

if err := client.NotebookDeleteNotebookThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NotebooksClient.NotebookGetNotebook`

```go
ctx := context.TODO()
id := notebooks.NewNotebookID("notebookName")

read, err := client.NotebookGetNotebook(ctx, id, notebooks.DefaultNotebookGetNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotebooksClient.NotebookGetNotebookSummaryByWorkSpace`

```go
ctx := context.TODO()


// alternatively `client.NotebookGetNotebookSummaryByWorkSpace(ctx)` can be used to do batched pagination
items, err := client.NotebookGetNotebookSummaryByWorkSpaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NotebooksClient.NotebookGetNotebooksByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.NotebookGetNotebooksByWorkspace(ctx)` can be used to do batched pagination
items, err := client.NotebookGetNotebooksByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NotebooksClient.NotebookRenameNotebook`

```go
ctx := context.TODO()
id := notebooks.NewNotebookID("notebookName")

payload := notebooks.ArtifactRenameRequest{
	// ...
}


if err := client.NotebookRenameNotebookThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
