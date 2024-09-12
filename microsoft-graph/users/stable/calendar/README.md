
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendar` Documentation

The `calendar` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendar"
```


### Client Initialization

```go
client := calendar.NewCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarClient.CreateCalendar`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

payload := calendar.Calendar{
	// ...
}


read, err := client.CreateCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.DeleteCalendar`

```go
ctx := context.TODO()
id := calendar.NewUserIdCalendarID("userIdValue", "calendarIdValue")

read, err := client.DeleteCalendar(ctx, id, calendar.DefaultDeleteCalendarOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetCalendar`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

read, err := client.GetCalendar(ctx, id, calendar.DefaultGetCalendarOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetCalendarSchedules`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

payload := calendar.GetCalendarSchedulesRequest{
	// ...
}


// alternatively `client.GetCalendarSchedules(ctx, id, payload, calendar.DefaultGetCalendarSchedulesOperationOptions())` can be used to do batched pagination
items, err := client.GetCalendarSchedulesComplete(ctx, id, payload, calendar.DefaultGetCalendarSchedulesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarClient.GetCalendarsCount`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

read, err := client.GetCalendarsCount(ctx, id, calendar.DefaultGetCalendarsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.ListCalendars`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

// alternatively `client.ListCalendars(ctx, id, calendar.DefaultListCalendarsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarsComplete(ctx, id, calendar.DefaultListCalendarsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarClient.UpdateCalendar`

```go
ctx := context.TODO()
id := calendar.NewUserID("userIdValue")

payload := calendar.Calendar{
	// ...
}


read, err := client.UpdateCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
