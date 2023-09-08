
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfoldermessage` Documentation

The `usermailfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermailfoldermessage"
```


### Client Initialization

```go
client := usermailfoldermessage.NewUserMailFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMailFolderMessageClient.CreateUpdateUserByIdMailFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateUserByIdMailFolderByIdMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessage`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

payload := usermailfoldermessage.Message{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdCopy`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdCreateForward`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdCreateReply`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdForward`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdMarkAsJunk`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdMarkAsJunkRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdMarkAsJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdMarkAsNotJunk`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdMarkAsNotJunkRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdMarkAsNotJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdMove`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdReply`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdReplyAll`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.CreateUserByIdMailFolderByIdMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMailFolderByIdMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdSend`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.CreateUserByIdMailFolderByIdMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.CreateUserByIdMailFolderByIdMessageByIdUnsubscribe`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.CreateUserByIdMailFolderByIdMessageByIdUnsubscribe(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.DeleteUserByIdMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.DeleteUserByIdMailFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.GetUserByIdMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.GetUserByIdMailFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.GetUserByIdMailFolderByIdMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.GetUserByIdMailFolderByIdMessageCount`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.GetUserByIdMailFolderByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMailFolderMessageClient.ListUserByIdMailFolderByIdMessages`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderID("userIdValue", "mailFolderIdValue")

// alternatively `client.ListUserByIdMailFolderByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMailFolderByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserMailFolderMessageClient.UpdateUserByIdMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := usermailfoldermessage.NewUserMailFolderMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := usermailfoldermessage.Message{
	// ...
}


read, err := client.UpdateUserByIdMailFolderByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
