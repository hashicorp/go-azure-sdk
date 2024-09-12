
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/message` Documentation

The `message` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/message"
```


### Client Initialization

```go
client := message.NewMessageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MessageClient.CreateCopyssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateCopyssageRequest{
	// ...
}


read, err := client.CreateCopyssage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateCreatessageForward`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateCreatessageForwardRequest{
	// ...
}


read, err := client.CreateCreatessageForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateCreatessageReply`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateCreatessageReplyRequest{
	// ...
}


read, err := client.CreateCreatessageReply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateCreatessageReplyAll`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateCreatessageReplyAllRequest{
	// ...
}


read, err := client.CreateCreatessageReplyAll(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateForwardssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateForwardssageRequest{
	// ...
}


read, err := client.CreateForwardssage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMarkssageAsJunk`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateMarkssageAsJunkRequest{
	// ...
}


read, err := client.CreateMarkssageAsJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMarkssageAsNotJunk`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateMarkssageAsNotJunkRequest{
	// ...
}


read, err := client.CreateMarkssageAsNotJunk(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateMovessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateMovessageRequest{
	// ...
}


read, err := client.CreateMovessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateReplyssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.CreateReplyssageRequest{
	// ...
}


read, err := client.CreateReplyssage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreateSendssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.CreateSendssage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.Createssage`

```go
ctx := context.TODO()

payload := message.Message{
	// ...
}


read, err := client.Createssage(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.CreatessageUnsubscribe`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.CreatessageUnsubscribe(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.Deletessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.Deletessage(ctx, id, message.DefaultDeletessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.Getssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.Getssage(ctx, id, message.DefaultGetssageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetssageValue`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.GetssageValue(ctx, id, message.DefaultGetssageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.GetssagesCount`

```go
ctx := context.TODO()


read, err := client.GetssagesCount(ctx, message.DefaultGetssagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.Listssages`

```go
ctx := context.TODO()


// alternatively `client.Listssages(ctx, message.DefaultListssagesOperationOptions())` can be used to do batched pagination
items, err := client.ListssagesComplete(ctx, message.DefaultListssagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MessageClient.RemovessageValue`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

read, err := client.RemovessageValue(ctx, id, message.DefaultRemovessageValueOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.ReplyAllssage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.ReplyAllssageRequest{
	// ...
}


read, err := client.ReplyAllssage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.SetssageValue`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")
var payload []byte

read, err := client.SetssageValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MessageClient.Updatessage`

```go
ctx := context.TODO()
id := message.NewMeMessageID("messageIdValue")

payload := message.Message{
	// ...
}


read, err := client.Updatessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
