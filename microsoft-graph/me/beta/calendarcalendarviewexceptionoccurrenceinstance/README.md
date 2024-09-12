
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceinstance` Documentation

The `calendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarcalendarviewexceptionoccurrenceinstance.NewCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.AcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.CancelCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.CancelCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.DeclineCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.DismissCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarViewExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.ForwardCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstance(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarIdCalendarViewIdExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.ListCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstances(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendarcalendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.SnoozeCalendarViewExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewExceptionOccurrenceInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarcalendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarViewExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
