
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendargroupcalendar` Documentation

The `usercalendargroupcalendar` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/usercalendargroupcalendar"
```


### Client Initialization

```go
client := usercalendargroupcalendar.NewUserCalendarGroupCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserCalendarGroupCalendarClient.CreateUserByIdCalendarGroupByIdCalendar`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupID("userIdValue", "calendarGroupIdValue")

payload := usercalendargroupcalendar.Calendar{
	// ...
}


read, err := client.CreateUserByIdCalendarGroupByIdCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarClient.DeleteUserByIdCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.DeleteUserByIdCalendarGroupByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarClient.GetUserByIdCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarClient.GetUserByIdCalendarGroupByIdCalendarByIdSchedule`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := usercalendargroupcalendar.GetUserByIdCalendarGroupByIdCalendarByIdScheduleRequest{
	// ...
}


read, err := client.GetUserByIdCalendarGroupByIdCalendarByIdSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarClient.GetUserByIdCalendarGroupByIdCalendarCount`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupID("userIdValue", "calendarGroupIdValue")

read, err := client.GetUserByIdCalendarGroupByIdCalendarCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserCalendarGroupCalendarClient.ListUserByIdCalendarGroupByIdCalendars`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupID("userIdValue", "calendarGroupIdValue")

// alternatively `client.ListUserByIdCalendarGroupByIdCalendars(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdCalendarGroupByIdCalendarsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserCalendarGroupCalendarClient.UpdateUserByIdCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := usercalendargroupcalendar.NewUserCalendarGroupCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := usercalendargroupcalendar.Calendar{
	// ...
}


read, err := client.UpdateUserByIdCalendarGroupByIdCalendarById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
