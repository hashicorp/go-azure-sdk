
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrence` Documentation

The `usercalendargroupcalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
