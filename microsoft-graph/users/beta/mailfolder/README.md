
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolder` Documentation

The `mailfolder` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolder"
```


### Client Initialization

```go
client := mailfolder.NewMailFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderClient.CopyMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

payload := mailfolder.CopyMailFolderRequest{
	// ...
}


read, err := client.CopyMailFolder(ctx, id, payload)
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
id := mailfolder.NewUserID("userIdValue")

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.CreateMailFolder(ctx, id, payload)
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
id := mailfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

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
id := mailfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

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
id := mailfolder.NewUserID("userIdValue")

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
id := mailfolder.NewUserID("userIdValue")

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
id := mailfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

payload := mailfolder.MoveMailFolderRequest{
	// ...
}


read, err := client.MoveMailFolder(ctx, id, payload)
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
id := mailfolder.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.UpdateMailFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
