
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memessage` Documentation

The `memessage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memessage"
```


### Client Initialization

```go
client := memessage.NewMeMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMessageClient.CreateMeMessage`

```go
ctx := context.TODO()

payload := memessage.Message{
	// ...
}


read, err := client.CreateMeMessage(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdCopy`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdCopyRequest{
	// ...
}


read, err := client.CreateMeMessageByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdCreateForward`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdCreateForwardRequest{
	// ...
}


read, err := client.CreateMeMessageByIdCreateForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdCreateReply`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdCreateReplyRequest{
	// ...
}


read, err := client.CreateMeMessageByIdCreateReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdCreateReplyAll`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdCreateReplyAllRequest{
	// ...
}


read, err := client.CreateMeMessageByIdCreateReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdForward`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdForwardRequest{
	// ...
}


read, err := client.CreateMeMessageByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdMarkAsJunk`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdMarkAsJunkRequest{
	// ...
}


read, err := client.CreateMeMessageByIdMarkAsJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdMarkAsNotJunk`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdMarkAsNotJunkRequest{
	// ...
}


read, err := client.CreateMeMessageByIdMarkAsNotJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdMove`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdMoveRequest{
	// ...
}


read, err := client.CreateMeMessageByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdReply`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdReplyRequest{
	// ...
}


read, err := client.CreateMeMessageByIdReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdReplyAll`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.CreateMeMessageByIdReplyAllRequest{
	// ...
}


read, err := client.CreateMeMessageByIdReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdSend`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

read, err := client.CreateMeMessageByIdSend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateMeMessageByIdUnsubscribe`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

read, err := client.CreateMeMessageByIdUnsubscribe(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.CreateUpdateMeMessageByIdValue`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")
var payload []byte

read, err := client.CreateUpdateMeMessageByIdValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.DeleteMeMessageById`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

read, err := client.DeleteMeMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.GetMeMessageById`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

read, err := client.GetMeMessageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.GetMeMessageByIdValue`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

read, err := client.GetMeMessageByIdValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.GetMeMessageCount`

```go
ctx := context.TODO()


read, err := client.GetMeMessageCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMessageClient.ListMeMessages`

```go
ctx := context.TODO()


// alternatively `client.ListMeMessages(ctx)` can be used to do batched pagination
items, err := client.ListMeMessagesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeMessageClient.UpdateMeMessageById`

```go
ctx := context.TODO()
id := memessage.NewMeMessageID("messageIdValue")

payload := memessage.Message{
	// ...
}


read, err := client.UpdateMeMessageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
