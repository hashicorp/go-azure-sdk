
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermessage` Documentation

The `usermessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usermessage"
```


### Client Initialization

```go
client := usermessage.NewUserMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserMessageClient.CreateUpdateUserByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateUserByIdMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessage`

```go
ctx := context.TODO()
id := usermessage.NewUserID("userIdValue")

payload := usermessage.Message{
	// ...
}


read, err := client.CreateUserByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdCopy`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdCreateForward`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdCreateReply`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdForward`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdMarkAsJunk`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdMarkAsJunkRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdMarkAsJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdMarkAsNotJunk`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdMarkAsNotJunkRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdMarkAsNotJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdMove`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdReply`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdReplyAll`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.CreateUserByIdMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateUserByIdMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdSend`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.CreateUserByIdMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.CreateUserByIdMessageByIdUnsubscribe`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.CreateUserByIdMessageByIdUnsubscribe(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.DeleteUserByIdMessageById`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.DeleteUserByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.GetUserByIdMessageById`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.GetUserByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.GetUserByIdMessageByIdValue`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

read, err := client.GetUserByIdMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.GetUserByIdMessageCount`

```go
ctx := context.TODO()
id := usermessage.NewUserID("userIdValue")

read, err := client.GetUserByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserMessageClient.ListUserByIdMessages`

```go
ctx := context.TODO()
id := usermessage.NewUserID("userIdValue")

// alternatively `client.ListUserByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserMessageClient.UpdateUserByIdMessageById`

```go
ctx := context.TODO()
id := usermessage.NewUserMessageID("userIdValue", "messageIdValue")

payload := usermessage.Message{
	// ...
}


read, err := client.UpdateUserByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
