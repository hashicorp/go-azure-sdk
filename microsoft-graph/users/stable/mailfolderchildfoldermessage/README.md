
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessage` Documentation

The `mailfolderchildfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessage"
```


### Client Initialization

```go
client := mailfolderchildfoldermessage.NewMailFolderChildFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

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


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageCopy`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageCopyRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageCreateForward`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageCreateForwardRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageCreateReply`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageCreateReplyRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageCreateReplyAll`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageCreateReplyAll(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageMove`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := mailfolderchildfoldermessage.CreateMailFolderChildFolderMessageMoveRequest{
	// ...
}


read, err := client.CreateMailFolderChildFolderMessageMove(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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


### Example Usage: `MailFolderChildFolderMessageClient.CreateMailFolderChildFolderMessageSend`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.CreateMailFolderChildFolderMessageSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.CreateUpdateMailFolderChildFolderMessageValue`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateMailFolderChildFolderMessageValue(ctx, id, payload)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.DeleteMailFolderChildFolderMessage(ctx, id)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMailFolderChildFolderMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderChildFolderMessageClient.GetMailFolderChildFolderMessageCount`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

read, err := client.GetMailFolderChildFolderMessageCount(ctx, id)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMailFolderChildFolderMessageValue(ctx, id)
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
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderID("userIdValue", "mailFolderIdValue", "mailFolderId1Value")

// alternatively `client.ListMailFolderChildFolderMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMailFolderChildFolderMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderChildFolderMessageClient.UpdateMailFolderChildFolderMessage`

```go
ctx := context.TODO()
id := mailfolderchildfoldermessage.NewUserIdMailFolderIdChildFolderIdMessageID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

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
