
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessage` Documentation

The `mailfoldermessage` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessage"
```


### Client Initialization

```go
client := mailfoldermessage.NewMailFolderMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderMessageClient.CopyMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.CopyMailFolderMessageRequest{
	// ...
}


read, err := client.CopyMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultCopyMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderID("userId", "mailFolderId")

payload := mailfoldermessage.Message{
	// ...
}


read, err := client.CreateMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultCreateMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.CreateMailFolderMessageForwardRequest{
	// ...
}


read, err := client.CreateMailFolderMessageForward(ctx, id, payload, mailfoldermessage.DefaultCreateMailFolderMessageForwardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessagePermanentDelete`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

read, err := client.CreateMailFolderMessagePermanentDelete(ctx, id, mailfoldermessage.DefaultCreateMailFolderMessagePermanentDeleteOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.CreateMailFolderMessageReplyRequest{
	// ...
}


read, err := client.CreateMailFolderMessageReply(ctx, id, payload, mailfoldermessage.DefaultCreateMailFolderMessageReplyOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.CreateMailFolderMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderMessageReplyAll(ctx, id, payload, mailfoldermessage.DefaultCreateMailFolderMessageReplyAllOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageUnsubscribe`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

read, err := client.CreateMailFolderMessageUnsubscribe(ctx, id, mailfoldermessage.DefaultCreateMailFolderMessageUnsubscribeOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.ForwardMailFolderMessageRequest{
	// ...
}


read, err := client.ForwardMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultForwardMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

read, err := client.GetMailFolderMessageValue(ctx, id, mailfoldermessage.DefaultGetMailFolderMessageValueOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderID("userId", "mailFolderId")

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
id := mailfoldermessage.NewUserIdMailFolderID("userId", "mailFolderId")

// alternatively `client.ListMailFolderMessages(ctx, id, mailfoldermessage.DefaultListMailFolderMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderMessagesComplete(ctx, id, mailfoldermessage.DefaultListMailFolderMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderMessageClient.MarkMailFolderMessageAsJunk`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.MarkMailFolderMessageAsJunkRequest{
	// ...
}


read, err := client.MarkMailFolderMessageAsJunk(ctx, id, payload, mailfoldermessage.DefaultMarkMailFolderMessageAsJunkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.MarkMailFolderMessageAsNotJunk`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.MarkMailFolderMessageAsNotJunkRequest{
	// ...
}


read, err := client.MarkMailFolderMessageAsNotJunk(ctx, id, payload, mailfoldermessage.DefaultMarkMailFolderMessageAsNotJunkOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.MoveMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.MoveMailFolderMessageRequest{
	// ...
}


read, err := client.MoveMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultMoveMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.ReplyAllMailFolderMessageRequest{
	// ...
}


read, err := client.ReplyAllMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultReplyAllMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.ReplyMailFolderMessageRequest{
	// ...
}


read, err := client.ReplyMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultReplyMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

read, err := client.SendMailFolderMessage(ctx, id, mailfoldermessage.DefaultSendMailFolderMessageOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")
var payload []byte

read, err := client.SetMailFolderMessageValue(ctx, id, payload, mailfoldermessage.DefaultSetMailFolderMessageValueOperationOptions())
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
id := mailfoldermessage.NewUserIdMailFolderIdMessageID("userId", "mailFolderId", "messageId")

payload := mailfoldermessage.Message{
	// ...
}


read, err := client.UpdateMailFolderMessage(ctx, id, payload, mailfoldermessage.DefaultUpdateMailFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
