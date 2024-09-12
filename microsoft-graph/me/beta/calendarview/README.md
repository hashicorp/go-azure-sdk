
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarview` Documentation

The `calendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendarview"
```


### Client Initialization

```go
client := calendarview.NewCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewClient.AcceptCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.AcceptCalendarViewRequest{
	// ...
}


read, err := client.AcceptCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CancelCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CancelCalendarViewRequest{
	// ...
}


read, err := client.CancelCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.DeclineCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.DeclineCalendarViewRequest{
	// ...
}


read, err := client.DeclineCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.DismissCalendarViewReminder`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

read, err := client.DismissCalendarViewReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.ForwardCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.ForwardCalendarViewRequest{
	// ...
}


read, err := client.ForwardCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.GetCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

read, err := client.GetCalendarView(ctx, id, calendarview.DefaultGetCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.GetCalendarViewCount`

```go
ctx := context.TODO()


read, err := client.GetCalendarViewCount(ctx, calendarview.DefaultGetCalendarViewCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.ListCalendarViews`

```go
ctx := context.TODO()


// alternatively `client.ListCalendarViews(ctx, calendarview.DefaultListCalendarViewsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewsComplete(ctx, calendarview.DefaultListCalendarViewsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewClient.SnoozeCalendarViewReminder`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.SnoozeCalendarViewReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.TentativelyAcceptCalendarView`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.TentativelyAcceptCalendarViewRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarView(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
