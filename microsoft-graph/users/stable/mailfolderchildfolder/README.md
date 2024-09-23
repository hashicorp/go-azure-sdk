
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfolder` Documentation

The `mailfolderchildfolder` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfolder"
```


### Client Initialization

```go
client := mailfolderchildfolder.NewMailFolderChildFolderClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderClient.CopyMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

payload := mailfolderchildfolder.CopyMailFolderChildFolderRequest{
	// ...
}


read, err := client.CopyMailFolderChildFolder(ctx, id, payload, mailfolderchildfolder.DefaultCopyMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.CreateMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderID("userId", "mailFolderId")

payload := mailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.CreateMailFolderChildFolder(ctx, id, payload, mailfolderchildfolder.DefaultCreateMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.DeleteMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

read, err := client.DeleteMailFolderChildFolder(ctx, id, mailfolderchildfolder.DefaultDeleteMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.GetMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

read, err := client.GetMailFolderChildFolder(ctx, id, mailfolderchildfolder.DefaultGetMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.GetMailFolderChildFoldersCount`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderID("userId", "mailFolderId")

read, err := client.GetMailFolderChildFoldersCount(ctx, id, mailfolderchildfolder.DefaultGetMailFolderChildFoldersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.ListMailFolderChildFolders`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderID("userId", "mailFolderId")

// alternatively `client.ListMailFolderChildFolders(ctx, id, mailfolderchildfolder.DefaultListMailFolderChildFoldersOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderChildFoldersComplete(ctx, id, mailfolderchildfolder.DefaultListMailFolderChildFoldersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderChildFolderClient.MoveMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

payload := mailfolderchildfolder.MoveMailFolderChildFolderRequest{
	// ...
}


read, err := client.MoveMailFolderChildFolder(ctx, id, payload, mailfolderchildfolder.DefaultMoveMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.UpdateMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

payload := mailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.UpdateMailFolderChildFolder(ctx, id, payload, mailfolderchildfolder.DefaultUpdateMailFolderChildFolderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
