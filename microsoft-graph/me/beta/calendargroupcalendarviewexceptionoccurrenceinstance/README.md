
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewexceptionoccurrenceinstance` Documentation

The `calendargroupcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendargroupcalendarviewexceptionoccurrenceinstance.NewCalendarGroupCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.AcceptCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.CancelCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.CancelCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.DeclineCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.DismissCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.ForwardCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstance(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.ListCalendarGroupCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarViewExceptionOccurrenceInstances(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendargroupcalendarviewexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarGroupCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.SnoozeCalendarGroupCalendarViewExceptionOccurrenceInstanceReminderRequest{
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


### Example Usage: `CalendarGroupCalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarGroupCalendarViewExceptionOccurrenceInstanceRequest{
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
