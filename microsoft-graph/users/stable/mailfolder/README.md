
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolder` Documentation

The `mailfolder` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolder"
```


### Client Initialization

```go
client := mailfolder.NewMailFolderClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderClient.CopyMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

payload := mailfolder.CopyMailFolderRequest{
	// ...
}


read, err := client.CopyMailFolder(ctx, id, payload, mailfolder.DefaultCopyMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.CreateMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserID("userId")

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.CreateMailFolder(ctx, id, payload, mailfolder.DefaultCreateMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.CreateMailFolderPermanentDelete`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

read, err := client.CreateMailFolderPermanentDelete(ctx, id, mailfolder.DefaultCreateMailFolderPermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.DeleteMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

read, err := client.DeleteMailFolder(ctx, id, mailfolder.DefaultDeleteMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.GetMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

read, err := client.GetMailFolder(ctx, id, mailfolder.DefaultGetMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.GetMailFoldersCount`

```go
ctx := context.TODO()
id := mailfolder.NewUserID("userId")

read, err := client.GetMailFoldersCount(ctx, id, mailfolder.DefaultGetMailFoldersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.ListMailFolders`

```go
ctx := context.TODO()
id := mailfolder.NewUserID("userId")

// alternatively `client.ListMailFolders(ctx, id, mailfolder.DefaultListMailFoldersOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFoldersComplete(ctx, id, mailfolder.DefaultListMailFoldersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderClient.MoveMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

payload := mailfolder.MoveMailFolderRequest{
	// ...
}


read, err := client.MoveMailFolder(ctx, id, payload, mailfolder.DefaultMoveMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.UpdateMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userId", "mailFolderId")

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.UpdateMailFolder(ctx, id, payload, mailfolder.DefaultUpdateMailFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
