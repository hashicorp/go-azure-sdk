
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendar` Documentation

The `calendargroupcalendar` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendar"
```


### Client Initialization

```go
client := calendargroupcalendar.NewCalendarGroupCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarClient.CreateCalendarGroupCalendar`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupID("userIdValue", "calendarGroupIdValue")

payload := calendargroupcalendar.Calendar{
	// ...
}


read, err := client.CreateCalendarGroupCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.DeleteCalendarGroupCalendar`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.DeleteCalendarGroupCalendar(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.GetCalendarGroupCalendar`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendar(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.GetCalendarGroupCalendarCount`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupID("userIdValue", "calendarGroupIdValue")

read, err := client.GetCalendarGroupCalendarCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.GetUserCalendarGroupCalendarSchedule`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := calendargroupcalendar.GetUserCalendarGroupCalendarScheduleRequest{
	// ...
}


read, err := client.GetUserCalendarGroupCalendarSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.ListCalendarGroupCalendars`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupID("userIdValue", "calendarGroupIdValue")

// alternatively `client.ListCalendarGroupCalendars(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarClient.UpdateCalendarGroupCalendar`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := calendargroupcalendar.Calendar{
	// ...
}


read, err := client.UpdateCalendarGroupCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
