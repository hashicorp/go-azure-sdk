
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendar` Documentation

The `usercalendar` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendar"
```


### Client Initialization

```go
client := usercalendar.NewUserCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarClient.CreateUserByIdCalendar`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

payload := usercalendar.Calendar{
	// ...
}


read, err := client.CreateUserByIdCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.DeleteUserByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendar.NewUserCalendarID("userIdValue", "calendarIdValue")

read, err := client.DeleteUserByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.GetUserByIdCalendar`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

read, err := client.GetUserByIdCalendar(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.GetUserByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendar.NewUserCalendarID("userIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.GetUserByIdCalendarByIdSchedule`

```go
ctx := context.TODO()
id := usercalendar.NewUserCalendarID("userIdValue", "calendarIdValue")

payload := usercalendar.GetUserByIdCalendarByIdScheduleRequest{
	// ...
}


read, err := client.GetUserByIdCalendarByIdSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.GetUserByIdCalendarCount`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

read, err := client.GetUserByIdCalendarCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.GetUserByIdCalendarSchedule`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

payload := usercalendar.GetUserByIdCalendarScheduleRequest{
	// ...
}


read, err := client.GetUserByIdCalendarSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.ListUserByIdCalendars`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

// alternatively `client.ListUserByIdCalendars(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarClient.UpdateUserByIdCalendar`

```go
ctx := context.TODO()
id := usercalendar.NewUserID("userIdValue")

payload := usercalendar.Calendar{
	// ...
}


read, err := client.UpdateUserByIdCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarClient.UpdateUserByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendar.NewUserCalendarID("userIdValue", "calendarIdValue")

payload := usercalendar.Calendar{
	// ...
}


read, err := client.UpdateUserByIdCalendarById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
