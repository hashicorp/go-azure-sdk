
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrence` Documentation

The `groupcalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventexceptionoccurrence.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.GetGroupByIdCalendarEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventExceptionOccurrenceClient.ListGroupByIdCalendarEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := groupcalendareventexceptionoccurrence.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
