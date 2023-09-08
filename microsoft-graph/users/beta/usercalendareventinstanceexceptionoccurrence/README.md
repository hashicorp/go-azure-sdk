
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventinstanceexceptionoccurrence` Documentation

The `usercalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventinstanceexceptionoccurrence.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventInstanceExceptionOccurrenceClient.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
