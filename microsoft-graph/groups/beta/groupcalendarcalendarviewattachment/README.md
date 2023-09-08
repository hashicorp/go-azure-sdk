
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewattachment` Documentation

The `groupcalendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewattachment"
```


### Client Initialization

```go
client := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.CreateGroupByIdCalendarCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarviewattachment.CreateGroupByIdCalendarCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.DeleteGroupByIdCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.GetGroupByIdCalendarCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.GetGroupByIdCalendarCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewAttachmentClient.ListGroupByIdCalendarCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
