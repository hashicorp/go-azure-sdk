
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarevent` Documentation

The `usercalendargroupcalendarevent` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarevent"
```


### Client Initialization

```go
client := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEvent`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := usercalendargroupcalendarevent.Event{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.DeleteUserByIdCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.GetUserByIdCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.GetUserByIdCalendarGroupByIdCalendarByIdEventCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.ListUserByIdCalendarGroupByIdCalendarByIdEvents`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarGroupCalendarEventClient.UpdateUserByIdCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarevent.Event{
	// ...
}


read, err := client.UpdateUserByIdCalendarGroupByIdCalendarByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
