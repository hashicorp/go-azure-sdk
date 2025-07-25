
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/event` Documentation

The `event` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/event"
```


### Client Initialization

```go
client := event.NewEventClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventClient.AcceptEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.AcceptEventRequest{
	// ...
}


read, err := client.AcceptEvent(ctx, id, payload, event.DefaultAcceptEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CancelEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.CancelEventRequest{
	// ...
}


read, err := client.CancelEvent(ctx, id, payload, event.DefaultCancelEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEvent`

```go
ctx := context.TODO()

payload := event.Event{
	// ...
}


read, err := client.CreateEvent(ctx, payload, event.DefaultCreateEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventPermanentDelete`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

read, err := client.CreateEventPermanentDelete(ctx, id, event.DefaultCreateEventPermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.DeclineEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.DeclineEventRequest{
	// ...
}


read, err := client.DeclineEvent(ctx, id, payload, event.DefaultDeclineEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.DeleteEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

read, err := client.DeleteEvent(ctx, id, event.DefaultDeleteEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.DismissEventReminder`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

read, err := client.DismissEventReminder(ctx, id, event.DefaultDismissEventReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.ForwardEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.ForwardEventRequest{
	// ...
}


read, err := client.ForwardEvent(ctx, id, payload, event.DefaultForwardEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.GetEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

read, err := client.GetEvent(ctx, id, event.DefaultGetEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.GetEventsCount`

```go
ctx := context.TODO()


read, err := client.GetEventsCount(ctx, event.DefaultGetEventsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.ListEvents`

```go
ctx := context.TODO()


// alternatively `client.ListEvents(ctx, event.DefaultListEventsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventsComplete(ctx, event.DefaultListEventsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventClient.SnoozeEventReminder`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.SnoozeEventReminderRequest{
	// ...
}


read, err := client.SnoozeEventReminder(ctx, id, payload, event.DefaultSnoozeEventReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.TentativelyAcceptEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.TentativelyAcceptEventRequest{
	// ...
}


read, err := client.TentativelyAcceptEvent(ctx, id, payload, event.DefaultTentativelyAcceptEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.UpdateEvent`

```go
ctx := context.TODO()
id := event.NewMeEventID("eventId")

payload := event.Event{
	// ...
}


read, err := client.UpdateEvent(ctx, id, payload, event.DefaultUpdateEventOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
