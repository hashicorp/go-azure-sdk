
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstance` Documentation

The `eventexceptionoccurrenceinstance` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := eventexceptionoccurrenceinstance.NewEventExceptionOccurrenceInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.AcceptEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.AcceptEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptEventExceptionOccurrenceInstance(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultAcceptEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.CancelEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.CancelEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelEventExceptionOccurrenceInstance(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultCancelEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.DeclineEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.DeclineEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineEventExceptionOccurrenceInstance(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultDeclineEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.DismissEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.DismissEventExceptionOccurrenceInstanceReminder(ctx, id, eventexceptionoccurrenceinstance.DefaultDismissEventExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.ForwardEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.ForwardEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardEventExceptionOccurrenceInstance(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultForwardEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.GetEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

read, err := client.GetEventExceptionOccurrenceInstance(ctx, id, eventexceptionoccurrenceinstance.DefaultGetEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.GetEventExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

read, err := client.GetEventExceptionOccurrenceInstancesCount(ctx, id, eventexceptionoccurrenceinstance.DefaultGetEventExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.ListEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceID("groupId", "eventId", "eventId1")

// alternatively `client.ListEventExceptionOccurrenceInstances(ctx, id, eventexceptionoccurrenceinstance.DefaultListEventExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventExceptionOccurrenceInstancesComplete(ctx, id, eventexceptionoccurrenceinstance.DefaultListEventExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.SnoozeEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.SnoozeEventExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeEventExceptionOccurrenceInstanceReminder(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultSnoozeEventExceptionOccurrenceInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventExceptionOccurrenceInstanceClient.TentativelyAcceptEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := eventexceptionoccurrenceinstance.NewGroupIdEventIdExceptionOccurrenceIdInstanceID("groupId", "eventId", "eventId1", "eventId2")

payload := eventexceptionoccurrenceinstance.TentativelyAcceptEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventExceptionOccurrenceInstance(ctx, id, payload, eventexceptionoccurrenceinstance.DefaultTentativelyAcceptEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
