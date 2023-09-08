
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceexceptionoccurrence` Documentation

The `mecalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarviewinstanceexceptionoccurrence.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceExceptionOccurrenceClient.ListMeCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
