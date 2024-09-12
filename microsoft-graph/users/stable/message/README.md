
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/message` Documentation

The `message` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/message"
```


### Client Initialization

```go
client := message.NewMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageClient.CopyMessage`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CopyMessageRequest{
	// ...
}


read, err := client.CopyMessage(ctx, id, payload)
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
id := message.NewUserID("userIdValue")

payload := message.Message{
	// ...
}


read, err := client.CreateMessage(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageForwardRequest{
	// ...
}


read, err := client.CreateMessageForward(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageReplyRequest{
	// ...
}


read, err := client.CreateMessageReply(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMessageReplyAll(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.ForwardMessageRequest{
	// ...
}


read, err := client.ForwardMessage(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

read, err := client.GetMessageValue(ctx, id)
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
id := message.NewUserID("userIdValue")

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
id := message.NewUserID("userIdValue")

// alternatively `client.ListMessages(ctx, id, message.DefaultListMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMessagesComplete(ctx, id, message.DefaultListMessagesOperationOptions())
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.MoveMessageRequest{
	// ...
}


read, err := client.MoveMessage(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.ReplyAllMessageRequest{
	// ...
}


read, err := client.ReplyAllMessage(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.ReplyMessageRequest{
	// ...
}


read, err := client.ReplyMessage(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

read, err := client.SendMessage(ctx, id)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")
var payload []byte

read, err := client.SetMessageValue(ctx, id, payload)
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
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.Message{
	// ...
}


read, err := client.UpdateMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
