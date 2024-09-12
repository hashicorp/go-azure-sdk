
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstance` Documentation

The `calendargroupcalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstance"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstance.NewCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.AcceptCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.AcceptCalendarGroupCalendarViewInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CancelCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CancelCalendarGroupCalendarViewInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.DeclineCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.DeclineCalendarGroupCalendarViewInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.DismissCalendarGroupCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarViewInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.ForwardCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.ForwardCalendarGroupCalendarViewInstanceRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.GetCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarViewInstance(ctx, id, calendargroupcalendarcalendarviewinstance.DefaultGetCalendarGroupCalendarViewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.GetCalendarGroupCalendarViewInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarViewInstancesCount(ctx, id, calendargroupcalendarcalendarviewinstance.DefaultGetCalendarGroupCalendarViewInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.ListCalendarGroupCalendarViewInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarViewInstances(ctx, id, calendargroupcalendarcalendarviewinstance.DefaultListCalendarGroupCalendarViewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarViewInstancesComplete(ctx, id, calendargroupcalendarcalendarviewinstance.DefaultListCalendarGroupCalendarViewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.SnoozeCalendarGroupCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.SnoozeCalendarGroupCalendarViewInstanceReminderRequest{
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


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.TentativelyAcceptCalendarGroupCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.TentativelyAcceptCalendarGroupCalendarViewInstanceRequest{
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
