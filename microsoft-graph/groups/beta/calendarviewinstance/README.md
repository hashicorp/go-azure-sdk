
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstance` Documentation

The `calendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstance"
```


### Client Initialization

```go
client := calendarviewinstance.NewCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceAccept`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceCancel`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceDecline`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarViewInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceForward`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CreateCalendarViewInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CreateCalendarViewInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarViewInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.GetCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.GetCalendarViewInstanceCount`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarViewInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.ListCalendarViewInstances`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListCalendarViewInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarViewInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
