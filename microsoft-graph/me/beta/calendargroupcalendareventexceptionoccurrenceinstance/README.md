
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventexceptionoccurrenceinstance` Documentation

The `calendargroupcalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendargroupcalendareventexceptionoccurrenceinstance.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.AcceptCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.AcceptCalendarGroupCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CancelCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.CancelCalendarGroupCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.DeclineCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.DeclineCalendarGroupCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.DismissCalendarGroupCalendarEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarGroupCalendarEventExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.ForwardCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.ForwardCalendarGroupCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, calendargroupcalendareventexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetCalendarGroupCalendarEventExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventExceptionOccurrenceInstancesCount(ctx, id, calendargroupcalendareventexceptionoccurrenceinstance.DefaultGetCalendarGroupCalendarEventExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.ListCalendarGroupCalendarEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarGroupCalendarEventExceptionOccurrenceInstances(ctx, id, calendargroupcalendareventexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarEventExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventExceptionOccurrenceInstancesComplete(ctx, id, calendargroupcalendareventexceptionoccurrenceinstance.DefaultListCalendarGroupCalendarEventExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.SnoozeCalendarGroupCalendarEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.SnoozeCalendarGroupCalendarEventExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarEventExceptionOccurrenceInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendargroupcalendareventexceptionoccurrenceinstance.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
