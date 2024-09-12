
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarcalendarview` Documentation

The `calendargroupcalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarcalendarview"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarview.NewCalendarGroupCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.AcceptCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.AcceptCalendarGroupCalendarViewRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CancelCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CancelCalendarGroupCalendarViewRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.DeclineCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.DeclineCalendarGroupCalendarViewRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.DismissCalendarGroupCalendarViewReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DismissCalendarGroupCalendarViewReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.ForwardCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.ForwardCalendarGroupCalendarViewRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.GetCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarView(ctx, id, calendargroupcalendarcalendarview.DefaultGetCalendarGroupCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.GetCalendarGroupCalendarViewCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendarViewCount(ctx, id, calendargroupcalendarcalendarview.DefaultGetCalendarGroupCalendarViewCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.ListCalendarGroupCalendarViews`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListCalendarGroupCalendarViews(ctx, id, calendargroupcalendarcalendarview.DefaultListCalendarGroupCalendarViewsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewsComplete(ctx, id, calendargroupcalendarcalendarview.DefaultListCalendarGroupCalendarViewsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.SnoozeCalendarGroupCalendarViewReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.SnoozeCalendarGroupCalendarViewReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarViewReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.TentativelyAcceptCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.TentativelyAcceptCalendarGroupCalendarViewRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
