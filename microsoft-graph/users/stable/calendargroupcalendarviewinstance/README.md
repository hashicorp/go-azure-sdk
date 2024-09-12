
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewinstance` Documentation

The `calendargroupcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarviewinstance"
```


### Client Initialization

```go
client := calendargroupcalendarviewinstance.NewCalendarGroupCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.AcceptCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.AcceptCalendarGroupCalendarViewInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.CancelCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.CancelCalendarGroupCalendarViewInstanceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.DeclineCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.DeclineCalendarGroupCalendarViewInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.DismissCalendarGroupCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarViewInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.ForwardCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.ForwardCalendarGroupCalendarViewInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.GetCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstance(ctx, id, calendargroupcalendarviewinstance.DefaultGetCalendarGroupCalendarViewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.GetCalendarGroupCalendarViewInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewInstancesCount(ctx, id, calendargroupcalendarviewinstance.DefaultGetCalendarGroupCalendarViewInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.ListCalendarGroupCalendarViewInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewInstances(ctx, id, calendargroupcalendarviewinstance.DefaultListCalendarGroupCalendarViewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstancesComplete(ctx, id, calendargroupcalendarviewinstance.DefaultListCalendarGroupCalendarViewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.SnoozeCalendarGroupCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.SnoozeCalendarGroupCalendarViewInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarViewInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarViewInstanceClient.TentativelyAcceptCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarviewinstance.TentativelyAcceptCalendarGroupCalendarViewInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarViewInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
