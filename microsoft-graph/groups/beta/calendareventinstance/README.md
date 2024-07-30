
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstance` Documentation

The `calendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstance"
```


### Client Initialization

```go
client := calendareventinstance.NewCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceAccept`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceCancel`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceDecline`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarEventInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceForward`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.CreateCalendarEventInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CreateCalendarEventInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarEventInstanceTentativelyAccept(ctx, id, payload)
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
id := calendareventinstance.NewGroupIdCalendarEventIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarEventInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarEventInstanceClient.GetCalendarEventInstanceCount`

```go
ctx := context.TODO()
id := calendareventinstance.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarEventInstanceCount(ctx, id)
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
id := calendareventinstance.NewGroupIdCalendarEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListCalendarEventInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarEventInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
