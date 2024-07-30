
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrence` Documentation

The `calendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendarviewinstanceexceptionoccurrence.NewCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceAccept`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceCancel`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceCancelRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceDecline`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceDismissReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateCalendarViewInstanceExceptionOccurrenceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceForward`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceForwardRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CreateCalendarViewInstanceExceptionOccurrenceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CreateCalendarViewInstanceExceptionOccurrenceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceExceptionOccurrenceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewInstanceExceptionOccurrence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstanceExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.ListCalendarViewInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewInstanceExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
