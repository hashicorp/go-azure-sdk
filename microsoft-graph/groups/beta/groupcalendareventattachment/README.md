
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventattachment` Documentation

The `groupcalendareventattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventattachment"
```


### Client Initialization

```go
client := groupcalendareventattachment.NewGroupCalendarEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventAttachmentClient.CreateGroupByIdCalendarEventByIdAttachment`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendareventattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventAttachmentClient.CreateGroupByIdCalendarEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendareventattachment.CreateGroupByIdCalendarEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventAttachmentClient.DeleteGroupByIdCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventAttachmentClient.GetGroupByIdCalendarEventByIdAttachmentById`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdCalendarEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventAttachmentClient.GetGroupByIdCalendarEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventAttachmentClient.ListGroupByIdCalendarEventByIdAttachments`

```go
ctx := context.TODO()
id := groupcalendareventattachment.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
