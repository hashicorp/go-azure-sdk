
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstance` Documentation

The `calendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := calendareventexceptionoccurrenceinstance.NewCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.AcceptCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.AcceptCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.CancelCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.CancelCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.CancelCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.DeclineCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.DeclineCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.DismissCalendarEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.DismissCalendarEventExceptionOccurrenceInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.ForwardCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.ForwardCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.GetCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetCalendarEventExceptionOccurrenceInstance(ctx, id, calendareventexceptionoccurrenceinstance.DefaultGetCalendarEventExceptionOccurrenceInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.GetCalendarEventExceptionOccurrenceInstancesCount`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventExceptionOccurrenceInstancesCount(ctx, id, calendareventexceptionoccurrenceinstance.DefaultGetCalendarEventExceptionOccurrenceInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.ListCalendarEventExceptionOccurrenceInstances`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListCalendarEventExceptionOccurrenceInstances(ctx, id, calendareventexceptionoccurrenceinstance.DefaultListCalendarEventExceptionOccurrenceInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventExceptionOccurrenceInstancesComplete(ctx, id, calendareventexceptionoccurrenceinstance.DefaultListCalendarEventExceptionOccurrenceInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.SnoozeCalendarEventExceptionOccurrenceInstanceReminder`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.SnoozeCalendarEventExceptionOccurrenceInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventExceptionOccurrenceInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventExceptionOccurrenceInstanceClient.TentativelyAcceptCalendarEventExceptionOccurrenceInstance`

```go
ctx := context.TODO()
id := calendareventexceptionoccurrenceinstance.NewGroupIdCalendarEventIdExceptionOccurrenceIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := calendareventexceptionoccurrenceinstance.TentativelyAcceptCalendarEventExceptionOccurrenceInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventExceptionOccurrenceInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
