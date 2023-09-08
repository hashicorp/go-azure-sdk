
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarview` Documentation

The `usercalendargroupcalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendargroupcalendarcalendarview"
```


### Client Initialization

```go
client := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendargroupcalendarcalendarview.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewById`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarCalendarViewClient.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViews`

```go
ctx := context.TODO()
id := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
