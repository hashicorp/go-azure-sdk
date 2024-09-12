
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstance` Documentation

The `calendargroupcalendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.AcceptCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.CancelCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CancelCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.DeclineCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.DismissCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.ForwardCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.ListCalendarGroupCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceInstances(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendargroupcalendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.SnoozeCalendarGroupCalendarViewExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
