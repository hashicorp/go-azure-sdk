
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendareventinstance` Documentation

The `usercalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendareventinstance"
```


### Client Initialization

```go
client := usercalendareventinstance.NewUserCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.CreateUserByIdCalendarEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := usercalendareventinstance.CreateUserByIdCalendarEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.GetUserByIdCalendarByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.GetUserByIdCalendarByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.GetUserByIdCalendarEventByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.GetUserByIdCalendarEventByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarEventInstanceClient.ListUserByIdCalendarByIdEventByIdInstances`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarEventInstanceClient.ListUserByIdCalendarEventByIdInstances`

```go
ctx := context.TODO()
id := usercalendareventinstance.NewUserCalendarEventID("userIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
