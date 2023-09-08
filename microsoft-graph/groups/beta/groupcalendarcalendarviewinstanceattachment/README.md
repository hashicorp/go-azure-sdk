
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceattachment` Documentation

The `groupcalendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstanceattachment.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.DeleteGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceAttachmentClient.ListGroupByIdCalendarCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
