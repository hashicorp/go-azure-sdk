
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarevent` Documentation

The `calendargroupcalendarevent` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/calendargroupcalendarevent"
```


### Client Initialization

```go
client := calendargroupcalendarevent.NewCalendarGroupCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventClient.AcceptCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.AcceptCalendarGroupCalendarEventRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.CancelCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.CancelCalendarGroupCalendarEventRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
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


### Example Usage: `CalendarGroupCalendarEventClient.DeclineCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.DeclineCalendarGroupCalendarEventRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarEvent(ctx, id, payload)
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

read, err := client.DeleteCalendarGroupCalendarEvent(ctx, id, calendargroupcalendarevent.DefaultDeleteCalendarGroupCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.DismissCalendarGroupCalendarEventReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DismissCalendarGroupCalendarEventReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.ForwardCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.ForwardCalendarGroupCalendarEventRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarEvent(ctx, id, payload)
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

read, err := client.GetCalendarGroupCalendarEvent(ctx, id, calendargroupcalendarevent.DefaultGetCalendarGroupCalendarEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.GetCalendarGroupCalendarEventsCount`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetCalendarGroupCalendarEventsCount(ctx, id, calendargroupcalendarevent.DefaultGetCalendarGroupCalendarEventsCountOperationOptions())
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

// alternatively `client.ListCalendarGroupCalendarEvents(ctx, id, calendargroupcalendarevent.DefaultListCalendarGroupCalendarEventsOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventsComplete(ctx, id, calendargroupcalendarevent.DefaultListCalendarGroupCalendarEventsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventClient.SnoozeCalendarGroupCalendarEventReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.SnoozeCalendarGroupCalendarEventReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarEventReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventClient.TentativelyAcceptCalendarGroupCalendarEvent`

```go
ctx := context.TODO()
id := calendargroupcalendarevent.NewMeCalendarGroupIdCalendarIdEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := calendargroupcalendarevent.TentativelyAcceptCalendarGroupCalendarEventRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
