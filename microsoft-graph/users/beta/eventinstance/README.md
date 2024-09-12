
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstance` Documentation

The `eventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstance"
```


### Client Initialization

```go
client := eventinstance.NewEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceClient.AcceptEventInstance`

```go
ctx := context.TODO()
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.AcceptEventInstanceRequest{
	// ...
}


read, err := client.AcceptEventInstance(ctx, id, payload)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.CancelEventInstanceRequest{
	// ...
}


read, err := client.CancelEventInstance(ctx, id, payload)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.DeclineEventInstanceRequest{
	// ...
}


read, err := client.DeclineEventInstance(ctx, id, payload)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissEventInstanceReminder(ctx, id)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.ForwardEventInstanceRequest{
	// ...
}


read, err := client.ForwardEventInstance(ctx, id, payload)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

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
id := eventinstance.NewUserIdEventID("userIdValue", "eventIdValue")

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
id := eventinstance.NewUserIdEventID("userIdValue", "eventIdValue")

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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.SnoozeEventInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeEventInstanceReminder(ctx, id, payload)
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
id := eventinstance.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := eventinstance.TentativelyAcceptEventInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
