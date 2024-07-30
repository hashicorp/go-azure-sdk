
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


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarGroupCalendarEventInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceForward`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.CreateCalendarGroupCalendarEventInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendareventinstance.CreateCalendarGroupCalendarEventInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarEventInstanceTentativelyAccept(ctx, id, payload)
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

read, err := client.GetCalendarGroupCalendarEventInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarEventInstanceClient.GetCalendarGroupCalendarEventInstanceCount`

```go
ctx := context.TODO()
id := calendargroupcalendareventinstance.NewUserIdCalendarGroupIdCalendarIdEventID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarEventInstanceCount(ctx, id)
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

// alternatively `client.ListCalendarGroupCalendarEventInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarEventInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
