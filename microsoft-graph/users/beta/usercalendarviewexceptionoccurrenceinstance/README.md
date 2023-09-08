
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceinstance` Documentation

The `usercalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendarviewexceptionoccurrenceinstance.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewExceptionOccurrenceInstanceClient.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
