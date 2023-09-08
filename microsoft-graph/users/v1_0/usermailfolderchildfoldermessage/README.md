
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfolderchildfoldermessage` Documentation

The `usermailfolderchildfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usermailfolderchildfoldermessage"
```


### Client Initialization

```go
client := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUpdateUserByIdMailFolderByIdChildFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateUserByIdMailFolderByIdChildFolderByIdMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessage`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

payload := usermailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCopy`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateForward`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReply`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdForward`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdMove`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReply`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReplyAll`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdSend`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.CreateUserByIdMailFolderByIdChildFolderByIdMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.DeleteUserByIdMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.DeleteUserByIdMailFolderByIdChildFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.GetUserByIdMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdChildFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdChildFolderByIdMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.GetUserByIdMailFolderByIdChildFolderByIdMessageCount`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

read, err := client.GetUserByIdMailFolderByIdChildFolderByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.ListUserByIdMailFolderByIdChildFolderByIdMessages`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

// alternatively `client.ListUserByIdMailFolderByIdChildFolderByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFolderByIdChildFolderByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserMailFolderChildFolderMessageClient.UpdateUserByIdMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := usermailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.UpdateUserByIdMailFolderByIdChildFolderByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
