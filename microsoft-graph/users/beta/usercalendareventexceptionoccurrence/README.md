
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrence` Documentation

The `usercalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventexceptionoccurrence.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.GetUserByIdCalendarEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.ListUserByIdCalendarByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceClient.ListUserByIdCalendarEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrence.NewUserCalendarEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
