
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/event` Documentation

The `event` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/event"
```


### Client Initialization

```go
client := event.NewEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventClient.AcceptEvent`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.AcceptEventRequest{
	// ...
}


read, err := client.AcceptEvent(ctx, id, payload)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.CancelEventRequest{
	// ...
}


read, err := client.CancelEvent(ctx, id, payload)
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


### Example Usage: `EventClient.DeclineEvent`

```go
ctx := context.TODO()
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.DeclineEventRequest{
	// ...
}


read, err := client.DeclineEvent(ctx, id, payload)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.DismissEventReminder(ctx, id)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.ForwardEventRequest{
	// ...
}


read, err := client.ForwardEvent(ctx, id, payload)
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
id := event.NewGroupID("groupIdValue")

read, err := client.GetEventsCount(ctx, id, event.DefaultGetEventsCountOperationOptions())
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

// alternatively `client.ListEvents(ctx, id, event.DefaultListEventsOperationOptions())` can be used to do batched pagination
items, err := client.ListEventsComplete(ctx, id, event.DefaultListEventsOperationOptions())
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.SnoozeEventReminderRequest{
	// ...
}


read, err := client.SnoozeEventReminder(ctx, id, payload)
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
id := event.NewGroupIdEventID("groupIdValue", "eventIdValue")

payload := event.TentativelyAcceptEventRequest{
	// ...
}


read, err := client.TentativelyAcceptEvent(ctx, id, payload)
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
