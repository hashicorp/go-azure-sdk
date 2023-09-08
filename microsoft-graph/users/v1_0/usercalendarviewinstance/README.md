
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarviewinstance` Documentation

The `usercalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendarviewinstance"
```


### Client Initialization

```go
client := usercalendarviewinstance.NewUserCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.CreateUserByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendarviewinstance.CreateUserByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.GetUserByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.GetUserByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewInstanceClient.ListUserByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := usercalendarviewinstance.NewUserCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
