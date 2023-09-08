
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendareventinstanceattachment` Documentation

The `groupcalendareventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendareventinstanceattachment"
```


### Client Initialization

```go
client := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.CreateGroupByIdCalendarEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.CreateGroupByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstanceattachment.CreateGroupByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.DeleteGroupByIdCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.GetGroupByIdCalendarEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.GetGroupByIdCalendarEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceAttachmentClient.ListGroupByIdCalendarEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
