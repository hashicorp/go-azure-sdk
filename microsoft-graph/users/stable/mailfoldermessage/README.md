
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessage` Documentation

The `mailfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessage"
```


### Client Initialization

```go
client := mailfoldermessage.NewMailFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderMessageClient.CopyMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CopyMailFolderMessageRequest{
	// ...
}


read, err := client.CopyMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

payload := mailfoldermessage.Message{
	// ...
}


read, err := client.CreateMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageForward`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageForwardRequest{
	// ...
}


read, err := client.CreateMailFolderMessageForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageReply`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageReplyRequest{
	// ...
}


read, err := client.CreateMailFolderMessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageReplyAll`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderMessageReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.DeleteMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.DeleteMailFolderMessage(ctx, id, mailfoldermessage.DefaultDeleteMailFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.ForwardMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.ForwardMailFolderMessageRequest{
	// ...
}


read, err := client.ForwardMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.GetMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.GetMailFolderMessage(ctx, id, mailfoldermessage.DefaultGetMailFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.GetMailFolderMessageValue`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.GetMailFolderMessageValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.GetMailFolderMessagesCount`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

read, err := client.GetMailFolderMessagesCount(ctx, id, mailfoldermessage.DefaultGetMailFolderMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.ListMailFolderMessages`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderID("userIdValue", "mailFolderIdValue")

// alternatively `client.ListMailFolderMessages(ctx, id, mailfoldermessage.DefaultListMailFolderMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderMessagesComplete(ctx, id, mailfoldermessage.DefaultListMailFolderMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderMessageClient.MoveMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.MoveMailFolderMessageRequest{
	// ...
}


read, err := client.MoveMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.RemoveMailFolderMessageValue`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.RemoveMailFolderMessageValue(ctx, id, mailfoldermessage.DefaultRemoveMailFolderMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.ReplyAllMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.ReplyAllMailFolderMessageRequest{
	// ...
}


read, err := client.ReplyAllMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.ReplyMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.ReplyMailFolderMessageRequest{
	// ...
}


read, err := client.ReplyMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.SendMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

read, err := client.SendMailFolderMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.SetMailFolderMessageValue`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")
var payload []byte

read, err := client.SetMailFolderMessageValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.UpdateMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userIdValue", "mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.Message{
	// ...
}


read, err := client.UpdateMailFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
