
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


### Example Usage: `MessageClient.CreateMessageCopy`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageCopyRequest{
	// ...
}


read, err := client.CreateMessageCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageCreateForward`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageCreateForwardRequest{
	// ...
}


read, err := client.CreateMessageCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageCreateReply`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageCreateReplyRequest{
	// ...
}


read, err := client.CreateMessageCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMessageCreateReplyAll`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMessageCreateReplyAll(ctx, id, payload)
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


### Example Usage: `MessageClient.CreateMessageMove`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

payload := message.CreateMessageMoveRequest{
	// ...
}


read, err := client.CreateMessageMove(ctx, id, payload)
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


### Example Usage: `MessageClient.CreateMessageSend`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")

read, err := client.CreateMessageSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateUpdateMessageValue`

```go
ctx := context.TODO()
id := message.NewUserIdMessageID("userIdValue", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateMessageValue(ctx, id, payload)
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

read, err := client.DeleteMessage(ctx, id)
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

read, err := client.GetMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetMessageCount`

```go
ctx := context.TODO()
id := message.NewUserID("userIdValue")

read, err := client.GetMessageCount(ctx, id)
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


### Example Usage: `MessageClient.ListMessages`

```go
ctx := context.TODO()
id := message.NewUserID("userIdValue")

// alternatively `client.ListMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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
