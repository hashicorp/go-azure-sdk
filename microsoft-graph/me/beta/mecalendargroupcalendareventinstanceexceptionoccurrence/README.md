
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrence` Documentation

The `mecalendargroupcalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventinstanceexceptionoccurrence.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
