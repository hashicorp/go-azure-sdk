
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/event` Documentation

The `event` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/event"
```


### Client Initialization

```go
client := event.NewEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventClient.CreateEvent`

```go
ctx := context.TODO()
id := event.NewGroupID("groupIdValue")

payload := event.Event{
	// ...
}


read, err := client.CreateEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventAccept`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventAcceptRequest{
	// ...
}


read, err := client.CreateEventAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventCancel`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventCancelRequest{
	// ...
}


read, err := client.CreateEventCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventDecline`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventDeclineRequest{
	// ...
}


read, err := client.CreateEventDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventDismissReminder`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.CreateEventDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventForward`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventForwardRequest{
	// ...
}


read, err := client.CreateEventForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventSnoozeReminder`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventSnoozeReminderRequest{
	// ...
}


read, err := client.CreateEventSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.CreateEventTentativelyAccept`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CreateEventTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateEventTentativelyAccept(ctx, id, payload)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.DeleteEvent(ctx, id)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.GetEvent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventClient.GetEventCount`

```go
ctx := context.TODO()
id := event.NewGroupID("groupIdValue")

read, err := client.GetEventCount(ctx, id)
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
id := event.NewGroupID("groupIdValue")

// alternatively `client.ListEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventClient.UpdateEvent`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.Event{
	// ...
}


read, err := client.UpdateEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
