
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstance` Documentation

The `calendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstance"
```


### Client Initialization

```go
client := calendarcalendarviewinstance.NewCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceCancel`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceCancelRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceDecline`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceDeclineRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceDismissReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateCalendarCalendarViewInstanceDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceForward`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceForwardRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceSnoozeReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarIdCalendarViewIdInstanceID("userIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.CreateCalendarCalendarViewInstanceTentativelyAccept`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CreateCalendarCalendarViewInstanceTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewInstanceTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.GetCalendarCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarCalendarViewIdInstanceID("userIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarCalendarViewInstance(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.GetCalendarCalendarViewInstanceCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarCalendarViewID("userIdValue", "eventIdValue")

read, err := client.GetCalendarCalendarViewInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.ListCalendarCalendarViewInstances`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewUserIdCalendarCalendarViewID("userIdValue", "eventIdValue")

// alternatively `client.ListCalendarCalendarViewInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
