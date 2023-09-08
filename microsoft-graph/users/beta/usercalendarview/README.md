
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarview` Documentation

The `usercalendarview` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usercalendarview"
```


### Client Initialization

```go
client := usercalendarview.NewUserCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

read, err := client.CreateUserByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.CreateUserByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

payload := usercalendarview.CreateUserByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateUserByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.GetUserByIdCalendarViewById`

```go
ctx := context.TODO()
id := usercalendarview.NewUserCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetUserByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.GetUserByIdCalendarViewCount`

```go
ctx := context.TODO()
id := usercalendarview.NewUserID("userIdValue")

read, err := client.GetUserByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarViewClient.ListUserByIdCalendarViews`

```go
ctx := context.TODO()
id := usercalendarview.NewUserID("userIdValue")

// alternatively `client.ListUserByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
