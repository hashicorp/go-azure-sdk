
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstance` Documentation

The `calendargroupcalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstance"
```


### Client Initialization

```go
client := calendargroupcalendareventinstance.NewCalendarGroupCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.AcceptCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.AcceptCalendarGroupCalendarEventInstanceRequest{
	// ...
}


read, err := client.AcceptCalendarGroupCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CancelCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CancelCalendarGroupCalendarEventInstanceRequest{
	// ...
}


read, err := client.CancelCalendarGroupCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.DeclineCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.DeclineCalendarGroupCalendarEventInstanceRequest{
	// ...
}


read, err := client.DeclineCalendarGroupCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.DismissCalendarGroupCalendarEventInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.DismissCalendarGroupCalendarEventInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.ForwardCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.ForwardCalendarGroupCalendarEventInstanceRequest{
	// ...
}


read, err := client.ForwardCalendarGroupCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.GetCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarEventInstance(ctx, id, calendargroupcalendareventinstance.DefaultGetCalendarGroupCalendarEventInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.GetCalendarGroupCalendarEventInstancesCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventInstancesCount(ctx, id, calendargroupcalendareventinstance.DefaultGetCalendarGroupCalendarEventInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.ListCalendarGroupCalendarEventInstances`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarEventInstances(ctx, id, calendargroupcalendareventinstance.DefaultListCalendarGroupCalendarEventInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstancesComplete(ctx, id, calendargroupcalendareventinstance.DefaultListCalendarGroupCalendarEventInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.SnoozeCalendarGroupCalendarEventInstanceReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.SnoozeCalendarGroupCalendarEventInstanceReminderRequest{
	// ...
}


read, err := client.SnoozeCalendarGroupCalendarEventInstanceReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.TentativelyAcceptCalendarGroupCalendarEventInstance`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.TentativelyAcceptCalendarGroupCalendarEventInstanceRequest{
	// ...
}


read, err := client.TentativelyAcceptCalendarGroupCalendarEventInstance(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
