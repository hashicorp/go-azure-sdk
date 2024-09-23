
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessage` Documentation

The `mailfolderchildfoldermessage` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessage"
```


### Client Initialization

```go
client := mailfolderchildfoldermessage.NewMailFolderChildFolderMessageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderMessageClient.CopyMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.CopyMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.CopyMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultCopyMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

payload := mailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultCreateMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageForward`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageForwardRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageForward(ctx, id, payload, mailfolderchildfoldermessage.DefaultCreateMailFolderChildFolderMessageForwardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageReply`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageReplyRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageReply(ctx, id, payload, mailfolderchildfoldermessage.DefaultCreateMailFolderChildFolderMessageReplyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageReplyAll`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageReplyAll(ctx, id, payload, mailfolderchildfoldermessage.DefaultCreateMailFolderChildFolderMessageReplyAllOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.DeleteMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

read, err := client.DeleteMailFolderChildFolderMessage(ctx, id, mailfolderchildfoldermessage.DefaultDeleteMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.ForwardMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.ForwardMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ForwardMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultForwardMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.GetMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

read, err := client.GetMailFolderChildFolderMessage(ctx, id, mailfolderchildfoldermessage.DefaultGetMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.GetMailFolderChildFolderMessageValue`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

read, err := client.GetMailFolderChildFolderMessageValue(ctx, id, mailfolderchildfoldermessage.DefaultGetMailFolderChildFolderMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.GetMailFolderChildFolderMessagesCount`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

read, err := client.GetMailFolderChildFolderMessagesCount(ctx, id, mailfolderchildfoldermessage.DefaultGetMailFolderChildFolderMessagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.ListMailFolderChildFolderMessages`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userId", "mailFolderId", "mailFolderId1")

// alternatively `client.ListMailFolderChildFolderMessages(ctx, id, mailfolderchildfoldermessage.DefaultListMailFolderChildFolderMessagesOperationOptions())` can be used to do batched pagination
items, err := client.ListMailFolderChildFolderMessagesComplete(ctx, id, mailfolderchildfoldermessage.DefaultListMailFolderChildFolderMessagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderChildFolderMessageClient.MoveMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.MoveMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.MoveMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultMoveMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.RemoveMailFolderChildFolderMessageValue`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

read, err := client.RemoveMailFolderChildFolderMessageValue(ctx, id, mailfolderchildfoldermessage.DefaultRemoveMailFolderChildFolderMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.ReplyAllMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.ReplyAllMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ReplyAllMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultReplyAllMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.ReplyMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.ReplyMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ReplyMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultReplyMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.SendMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

read, err := client.SendMailFolderChildFolderMessage(ctx, id, mailfolderchildfoldermessage.DefaultSendMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.SetMailFolderChildFolderMessageValue`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")
var payload []byte

read, err := client.SetMailFolderChildFolderMessageValue(ctx, id, payload, mailfolderchildfoldermessage.DefaultSetMailFolderChildFolderMessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.UpdateMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

payload := mailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.UpdateMailFolderChildFolderMessage(ctx, id, payload, mailfolderchildfoldermessage.DefaultUpdateMailFolderChildFolderMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
