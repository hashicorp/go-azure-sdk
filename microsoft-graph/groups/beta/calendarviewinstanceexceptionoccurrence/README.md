
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


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.AcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.AcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.CancelCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.CancelCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.DeclineCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.DeclineCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.DismissCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarViewInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.ForwardCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.ForwardCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
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

read, err := client.GetCalendarViewInstanceExceptionOccurrence(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.GetCalendarViewInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstanceExceptionOccurrencesCount(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultGetCalendarViewInstanceExceptionOccurrencesCountOperationOptions())
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

// alternatively `client.ListCalendarViewInstanceExceptionOccurrences(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstanceExceptionOccurrencesComplete(ctx, id, calendarviewinstanceexceptionoccurrence.DefaultListCalendarViewInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.SnoozeCalendarViewInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.SnoozeCalendarViewInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewInstanceExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarViewInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendarviewinstanceexceptionoccurrence.NewGroupIdCalendarViewIdInstanceIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewinstanceexceptionoccurrence.TentativelyAcceptCalendarViewInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
