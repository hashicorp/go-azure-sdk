
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrenceattachment` Documentation

The `usereventexceptionoccurrenceattachment` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usereventexceptionoccurrenceattachment"
```


### Client Initialization

```go
client := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.CreateUserByIdEventByIdExceptionOccurrenceByIdAttachment`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrenceattachment.Attachment{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.CreateUserByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usereventexceptionoccurrenceattachment.CreateUserByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateUserByIdEventByIdExceptionOccurrenceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.DeleteUserByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteUserByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.GetUserByIdEventByIdExceptionOccurrenceByIdAttachmentById`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceAttachmentID("userIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.GetUserByIdEventByIdExceptionOccurrenceByIdAttachmentCount`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdEventByIdExceptionOccurrenceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserEventExceptionOccurrenceAttachmentClient.ListUserByIdEventByIdExceptionOccurrenceByIdAttachments`

```go
ctx := context.TODO()
id := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdEventByIdExceptionOccurrenceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdEventByIdExceptionOccurrenceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
