
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceattachment` Documentation

The `calendarcalendarviewinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceattachment"
```


### Client Initialization

```go
client := calendarcalendarviewinstanceattachment.NewCalendarCalendarViewInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.CreateCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.CreateCalendarCalendarViewInstanceAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstanceattachment.CreateCalendarCalendarViewInstanceAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.DeleteCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteCalendarCalendarViewInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.GetCalendarCalendarViewInstanceAttachment`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceIdAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetCalendarCalendarViewInstanceAttachment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.GetCalendarCalendarViewInstanceAttachmentCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewInstanceAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceAttachmentClient.ListCalendarCalendarViewInstanceAttachments`

```go
ctx := context.TODO()
id := calendarcalendarviewinstanceattachment.NewGroupIdCalendarCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarCalendarViewInstanceAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewInstanceAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
