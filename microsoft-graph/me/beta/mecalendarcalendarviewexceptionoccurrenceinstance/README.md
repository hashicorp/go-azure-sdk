
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceinstance` Documentation

The `mecalendarcalendarviewexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewexceptionoccurrenceinstance.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceInstanceClient.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
