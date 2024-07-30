
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment` Documentation

The `calendarcalendarviewattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment"
```


### Client Initialization

```go
client := calendarcalendarviewattachment.NewCalendarCalendarViewAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewAttachmentClient.CreateCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarviewattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarCalendarViewAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.CreateCalendarCalendarViewAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarviewattachment.CreateCalendarCalendarViewAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.DeleteCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteCalendarCalendarViewAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.GetCalendarCalendarViewAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewIdAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetCalendarCalendarViewAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.GetCalendarCalendarViewAttachmentCount`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarCalendarViewAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewAttachmentClient.ListCalendarCalendarViewAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewattachment.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListCalendarCalendarViewAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
