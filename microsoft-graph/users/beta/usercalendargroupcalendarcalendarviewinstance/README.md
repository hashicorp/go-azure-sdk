
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstance` Documentation

The `usercalendargroupcalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarviewinstance"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := usercalendargroupcalendarcalendarviewinstance.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewInstanceClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
