
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewattachment` Documentation

The `groupcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewattachment"
```


### Client Initialization

```go
client := groupcalendarviewattachment.NewGroupCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewAttachmentClient.CreateGroupByIdCalendarViewByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewAttachmentClient.CreateGroupByIdCalendarViewByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarviewattachment.CreateGroupByIdCalendarViewByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewAttachmentClient.DeleteGroupByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewAttachmentClient.GetGroupByIdCalendarViewByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarViewByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewAttachmentClient.GetGroupByIdCalendarViewByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarViewByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewAttachmentClient.ListGroupByIdCalendarViewByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendarviewattachment.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarViewByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
