
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstance` Documentation

The `calendargroupcalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstance"
```


### Client Initialization

```go
client := calendargroupcalendarcalendarviewinstance.NewCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceCancel`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceDecline`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceForward`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.CreateCalendarGroupCalendarCalendarViewInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendargroupcalendarcalendarviewinstance.CreateCalendarGroupCalendarCalendarViewInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarGroupCalendarCalendarViewInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.GetCalendarGroupCalendarCalendarViewInstance`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarGroupCalendarCalendarViewInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.GetCalendarGroupCalendarCalendarViewInstanceCount`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetCalendarGroupCalendarCalendarViewInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarGroupCalendarCalendarViewInstanceClient.ListCalendarGroupCalendarCalendarViewInstances`

```go
ctx := context.TODO()
id := calendargroupcalendarcalendarviewinstance.NewUserIdCalendarGroupIdCalendarIdCalendarViewID("userIdValue", "calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListCalendarGroupCalendarCalendarViewInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarGroupCalendarCalendarViewInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
