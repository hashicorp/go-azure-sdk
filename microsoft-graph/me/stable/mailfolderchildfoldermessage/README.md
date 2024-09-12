
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessage` Documentation

The `mailfolderchildfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfolderchildfoldermessage"
```


### Client Initialization

```go
client := mailfolderchildfoldermessage.NewMailFolderChildFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderMessageClient.CopyMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CopyMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.CopyMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderID("mailFolderIdValue", "mailFolderId1Value")

payload := mailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageForwardRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageForward(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageReplyRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageReply(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageReplyAll(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.ForwardMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ForwardMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMailFolderChildFolderMessageValue(ctx, id)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderID("mailFolderIdValue", "mailFolderId1Value")

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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderID("mailFolderIdValue", "mailFolderId1Value")

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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.MoveMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.MoveMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.ReplyAllMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ReplyAllMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.ReplyMailFolderChildFolderMessageRequest{
	// ...
}


read, err := client.ReplyMailFolderChildFolderMessage(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.SendMailFolderChildFolderMessage(ctx, id)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")
var payload []byte

read, err := client.SetMailFolderChildFolderMessageValue(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewMeMailFolderIdChildFolderIdMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.UpdateMailFolderChildFolderMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
