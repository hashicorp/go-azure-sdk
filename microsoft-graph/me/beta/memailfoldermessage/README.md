
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfoldermessage` Documentation

The `memailfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfoldermessage"
```


### Client Initialization

```go
client := memailfoldermessage.NewMeMailFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessage`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

payload := memailfoldermessage.Message{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdCopy`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdCreateForward`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdCreateReply`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdForward`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdMarkAsJunk`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdMarkAsJunkRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdMarkAsJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdMarkAsNotJunk`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdMarkAsNotJunkRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdMarkAsNotJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdMove`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdReply`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdReplyAll`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.CreateMeMailFolderByIdMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdSend`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.CreateMeMailFolderByIdMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateMeMailFolderByIdMessageByIdUnsubscribe`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.CreateMeMailFolderByIdMessageByIdUnsubscribe(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.CreateUpdateMeMailFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateMeMailFolderByIdMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.DeleteMeMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.DeleteMeMailFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.GetMeMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMeMailFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.GetMeMailFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

read, err := client.GetMeMailFolderByIdMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.GetMeMailFolderByIdMessageCount`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

read, err := client.GetMeMailFolderByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderMessageClient.ListMeMailFolderByIdMessages`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderID("mailFolderIdValue")

// alternatively `client.ListMeMailFolderByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMailFolderByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeMailFolderMessageClient.UpdateMeMailFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfoldermessage.NewMeMailFolderMessageID("mailFolderIdValue", "messageIdValue")

payload := memailfoldermessage.Message{
	// ...
}


read, err := client.UpdateMeMailFolderByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
