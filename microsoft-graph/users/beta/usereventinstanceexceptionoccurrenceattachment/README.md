
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceexceptionoccurrenceattachment` Documentation

The `usereventinstanceexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventinstanceexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usereventinstanceexceptionoccurrenceattachment.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.DeleteUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.DeleteUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value", "attachmentIdValue")

read, err := client.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventInstanceExceptionOccurrenceAttachmentClient.ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

// alternatively `client.ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
