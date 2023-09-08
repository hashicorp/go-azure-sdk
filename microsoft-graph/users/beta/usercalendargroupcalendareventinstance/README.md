
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstance` Documentation

The `usercalendargroupcalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendareventinstance"
```


### Client Initialization

```go
client := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendareventinstance.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarEventInstanceClient.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstances`

```go
ctx := context.TODO()
id := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
