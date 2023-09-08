
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
