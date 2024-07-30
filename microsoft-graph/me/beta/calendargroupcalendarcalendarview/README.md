
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarview` Documentation

The `calendargroupcalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarcalendarview"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarview.NewCalendarGroupCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateCalendarGroupCalendarCalendarViewDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewForward`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.CreateCalendarGroupCalendarCalendarViewTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarcalendarview.CreateCalendarGroupCalendarCalendarViewTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.GetCalendarGroupCalendarCalendarView`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarIdCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarCalendarView(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.GetCalendarGroupCalendarCalendarViewCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewClient.ListCalendarGroupCalendarCalendarViews`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarview.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListCalendarGroupCalendarCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
