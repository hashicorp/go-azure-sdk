
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarview` Documentation

The `calendarcalendarview` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarview"
```


### Client Initialization

```go
client := calendarcalendarview.NewCalendarCalendarViewClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewClient.AcceptCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.AcceptCalendarViewRequest{
	// ...
}


read, err := client.AcceptCalendarView(ctx, id, payload, calendarcalendarview.DefaultAcceptCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CancelCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.CancelCalendarViewRequest{
	// ...
}


read, err := client.CancelCalendarView(ctx, id, payload, calendarcalendarview.DefaultCancelCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.DeclineCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.DeclineCalendarViewRequest{
	// ...
}


read, err := client.DeclineCalendarView(ctx, id, payload, calendarcalendarview.DefaultDeclineCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.DismissCalendarViewReminder`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

read, err := client.DismissCalendarViewReminder(ctx, id, calendarcalendarview.DefaultDismissCalendarViewReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.ForwardCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.ForwardCalendarViewRequest{
	// ...
}


read, err := client.ForwardCalendarView(ctx, id, payload, calendarcalendarview.DefaultForwardCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.GetCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

read, err := client.GetCalendarView(ctx, id, calendarcalendarview.DefaultGetCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.GetCalendarViewCount`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupID("groupId")

read, err := client.GetCalendarViewCount(ctx, id, calendarcalendarview.DefaultGetCalendarViewCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.ListCalendarViews`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupID("groupId")

// alternatively `client.ListCalendarViews(ctx, id, calendarcalendarview.DefaultListCalendarViewsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewsComplete(ctx, id, calendarcalendarview.DefaultListCalendarViewsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewClient.SnoozeCalendarViewReminder`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.SnoozeCalendarViewReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewReminder(ctx, id, payload, calendarcalendarview.DefaultSnoozeCalendarViewReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.TentativelyAcceptCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupId", "eventId")

payload := calendarcalendarview.TentativelyAcceptCalendarViewRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarView(ctx, id, payload, calendarcalendarview.DefaultTentativelyAcceptCalendarViewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
