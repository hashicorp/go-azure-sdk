
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceinstance` Documentation

The `mecalendareventexceptionoccurrenceinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrenceinstance"
```


### Client Initialization

```go
client := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventexceptionoccurrenceinstance.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceID("calendarIdValue", "eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceInstanceClient.ListMeCalendarEventByIdExceptionOccurrenceByIdInstances`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarEventByIdExceptionOccurrenceByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdExceptionOccurrenceByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
