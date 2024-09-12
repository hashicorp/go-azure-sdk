
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarcalendarviewinstance` Documentation

The `calendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/calendarcalendarviewinstance"
```


### Client Initialization

```go
client := calendarcalendarviewinstance.NewCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewInstanceClient.AcceptCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarIdCalendarViewIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.AcceptCalendarViewInstanceRequest{
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


### Example Usage: `CalendarCalendarViewInstanceClient.CancelCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.CancelCalendarViewInstanceRequest{
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


### Example Usage: `CalendarCalendarViewInstanceClient.DeclineCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarIdCalendarViewIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.DeclineCalendarViewInstanceRequest{
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


### Example Usage: `CalendarCalendarViewInstanceClient.DismissCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

read, err := client.DismissCalendarViewInstanceReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.ForwardCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.ForwardCalendarViewInstanceRequest{
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


### Example Usage: `CalendarCalendarViewInstanceClient.GetCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarIdCalendarViewIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetCalendarViewInstance(ctx, id, calendarcalendarviewinstance.DefaultGetCalendarViewInstanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.GetCalendarViewInstancesCount`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarIdCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetCalendarViewInstancesCount(ctx, id, calendarcalendarviewinstance.DefaultGetCalendarViewInstancesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.ListCalendarViewInstances`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarCalendarViewID("eventIdValue")

// alternatively `client.ListCalendarViewInstances(ctx, id, calendarcalendarviewinstance.DefaultListCalendarViewInstancesOperationOptions())` can be used to do batched pagination
items, err := client.ListCalendarViewInstancesComplete(ctx, id, calendarcalendarviewinstance.DefaultListCalendarViewInstancesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CalendarCalendarViewInstanceClient.SnoozeCalendarViewInstanceReminder`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarIdCalendarViewIdInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.SnoozeCalendarViewInstanceReminderRequest{
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


### Example Usage: `CalendarCalendarViewInstanceClient.TentativelyAcceptCalendarViewInstance`

```go
ctx := context.TODO()
id := calendarcalendarviewinstance.NewMeCalendarCalendarViewIdInstanceID("eventIdValue", "eventId1Value")

payload := calendarcalendarviewinstance.TentativelyAcceptCalendarViewInstanceRequest{
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
