
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfolder` Documentation

The `usermailfolder` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfolder"
```


### Client Initialization

```go
client := usermailfolder.NewUserMailFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderClient.CreateUserByIdMailFolder`

```go
ctx := context.TODO()
id := usermailfolder.NewUserID("userIdValue")

payload := usermailfolder.MailFolder{
	// ...
}


read, err := client.CreateUserByIdMailFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.CreateUserByIdMailFolderByIdCopy`

```go
ctx := context.TODO()
id := usermailfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

payload := usermailfolder.CreateUserByIdMailFolderByIdCopyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.CreateUserByIdMailFolderByIdMove`

```go
ctx := context.TODO()
id := usermailfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

payload := usermailfolder.CreateUserByIdMailFolderByIdMoveRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.DeleteUserByIdMailFolderById`

```go
ctx := context.TODO()
id := usermailfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.DeleteUserByIdMailFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.GetUserByIdMailFolderById`

```go
ctx := context.TODO()
id := usermailfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.GetUserByIdMailFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.GetUserByIdMailFolderCount`

```go
ctx := context.TODO()
id := usermailfolder.NewUserID("userIdValue")

read, err := client.GetUserByIdMailFolderCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderClient.ListUserByIdMailFolders`

```go
ctx := context.TODO()
id := usermailfolder.NewUserID("userIdValue")

// alternatively `client.ListUserByIdMailFolders(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFoldersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserMailFolderClient.UpdateUserByIdMailFolderById`

```go
ctx := context.TODO()
id := usermailfolder.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

payload := usermailfolder.MailFolder{
	// ...
}


read, err := client.UpdateUserByIdMailFolderById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
