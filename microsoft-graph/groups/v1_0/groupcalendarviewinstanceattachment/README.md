
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarviewinstanceattachment` Documentation

The `groupcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.CreateGroupByIdCalendarViewByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.CreateGroupByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstanceattachment.CreateGroupByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.DeleteGroupByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.GetGroupByIdCalendarViewByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarViewByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.GetGroupByIdCalendarViewByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarViewByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceAttachmentClient.ListGroupByIdCalendarViewByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
