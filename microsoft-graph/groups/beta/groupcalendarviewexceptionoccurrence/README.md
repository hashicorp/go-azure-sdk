
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrence` Documentation

The `groupcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewexceptionoccurrence.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.GetGroupByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewExceptionOccurrenceClient.ListGroupByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
