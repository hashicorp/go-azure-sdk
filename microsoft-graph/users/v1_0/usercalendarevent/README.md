
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarevent` Documentation

The `usercalendarevent` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarevent"
```


### Client Initialization

```go
client := usercalendarevent.NewUserCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEvent`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarID("userIdValue", "calendarIdValue")

payload := usercalendarevent.Event{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdAccept`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdCancel`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdDecline`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdForward`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEvent`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserID("userIdValue")

payload := usercalendarevent.Event{
	// ...
}


read, err := client.CreateUserByIdCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdAccept`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdCancel`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdDecline`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdForward`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.CreateUserByIdCalendarEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.CreateUserByIdCalendarEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.DeleteUserByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteUserByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.DeleteUserByIdCalendarEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteUserByIdCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.GetUserByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.GetUserByIdCalendarByIdEventCount`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarID("userIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.GetUserByIdCalendarEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.GetUserByIdCalendarEventCount`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserID("userIdValue")

read, err := client.GetUserByIdCalendarEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.ListUserByIdCalendarByIdEvents`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarID("userIdValue", "calendarIdValue")

// alternatively `client.ListUserByIdCalendarByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventClient.ListUserByIdCalendarEvents`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserID("userIdValue")

// alternatively `client.ListUserByIdCalendarEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventClient.UpdateUserByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.Event{
	// ...
}


read, err := client.UpdateUserByIdCalendarByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventClient.UpdateUserByIdCalendarEventById`

```go
ctx := context.TODO()
id := usercalendarevent.NewUserCalendarEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarevent.Event{
	// ...
}


read, err := client.UpdateUserByIdCalendarEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
