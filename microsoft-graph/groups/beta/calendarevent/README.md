
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarevent` Documentation

The `calendarevent` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarevent"
```


### Client Initialization

```go
client := calendarevent.NewCalendarEventClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventClient.AcceptCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.AcceptCalendarEventRequest{
	// ...
}


read, err := client.AcceptCalendarEvent(ctx, id, payload, calendarevent.DefaultAcceptCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CancelCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.CancelCalendarEventRequest{
	// ...
}


read, err := client.CancelCalendarEvent(ctx, id, payload, calendarevent.DefaultCancelCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupID("groupId")

payload := calendarevent.Event{
	// ...
}


read, err := client.CreateCalendarEvent(ctx, id, payload, calendarevent.DefaultCreateCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventPermanentDelete`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.CreateCalendarEventPermanentDelete(ctx, id, calendarevent.DefaultCreateCalendarEventPermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.DeclineCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.DeclineCalendarEventRequest{
	// ...
}


read, err := client.DeclineCalendarEvent(ctx, id, payload, calendarevent.DefaultDeclineCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.DeleteCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.DeleteCalendarEvent(ctx, id, calendarevent.DefaultDeleteCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.DismissCalendarEventReminder`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.DismissCalendarEventReminder(ctx, id, calendarevent.DefaultDismissCalendarEventReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.ForwardCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.ForwardCalendarEventRequest{
	// ...
}


read, err := client.ForwardCalendarEvent(ctx, id, payload, calendarevent.DefaultForwardCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.GetCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.GetCalendarEvent(ctx, id, calendarevent.DefaultGetCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.GetCalendarEventsCount`

```go
ctx := context.TODO()
id := calendarevent.NewGroupID("groupId")

read, err := client.GetCalendarEventsCount(ctx, id, calendarevent.DefaultGetCalendarEventsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.ListCalendarEvents`

```go
ctx := context.TODO()
id := calendarevent.NewGroupID("groupId")

// alternatively `client.ListCalendarEvents(ctx, id, calendarevent.DefaultListCalendarEventsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventsComplete(ctx, id, calendarevent.DefaultListCalendarEventsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventClient.SnoozeCalendarEventReminder`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.SnoozeCalendarEventReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventReminder(ctx, id, payload, calendarevent.DefaultSnoozeCalendarEventReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.TentativelyAcceptCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.TentativelyAcceptCalendarEventRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEvent(ctx, id, payload, calendarevent.DefaultTentativelyAcceptCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.UpdateCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewGroupIdCalendarEventID("groupId", "eventId")

payload := calendarevent.Event{
	// ...
}


read, err := client.UpdateCalendarEvent(ctx, id, payload, calendarevent.DefaultUpdateCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
