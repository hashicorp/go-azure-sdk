
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarview` Documentation

The `calendargroupcalendarview` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarview"
```


### Client Initialization

```go
client := calendargroupcalendarview.NewCalendarGroupCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewClient.AcceptCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.AcceptCalendarGroupCalendarViewRequest{
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


### Example Usage: `CalendarGroupCalendarViewClient.CancelCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.CancelCalendarGroupCalendarViewRequest{
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


### Example Usage: `CalendarGroupCalendarViewClient.DeclineCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.DeclineCalendarGroupCalendarViewRequest{
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


### Example Usage: `CalendarGroupCalendarViewClient.DismissCalendarGroupCalendarViewReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DismissCalendarGroupCalendarViewReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewClient.ForwardCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.ForwardCalendarGroupCalendarViewRequest{
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


### Example Usage: `CalendarGroupCalendarViewClient.GetCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarView(ctx, id, calendargroupcalendarview.DefaultGetCalendarGroupCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewClient.GetCalendarGroupCalendarViewCount`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendarViewCount(ctx, id, calendargroupcalendarview.DefaultGetCalendarGroupCalendarViewCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewClient.ListCalendarGroupCalendarViews`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarID("userIdValue", "calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListCalendarGroupCalendarViews(ctx, id, calendargroupcalendarview.DefaultListCalendarGroupCalendarViewsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewsComplete(ctx, id, calendargroupcalendarview.DefaultListCalendarGroupCalendarViewsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarViewClient.SnoozeCalendarGroupCalendarViewReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.SnoozeCalendarGroupCalendarViewReminderRequest{
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


### Example Usage: `CalendarGroupCalendarViewClient.TentativelyAcceptCalendarGroupCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarview.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarview.TentativelyAcceptCalendarGroupCalendarViewRequest{
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
