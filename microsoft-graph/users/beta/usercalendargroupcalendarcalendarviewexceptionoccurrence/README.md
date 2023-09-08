
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrence` Documentation

The `usercalendargroupcalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
