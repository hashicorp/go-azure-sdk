
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstance` Documentation

The `eventinstance` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstance"
```


### Client Initialization

```go
client := eventinstance.NewEventInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceClient.AcceptEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.AcceptEventInstanceRequest{
	// ...
}


read, err := client.AcceptEventInstance(ctx, id, payload, eventinstance.DefaultAcceptEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CancelEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.CancelEventInstanceRequest{
	// ...
}


read, err := client.CancelEventInstance(ctx, id, payload, eventinstance.DefaultCancelEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.DeclineEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.DeclineEventInstanceRequest{
	// ...
}


read, err := client.DeclineEventInstance(ctx, id, payload, eventinstance.DefaultDeclineEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.DismissEventInstanceReminder`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.DismissEventInstanceReminder(ctx, id, eventinstance.DefaultDismissEventInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.ForwardEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.ForwardEventInstanceRequest{
	// ...
}


read, err := client.ForwardEventInstance(ctx, id, payload, eventinstance.DefaultForwardEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.GetEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetEventInstance(ctx, id, eventinstance.DefaultGetEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.GetEventInstancesCount`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventID("groupId", "eventId")

read, err := client.GetEventInstancesCount(ctx, id, eventinstance.DefaultGetEventInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.ListEventInstances`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventID("groupId", "eventId")

// alternatively `client.ListEventInstances(ctx, id, eventinstance.DefaultListEventInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventInstancesComplete(ctx, id, eventinstance.DefaultListEventInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventInstanceClient.SnoozeEventInstanceReminder`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.SnoozeEventInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeEventInstanceReminder(ctx, id, payload, eventinstance.DefaultSnoozeEventInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.TentativelyAcceptEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupId", "eventId", "eventId1")

payload := eventinstance.TentativelyAcceptEventInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventInstance(ctx, id, payload, eventinstance.DefaultTentativelyAcceptEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
