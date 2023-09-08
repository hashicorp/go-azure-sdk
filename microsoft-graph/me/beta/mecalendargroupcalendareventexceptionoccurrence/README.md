
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrence` Documentation

The `mecalendargroupcalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceClient.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
