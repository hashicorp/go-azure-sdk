
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `groupcalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := groupcalendarcalendarviewinstanceexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceExceptionOccurrenceClient.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
