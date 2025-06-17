
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/message` Documentation

The `message` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/message"
```


### Client Initialization

```go
client := message.NewMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageClient.CopyMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.CopyMessageRequest{
	// ...
}


read, err := client.CopyMessage(ctx, id, payload, message.DefaultCopyMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessage`

```go
ctx := context.TODO()
id := message.NewUserID("userId")

payload := message.Message{
	// ...
}


read, err := client.CreateMessage(ctx, id, payload, message.DefaultCreateMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageForward`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.CreateMessageForwardRequest{
	// ...
}


read, err := client.CreateMessageForward(ctx, id, payload, message.DefaultCreateMessageForwardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessagePermanentDelete`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.CreateMessagePermanentDelete(ctx, id, message.DefaultCreateMessagePermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageReply`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.CreateMessageReplyRequest{
	// ...
}


read, err := client.CreateMessageReply(ctx, id, payload, message.DefaultCreateMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageReplyAll`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.CreateMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMessageReplyAll(ctx, id, payload, message.DefaultCreateMessageReplyAllOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageUnsubscribe`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.CreateMessageUnsubscribe(ctx, id, message.DefaultCreateMessageUnsubscribeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.DeleteMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.DeleteMessage(ctx, id, message.DefaultDeleteMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.ForwardMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.ForwardMessageRequest{
	// ...
}


read, err := client.ForwardMessage(ctx, id, payload, message.DefaultForwardMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.GetMessage(ctx, id, message.DefaultGetMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetMessageValue`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.GetMessageValue(ctx, id, message.DefaultGetMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetMessagesCount`

```go
ctx := context.TODO()
id := message.NewUserID("userId")

read, err := client.GetMessagesCount(ctx, id, message.DefaultGetMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.ListMessages`

```go
ctx := context.TODO()
id := message.NewUserID("userId")

// alternatively `client.ListMessages(ctx, id, message.DefaultListMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMessagesComplete(ctx, id, message.DefaultListMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MessageClient.MarkMessageAsJunk`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.MarkMessageAsJunkRequest{
	// ...
}


read, err := client.MarkMessageAsJunk(ctx, id, payload, message.DefaultMarkMessageAsJunkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.MarkMessageAsNotJunk`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.MarkMessageAsNotJunkRequest{
	// ...
}


read, err := client.MarkMessageAsNotJunk(ctx, id, payload, message.DefaultMarkMessageAsNotJunkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.MoveMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.MoveMessageRequest{
	// ...
}


read, err := client.MoveMessage(ctx, id, payload, message.DefaultMoveMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.RemoveMessageValue`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.RemoveMessageValue(ctx, id, message.DefaultRemoveMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.ReplyAllMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.ReplyAllMessageRequest{
	// ...
}


read, err := client.ReplyAllMessage(ctx, id, payload, message.DefaultReplyAllMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.ReplyMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.ReplyMessageRequest{
	// ...
}


read, err := client.ReplyMessage(ctx, id, payload, message.DefaultReplyMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.SendMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

read, err := client.SendMessage(ctx, id, message.DefaultSendMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.SetMessageValue`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")
var payload []byte

read, err := client.SetMessageValue(ctx, id, payload, message.DefaultSetMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.UpdateMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userId", "messageId")

payload := message.Message{
	// ...
}


read, err := client.UpdateMessage(ctx, id, payload, message.DefaultUpdateMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
