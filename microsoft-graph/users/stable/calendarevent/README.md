
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarevent` Documentation

The `calendarevent` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarevent"
```


### Client Initialization

```go
client := calendarevent.NewCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventClient.CreateCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewUserID("userIdValue")

payload := calendarevent.Event{
	// ...
}


read, err := client.CreateCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventAccept`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarEventID("userIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventCancel`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventCancelRequest{
	// ...
}


read, err := client.CreateCalendarEventCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventDecline`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarEventID("userIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventDeclineRequest{
	// ...
}


read, err := client.CreateCalendarEventDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventDismissReminder`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarEventID("userIdValue", "eventIdValue")

read, err := client.CreateCalendarEventDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventForward`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventForwardRequest{
	// ...
}


read, err := client.CreateCalendarEventForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventSnoozeReminder`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarEventSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.CreateCalendarEventTentativelyAccept`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := calendarevent.CreateCalendarEventTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventTentativelyAccept(ctx, id, payload)
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
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteCalendarEvent(ctx, id)
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
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarEvent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventClient.GetCalendarEventCount`

```go
ctx := context.TODO()
id := calendarevent.NewUserID("userIdValue")

read, err := client.GetCalendarEventCount(ctx, id)
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
id := calendarevent.NewUserID("userIdValue")

// alternatively `client.ListCalendarEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventClient.UpdateCalendarEvent`

```go
ctx := context.TODO()
id := calendarevent.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

payload := calendarevent.Event{
	// ...
}


read, err := client.UpdateCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
