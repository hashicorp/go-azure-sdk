
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfolderchildfolder` Documentation

The `usermailfolderchildfolder` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfolderchildfolder"
```


### Client Initialization

```go
client := usermailfolderchildfolder.NewUserMailFolderChildFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderChildFolderClient.CreateUserByIdMailFolderByIdChildFolder`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

payload := usermailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.CreateUserByIdMailFolderByIdChildFolderByIdCopy`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := usermailfolderchildfolder.CreateUserByIdMailFolderByIdChildFolderByIdCopyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.CreateUserByIdMailFolderByIdChildFolderByIdMove`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := usermailfolderchildfolder.CreateUserByIdMailFolderByIdChildFolderByIdMoveRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.DeleteUserByIdMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

read, err := client.DeleteUserByIdMailFolderByIdChildFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.GetUserByIdMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

read, err := client.GetUserByIdMailFolderByIdChildFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.GetUserByIdMailFolderByIdChildFolderCount`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.GetUserByIdMailFolderByIdChildFolderCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderClient.ListUserByIdMailFolderByIdChildFolders`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

// alternatively `client.ListUserByIdMailFolderByIdChildFolders(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFolderByIdChildFoldersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserMailFolderChildFolderClient.UpdateUserByIdMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := usermailfolderchildfolder.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := usermailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.UpdateUserByIdMailFolderByIdChildFolderById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
