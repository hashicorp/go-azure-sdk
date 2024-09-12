
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventinstanceexceptionoccurrence` Documentation

The `calendargroupcalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := calendargroupcalendareventinstanceexceptionoccurrence.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.AcceptCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.AcceptCalendarGroupCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CancelCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.CancelCalendarGroupCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.DeclineCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.DeclineCalendarGroupCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.DismissCalendarGroupCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarEventInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.ForwardCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.ForwardCalendarGroupCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, calendargroupcalendareventinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetCalendarGroupCalendarEventInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventInstanceExceptionOccurrencesCount(ctx, id, calendargroupcalendareventinstanceexceptionoccurrence.DefaultGetCalendarGroupCalendarEventInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.ListCalendarGroupCalendarEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventInstanceExceptionOccurrences(ctx, id, calendargroupcalendareventinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstanceExceptionOccurrencesComplete(ctx, id, calendargroupcalendareventinstanceexceptionoccurrence.DefaultListCalendarGroupCalendarEventInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.SnoozeCalendarGroupCalendarEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.SnoozeCalendarGroupCalendarEventInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarEventInstanceExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceExceptionOccurrenceClient.TentativelyAcceptCalendarGroupCalendarEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventinstanceexceptionoccurrence.TentativelyAcceptCalendarGroupCalendarEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
