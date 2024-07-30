
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarevent` Documentation

The `calendargroupcalendarevent` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendargroupcalendarevent"
```


### Client Initialization

```go
client := calendargroupcalendarevent.NewCalendarGroupCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

payload := calendargroupcalendarevent.Event{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateCalendarGroupCalendarEventDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventForward`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CreateCalendarGroupCalendarEventTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CreateCalendarGroupCalendarEventTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.DeleteCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteCalendarGroupCalendarEvent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.GetCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEvent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.GetCalendarGroupCalendarEventCount`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendarEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.ListCalendarGroupCalendarEvents`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListCalendarGroupCalendarEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventClient.UpdateCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.Event{
	// ...
}


read, err := client.UpdateCalendarGroupCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
