
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrence` Documentation

The `eventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := eventinstanceexceptionoccurrence.NewEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.AcceptEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.AcceptEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.AcceptEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.CancelEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.CancelEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.CancelEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.DeclineEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.DeclineEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.DeclineEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.DismissEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissEventInstanceExceptionOccurrenceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.ForwardEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.ForwardEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.ForwardEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.GetEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetEventInstanceExceptionOccurrence(ctx, id, eventinstanceexceptionoccurrence.DefaultGetEventInstanceExceptionOccurrenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.GetEventInstanceExceptionOccurrencesCount`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetEventInstanceExceptionOccurrencesCount(ctx, id, eventinstanceexceptionoccurrence.DefaultGetEventInstanceExceptionOccurrencesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.ListEventInstanceExceptionOccurrences`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListEventInstanceExceptionOccurrences(ctx, id, eventinstanceexceptionoccurrence.DefaultListEventInstanceExceptionOccurrencesOperationOptions())` can be used to do batched pagination
items, err := client.ListEventInstanceExceptionOccurrencesComplete(ctx, id, eventinstanceexceptionoccurrence.DefaultListEventInstanceExceptionOccurrencesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.SnoozeEventInstanceExceptionOccurrenceReminder`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.SnoozeEventInstanceExceptionOccurrenceReminderRequest{
	// ...
}


read, err := client.SnoozeEventInstanceExceptionOccurrenceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventInstanceExceptionOccurrenceClient.TentativelyAcceptEventInstanceExceptionOccurrence`

```go
ctx := context.TODO()
id := eventinstanceexceptionoccurrence.NewUserIdEventIdInstanceIdExceptionOccurrenceID("userIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := eventinstanceexceptionoccurrence.TentativelyAcceptEventInstanceExceptionOccurrenceRequest{
	// ...
}


read, err := client.TentativelyAcceptEventInstanceExceptionOccurrence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
