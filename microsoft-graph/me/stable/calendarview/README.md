
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarview` Documentation

The `calendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarview"
```


### Client Initialization

```go
client := calendarview.NewCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewClient.CreateCalendarViewAccept`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewCancel`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewCancelRequest{
	// ...
}


read, err := client.CreateCalendarViewCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewDecline`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewDeclineRequest{
	// ...
}


read, err := client.CreateCalendarViewDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewDismissReminder`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

read, err := client.CreateCalendarViewDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewForward`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewForwardRequest{
	// ...
}


read, err := client.CreateCalendarViewForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewSnoozeReminder`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarViewSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewClient.CreateCalendarViewTentativelyAccept`

```go
ctx := context.TODO()
id := calendarview.NewMeCalendarViewID("eventIdValue")

payload := calendarview.CreateCalendarViewTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewTentativelyAccept(ctx, id, payload)
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

read, err := client.GetCalendarView(ctx, id)
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


read, err := client.GetCalendarViewCount(ctx)
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


// alternatively `client.ListCalendarViews(ctx)` can be used to do batched pagination
items, err := client.ListCalendarViewsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
