
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfolder` Documentation

The `mailfolderchildfolder` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfolder"
```


### Client Initialization

```go
client := mailfolderchildfolder.NewMailFolderChildFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
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


### Example Usage: `MailFolderChildFolderClient.CreateMailFolderChildFolderCopy`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfolder.CreateMailFolderChildFolderCopyRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.CreateMailFolderChildFolderMove`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfolder.CreateMailFolderChildFolderMoveRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMove(ctx, id, payload)
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

read, err := client.DeleteMailFolderChildFolder(ctx, id)
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

read, err := client.GetMailFolderChildFolder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderClient.GetMailFolderChildFolderCount`

```go
ctx := context.TODO()
id := mailfolderchildfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.GetMailFolderChildFolderCount(ctx, id)
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

// alternatively `client.ListMailFolderChildFolders(ctx, id)` can be used to do batched pagination
items, err := client.ListMailFolderChildFoldersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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
