
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolderchildfoldermessage` Documentation

The `memailfolderchildfoldermessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolderchildfoldermessage"
```


### Client Initialization

```go
client := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessage`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

payload := memailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdCopy`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateForward`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReply`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdForward`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdMove`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdReply`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdReplyAll`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.CreateMeMailFolderByIdChildFolderByIdMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateMeMailFolderByIdChildFolderByIdMessageByIdSend`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.CreateMeMailFolderByIdChildFolderByIdMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.CreateUpdateMeMailFolderByIdChildFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")
var payload []byte

read, err := client.CreateUpdateMeMailFolderByIdChildFolderByIdMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.DeleteMeMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.DeleteMeMailFolderByIdChildFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.GetMeMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMeMailFolderByIdChildFolderByIdMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.GetMeMailFolderByIdChildFolderByIdMessageByIdValue`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

read, err := client.GetMeMailFolderByIdChildFolderByIdMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.GetMeMailFolderByIdChildFolderByIdMessageCount`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

read, err := client.GetMeMailFolderByIdChildFolderByIdMessageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.ListMeMailFolderByIdChildFolderByIdMessages`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

// alternatively `client.ListMeMailFolderByIdChildFolderByIdMessages(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMailFolderByIdChildFolderByIdMessagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeMailFolderChildFolderMessageClient.UpdateMeMailFolderByIdChildFolderByIdMessageById`

```go
ctx := context.TODO()
id := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageID("mailFolderIdValue", "mailFolderId1Value", "messageIdValue")

payload := memailfolderchildfoldermessage.Message{
	// ...
}


read, err := client.UpdateMeMailFolderByIdChildFolderByIdMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
