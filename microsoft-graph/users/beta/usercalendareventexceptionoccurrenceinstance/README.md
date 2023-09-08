
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceinstance` Documentation

The `usercalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := usercalendareventexceptionoccurrenceinstance.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventExceptionOccurrenceInstanceClient.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
