
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfoldermessage` Documentation

The `mailfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/mailfoldermessage"
```


### Client Initialization

```go
client := mailfoldermessage.NewMailFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

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


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageCopy`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageCopyRequest{
	// ...
}


read, err := client.CreateMailFolderMessageCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageCreateForward`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageCreateForwardRequest{
	// ...
}


read, err := client.CreateMailFolderMessageCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageCreateReply`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageCreateReplyRequest{
	// ...
}


read, err := client.CreateMailFolderMessageCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageCreateReplyAll`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMailFolderMessageCreateReplyAll(ctx, id, payload)
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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

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


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageMove`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

payload := mailfoldermessage.CreateMailFolderMessageMoveRequest{
	// ...
}


read, err := client.CreateMailFolderMessageMove(ctx, id, payload)
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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

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


### Example Usage: `MailFolderMessageClient.CreateMailFolderMessageSend`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.CreateMailFolderMessageSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.CreateUpdateMailFolderMessageValue`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateMailFolderMessageValue(ctx, id, payload)
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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.DeleteMailFolderMessage(ctx, id)
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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMailFolderMessage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderMessageClient.GetMailFolderMessageCount`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

read, err := client.GetMailFolderMessageCount(ctx, id)
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
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMailFolderMessageValue(ctx, id)
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
id := mailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

// alternatively `client.ListMailFolderMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMailFolderMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderMessageClient.UpdateMailFolderMessage`

```go
ctx := context.TODO()
id := mailfoldermessage.NewMeMailFolderIdMessageID("mailFolderIdValue", "messageIdValue")

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
