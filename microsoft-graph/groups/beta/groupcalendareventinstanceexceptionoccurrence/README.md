
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventinstanceexceptionoccurrence` Documentation

The `groupcalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendareventinstanceexceptionoccurrence.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.GetGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.GetGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceExceptionOccurrenceClient.ListGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
