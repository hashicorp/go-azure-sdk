
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstance` Documentation

The `eventinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstance"
```


### Client Initialization

```go
client := eventinstance.NewEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceClient.CreateEventInstanceAccept`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceAcceptRequest{
	// ...
}


read, err := client.CreateEventInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceCancel`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceCancelRequest{
	// ...
}


read, err := client.CreateEventInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceDecline`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceDeclineRequest{
	// ...
}


read, err := client.CreateEventInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceDismissReminder`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateEventInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceForward`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceForwardRequest{
	// ...
}


read, err := client.CreateEventInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateEventInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.CreateEventInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CreateEventInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateEventInstanceTentativelyAccept(ctx, id, payload)
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
id := eventinstance.NewGroupIdEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceClient.GetEventInstanceCount`

```go
ctx := context.TODO()
id := eventinstance.NewGroupIdEventID("groupIdValue", "eventIdValue")

read, err := client.GetEventInstanceCount(ctx, id)
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
id := eventinstance.NewGroupIdEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListEventInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListEventInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
