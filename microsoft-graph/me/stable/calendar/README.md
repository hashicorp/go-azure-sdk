
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendar` Documentation

The `calendar` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendar"
```


### Client Initialization

```go
client := calendar.NewCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarClient.CreateCalendar`

```go
ctx := context.TODO()

payload := calendar.Calendar{
	// ...
}


read, err := client.CreateCalendar(ctx, payload)
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
id := calendar.NewMeCalendarID("calendarIdValue")

read, err := client.DeleteCalendar(ctx, id)
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


read, err := client.GetCalendar(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetCalendarCount`

```go
ctx := context.TODO()


read, err := client.GetCalendarCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetMeCalendarSchedule`

```go
ctx := context.TODO()

payload := calendar.GetMeCalendarScheduleRequest{
	// ...
}


read, err := client.GetMeCalendarSchedule(ctx, payload)
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


// alternatively `client.ListCalendars(ctx)` can be used to do batched pagination
items, err := client.ListCalendarsComplete(ctx)
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

payload := calendar.Calendar{
	// ...
}


read, err := client.UpdateCalendar(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
