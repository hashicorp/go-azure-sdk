
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarview` Documentation

The `usercalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarcalendarview"
```


### Client Initialization

```go
client := usercalendarcalendarview.NewUserCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdForward`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.CreateUserByIdCalendarCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

payload := usercalendarcalendarview.CreateUserByIdCalendarCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.GetUserByIdCalendarByIdCalendarViewById`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.GetUserByIdCalendarByIdCalendarViewCount`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarID("userIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.GetUserByIdCalendarCalendarViewById`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarCalendarViewID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.GetUserByIdCalendarCalendarViewCount`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserID("userIdValue")

read, err := client.GetUserByIdCalendarCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarCalendarViewClient.ListUserByIdCalendarByIdCalendarViews`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserCalendarID("userIdValue", "calendarIdValue")

// alternatively `client.ListUserByIdCalendarByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarCalendarViewClient.ListUserByIdCalendarCalendarViews`

```go
ctx := context.TODO()
id := usercalendarcalendarview.NewUserID("userIdValue")

// alternatively `client.ListUserByIdCalendarCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
