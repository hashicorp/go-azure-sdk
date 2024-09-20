
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/message` Documentation

The `message` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/message"
```


### Client Initialization

```go
client := message.NewMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageClient.CopyMessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageId")

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

payload := message.Message{
	// ...
}


read, err := client.CreateMessage(ctx, payload, message.DefaultCreateMessageOperationOptions())
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
id := message.NewMeMessageID("messageId")

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


### Example Usage: `MessageClient.CreateMessageReply`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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


### Example Usage: `MessageClient.DeleteMessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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


read, err := client.GetMessagesCount(ctx, message.DefaultGetMessagesCountOperationOptions())
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


// alternatively `client.ListMessages(ctx, message.DefaultListMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMessagesComplete(ctx, message.DefaultListMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MessageClient.MoveMessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")

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
id := message.NewMeMessageID("messageId")
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
id := message.NewMeMessageID("messageId")

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
