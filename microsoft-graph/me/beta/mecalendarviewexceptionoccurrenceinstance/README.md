
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceinstance` Documentation

The `mecalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewexceptionoccurrenceinstance.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceInstanceClient.ListMeCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
