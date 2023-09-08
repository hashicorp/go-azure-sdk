
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrence` Documentation

The `mecalendargroupcalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
