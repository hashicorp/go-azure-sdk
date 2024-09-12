
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfolder` Documentation

The `mailfolderchildfolder` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfolder"
```


### Client Initialization

```go
client := mailfolderchildfolder.NewMailFolderChildFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderClient.CopyMailFolderChildFolder`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfolder.CopyMailFolderChildFolderRequest{
	// ...
}


read, err := client.CopyMailFolderChildFolder(ctx, id, payload)
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
id := mailfolderchildfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

payload := mailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.CreateMailFolderChildFolder(ctx, id, payload)
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
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

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
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

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
id := mailfolderchildfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

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
id := mailfolderchildfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

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
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfolder.MoveMailFolderChildFolderRequest{
	// ...
}


read, err := client.MoveMailFolderChildFolder(ctx, id, payload)
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
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.UpdateMailFolderChildFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
