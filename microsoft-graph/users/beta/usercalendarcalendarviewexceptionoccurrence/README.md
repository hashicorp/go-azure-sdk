
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrence` Documentation

The `usercalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewExceptionOccurrenceClient.ListUserByIdCalendarCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
