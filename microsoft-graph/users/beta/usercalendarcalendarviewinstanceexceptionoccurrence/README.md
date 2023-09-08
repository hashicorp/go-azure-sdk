
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `usercalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarcalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewInstanceExceptionOccurrenceClient.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
