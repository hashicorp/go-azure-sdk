
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrence` Documentation

The `mecalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.CreateMeCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarviewexceptionoccurrence.CreateMeCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.GetMeCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.GetMeCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewExceptionOccurrenceClient.ListMeCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarviewexceptionoccurrence.NewMeCalendarViewID("eventIdValue")

// alternatively `client.ListMeCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
