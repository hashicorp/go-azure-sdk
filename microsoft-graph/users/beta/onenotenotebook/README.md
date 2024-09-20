
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebook` Documentation

The `onenotenotebook` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebook"
```


### Client Initialization

```go
client := onenotenotebook.NewOnenoteNotebookClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenoteNotebookClient.CopyOnenoteNotebookNotebook`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserIdOnenoteNotebookID("userId", "notebookId")

payload := onenotenotebook.CopyOnenoteNotebookNotebookRequest{
	// ...
}


read, err := client.CopyOnenoteNotebookNotebook(ctx, id, payload, onenotenotebook.DefaultCopyOnenoteNotebookNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.CreateOnenoteNotebook`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserID("userId")

payload := onenotenotebook.Notebook{
	// ...
}


read, err := client.CreateOnenoteNotebook(ctx, id, payload, onenotenotebook.DefaultCreateOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.DeleteOnenoteNotebook`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserIdOnenoteNotebookID("userId", "notebookId")

read, err := client.DeleteOnenoteNotebook(ctx, id, onenotenotebook.DefaultDeleteOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.GetOnenoteNotebook`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserIdOnenoteNotebookID("userId", "notebookId")

read, err := client.GetOnenoteNotebook(ctx, id, onenotenotebook.DefaultGetOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.GetOnenoteNotebooksCount`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserID("userId")

read, err := client.GetOnenoteNotebooksCount(ctx, id, onenotenotebook.DefaultGetOnenoteNotebooksCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.GetOnenoteNotebooksFromWebUrl`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserID("userId")

payload := onenotenotebook.GetOnenoteNotebooksFromWebUrlRequest{
	// ...
}


read, err := client.GetOnenoteNotebooksFromWebUrl(ctx, id, payload, onenotenotebook.DefaultGetOnenoteNotebooksFromWebUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenoteNotebookClient.ListOnenoteNotebooks`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserID("userId")

// alternatively `client.ListOnenoteNotebooks(ctx, id, onenotenotebook.DefaultListOnenoteNotebooksOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenoteNotebooksComplete(ctx, id, onenotenotebook.DefaultListOnenoteNotebooksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenoteNotebookClient.UpdateOnenoteNotebook`

```go
ctx := context.TODO()
id := onenotenotebook.NewUserIdOnenoteNotebookID("userId", "notebookId")

payload := onenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateOnenoteNotebook(ctx, id, payload, onenotenotebook.DefaultUpdateOnenoteNotebookOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
