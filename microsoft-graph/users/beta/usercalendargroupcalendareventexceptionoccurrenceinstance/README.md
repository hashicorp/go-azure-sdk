
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstance` Documentation

The `usercalendargroupcalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendargroupcalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
