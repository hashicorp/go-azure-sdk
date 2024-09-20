
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstance` Documentation

The `calendareventinstance` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstance"
```


### Client Initialization

```go
client := calendareventinstance.NewCalendarEventInstanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceClient.AcceptCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.AcceptCalendarEventInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarEventInstance(ctx, id, payload, calendareventinstance.DefaultAcceptCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CancelCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.CancelCalendarEventInstanceRequest{
	// ...
}


read, err := client.CancelCalendarEventInstance(ctx, id, payload, calendareventinstance.DefaultCancelCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.DeclineCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.DeclineCalendarEventInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarEventInstance(ctx, id, payload, calendareventinstance.DefaultDeclineCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.DismissCalendarEventInstanceReminder`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.DismissCalendarEventInstanceReminder(ctx, id, calendareventinstance.DefaultDismissCalendarEventInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.ForwardCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.ForwardCalendarEventInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarEventInstance(ctx, id, payload, calendareventinstance.DefaultForwardCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.GetCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

read, err := client.GetCalendarEventInstance(ctx, id, calendareventinstance.DefaultGetCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.GetCalendarEventInstancesCount`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventID("groupId", "eventId")

read, err := client.GetCalendarEventInstancesCount(ctx, id, calendareventinstance.DefaultGetCalendarEventInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.ListCalendarEventInstances`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventID("groupId", "eventId")

// alternatively `client.ListCalendarEventInstances(ctx, id, calendareventinstance.DefaultListCalendarEventInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarEventInstancesComplete(ctx, id, calendareventinstance.DefaultListCalendarEventInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarEventInstanceClient.SnoozeCalendarEventInstanceReminder`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.SnoozeCalendarEventInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventInstanceReminder(ctx, id, payload, calendareventinstance.DefaultSnoozeCalendarEventInstanceReminderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.TentativelyAcceptCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupId", "eventId", "eventId1")

payload := calendareventinstance.TentativelyAcceptCalendarEventInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventInstance(ctx, id, payload, calendareventinstance.DefaultTentativelyAcceptCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
