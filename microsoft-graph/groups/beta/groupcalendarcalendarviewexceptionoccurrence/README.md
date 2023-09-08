
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrence` Documentation

The `groupcalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewexceptionoccurrence.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewExceptionOccurrenceClient.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
