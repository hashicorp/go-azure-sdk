
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrence` Documentation

The `usercalendargroupcalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
