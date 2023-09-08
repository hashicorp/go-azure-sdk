
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarevent` Documentation

The `groupcalendarevent` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarevent"
```


### Client Initialization

```go
client := groupcalendarevent.NewGroupCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEvent`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupID("groupIdValue")

payload := groupcalendarevent.Event{
	// ...
}


read, err := client.CreateGroupByIdCalendarEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.CreateGroupByIdCalendarEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdForward`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.CreateGroupByIdCalendarEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.CreateGroupByIdCalendarEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.DeleteGroupByIdCalendarEventById`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.DeleteGroupByIdCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.GetGroupByIdCalendarEventById`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.GetGroupByIdCalendarEventCount`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdCalendarEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventClient.ListGroupByIdCalendarEvents`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdCalendarEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupCalendarEventClient.UpdateGroupByIdCalendarEventById`

```go
ctx := context.TODO()
id := groupcalendarevent.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

payload := groupcalendarevent.Event{
	// ...
}


read, err := client.UpdateGroupByIdCalendarEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
