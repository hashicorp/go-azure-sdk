
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstance` Documentation

The `calendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstance"
```


### Client Initialization

```go
client := calendarviewinstance.NewCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarViewInstanceClient.AcceptCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.AcceptCalendarViewInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.CancelCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.CancelCalendarViewInstanceRequest{
	// ...
}


read, err := client.CancelCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.DeclineCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.DeclineCalendarViewInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.DismissCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarViewInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.ForwardCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.ForwardCalendarViewInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarViewInstance(ctx, id, payload)
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

read, err := client.GetCalendarViewInstance(ctx, id, calendarviewinstance.DefaultGetCalendarViewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.GetCalendarViewInstancesCount`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarViewInstancesCount(ctx, id, calendarviewinstance.DefaultGetCalendarViewInstancesCountOperationOptions())
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

// alternatively `client.ListCalendarViewInstances(ctx, id, calendarviewinstance.DefaultListCalendarViewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstancesComplete(ctx, id, calendarviewinstance.DefaultListCalendarViewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarViewInstanceClient.SnoozeCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.SnoozeCalendarViewInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarViewInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarViewInstanceClient.TentativelyAcceptCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarviewinstance.NewGroupIdCalendarViewIdInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := calendarviewinstance.TentativelyAcceptCalendarViewInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
