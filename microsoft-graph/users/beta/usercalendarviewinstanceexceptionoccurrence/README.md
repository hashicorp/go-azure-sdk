
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewinstanceexceptionoccurrence` Documentation

The `usercalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewinstanceexceptionoccurrence.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.GetUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceExceptionOccurrenceClient.ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
