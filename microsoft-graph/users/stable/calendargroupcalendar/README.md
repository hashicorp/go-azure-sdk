
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendar` Documentation

The `calendargroupcalendar` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendar"
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

read, err := client.DeleteCalendarGroupCalendar(ctx, id, calendargroupcalendar.DefaultDeleteCalendarGroupCalendarOperationOptions())
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

read, err := client.GetCalendarGroupCalendar(ctx, id, calendargroupcalendar.DefaultGetCalendarGroupCalendarOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarClient.GetCalendarGroupCalendarSchedules`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

payload := calendargroupcalendar.GetCalendarGroupCalendarSchedulesRequest{
	// ...
}


// alternatively `client.GetCalendarGroupCalendarSchedules(ctx, id, payload, calendargroupcalendar.DefaultGetCalendarGroupCalendarSchedulesOperationOptions())` can be used to do batched pagination
items, err := client.GetCalendarGroupCalendarSchedulesComplete(ctx, id, payload, calendargroupcalendar.DefaultGetCalendarGroupCalendarSchedulesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarClient.GetCalendarGroupCalendarsCount`

```go
ctx := context.TODO()
id := calendargroupcalendar.NewUserIdCalendarGroupID("userIdValue", "calendarGroupIdValue")

read, err := client.GetCalendarGroupCalendarsCount(ctx, id, calendargroupcalendar.DefaultGetCalendarGroupCalendarsCountOperationOptions())
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

// alternatively `client.ListCalendarGroupCalendars(ctx, id, calendargroupcalendar.DefaultListCalendarGroupCalendarsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarsComplete(ctx, id, calendargroupcalendar.DefaultListCalendarGroupCalendarsOperationOptions())
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
