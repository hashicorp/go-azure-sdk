
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstance` Documentation

The `usercalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarviewinstance"
```


### Client Initialization

```go
client := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.CreateUserByIdCalendarCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarcalendarviewinstance.CreateUserByIdCalendarCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.GetUserByIdCalendarByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.GetUserByIdCalendarCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.GetUserByIdCalendarCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.ListUserByIdCalendarByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewInstanceClient.ListUserByIdCalendarCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
