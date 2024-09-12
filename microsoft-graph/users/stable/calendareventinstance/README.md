
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstance` Documentation

The `calendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstance"
```


### Client Initialization

```go
client := calendareventinstance.NewCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarEventInstanceClient.AcceptCalendarEventInstance`

```go
ctx := context.TODO()
id := calendareventinstance.NewUserIdCalendarIdEventIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.AcceptCalendarEventInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarEventInstance(ctx, id, payload)
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
id := calendareventinstance.NewUserIdCalendarIdEventIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.CancelCalendarEventInstanceRequest{
	// ...
}


read, err := client.CancelCalendarEventInstance(ctx, id, payload)
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
id := calendareventinstance.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.DeclineCalendarEventInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarEventInstance(ctx, id, payload)
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
id := calendareventinstance.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarEventInstanceReminder(ctx, id)
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
id := calendareventinstance.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.ForwardCalendarEventInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarEventInstance(ctx, id, payload)
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
id := calendareventinstance.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

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
id := calendareventinstance.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

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
id := calendareventinstance.NewUserIdCalendarIdEventID("userIdValue", "calendarIdValue", "eventIdValue")

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
id := calendareventinstance.NewUserIdCalendarEventIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.SnoozeCalendarEventInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarEventInstanceReminder(ctx, id, payload)
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
id := calendareventinstance.NewUserIdCalendarIdEventIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendareventinstance.TentativelyAcceptCalendarEventInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
