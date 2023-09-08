
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebook` Documentation

The `useronenotenotebook` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebook"
```


### Client Initialization

```go
client := useronenotenotebook.NewUserOnenoteNotebookClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteNotebookClient.CreateUserByIdOnenoteNotebook`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserID("userIdValue")

payload := useronenotenotebook.Notebook{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.CreateUserByIdOnenoteNotebookByIdCopyNotebook`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

payload := useronenotenotebook.CreateUserByIdOnenoteNotebookByIdCopyNotebookRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdCopyNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.DeleteUserByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

read, err := client.DeleteUserByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.GetUserByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

read, err := client.GetUserByIdOnenoteNotebookById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.GetUserByIdOnenoteNotebookCount`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserID("userIdValue")

read, err := client.GetUserByIdOnenoteNotebookCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.GetUserByIdOnenoteNotebooksNotebookFromWebUrl`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserID("userIdValue")

payload := useronenotenotebook.GetUserByIdOnenoteNotebooksNotebookFromWebUrlRequest{
	// ...
}


read, err := client.GetUserByIdOnenoteNotebooksNotebookFromWebUrl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookClient.ListUserByIdOnenoteNotebooks`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserID("userIdValue")

// alternatively `client.ListUserByIdOnenoteNotebooks(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteNotebooksComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteNotebookClient.UpdateUserByIdOnenoteNotebookById`

```go
ctx := context.TODO()
id := useronenotenotebook.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

payload := useronenotenotebook.Notebook{
	// ...
}


read, err := client.UpdateUserByIdOnenoteNotebookById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
