
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebook` Documentation

The `meonenotenotebook` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotenotebook"
```


### Client Initialization

```go
client := meonenotenotebook.NewMeOnenoteNotebookClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteNotebookClient.CreateMeOnenoteNotebook`

```go
ctx := context.TODO()

payload := meonenotenotebook.Notebook{
	// ...
}


read, err := client.CreateMeOnenoteNotebook(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.CreateMeOnenoteNotebookByIdCopyNotebook`

```go
ctx := context.TODO()
id := meonenotenotebook.NewMeOnenoteNotebookID("notebookIdValue")

payload := meonenotenotebook.CreateMeOnenoteNotebookByIdCopyNotebookRequest{
	// ...
}


read, err := client.CreateMeOnenoteNotebookByIdCopyNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.DeleteMeOnenoteNotebookById`

```go
ctx := context.TODO()
id := meonenotenotebook.NewMeOnenoteNotebookID("notebookIdValue")

read, err := client.DeleteMeOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.GetMeOnenoteNotebookById`

```go
ctx := context.TODO()
id := meonenotenotebook.NewMeOnenoteNotebookID("notebookIdValue")

read, err := client.GetMeOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.GetMeOnenoteNotebookCount`

```go
ctx := context.TODO()


read, err := client.GetMeOnenoteNotebookCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.GetMeOnenoteNotebooksNotebookFromWebUrl`

```go
ctx := context.TODO()

payload := meonenotenotebook.GetMeOnenoteNotebooksNotebookFromWebUrlRequest{
	// ...
}


read, err := client.GetMeOnenoteNotebooksNotebookFromWebUrl(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteNotebookClient.ListMeOnenoteNotebooks`

```go
ctx := context.TODO()


// alternatively `client.ListMeOnenoteNotebooks(ctx)` can be used to do batched pagination
items, err := client.ListMeOnenoteNotebooksComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteNotebookClient.UpdateMeOnenoteNotebookById`

```go
ctx := context.TODO()
id := meonenotenotebook.NewMeOnenoteNotebookID("notebookIdValue")

payload := meonenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateMeOnenoteNotebookById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
