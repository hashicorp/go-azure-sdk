
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrence` Documentation

The `usercalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewexceptionoccurrence.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceClient.ListUserByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrence.NewUserCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
