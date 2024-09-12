
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrenceinstance` Documentation

The `calendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendarviewexceptionoccurrenceinstance.NewCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.AcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.AcceptCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.CancelCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.CancelCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.DeclineCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.DeclineCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.DismissCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarViewExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.ForwardCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.ForwardCalendarViewExceptionOccurrenceInstanceRequest{
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


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstance(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.GetCalendarViewExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewExceptionOccurrenceInstancesCount(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultGetCalendarViewExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.ListCalendarViewExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarViewExceptionOccurrenceInstances(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewExceptionOccurrenceInstancesComplete(ctx, id, calendarviewexceptionoccurrenceinstance.DefaultListCalendarViewExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.SnoozeCalendarViewExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.SnoozeCalendarViewExceptionOccurrenceInstanceReminderRequest{
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


### Example Usage: `CalendarViewExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarViewExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendarviewexceptionoccurrenceinstance.NewMeCalendarViewIdExceptionOccurrenceIdInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := calendarviewexceptionoccurrenceinstance.TentativelyAcceptCalendarViewExceptionOccurrenceInstanceRequest{
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
