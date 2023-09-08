
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstance` Documentation

The `mecalendargroupcalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendargroupcalendareventexceptionoccurrenceinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
