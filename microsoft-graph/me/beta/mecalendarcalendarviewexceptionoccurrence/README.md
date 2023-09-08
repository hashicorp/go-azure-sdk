
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrence` Documentation

The `mecalendarcalendarviewexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewexceptionoccurrence.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.GetMeCalendarCalendarViewByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarCalendarViewByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.ListMeCalendarByIdCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewExceptionOccurrenceClient.ListMeCalendarCalendarViewByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarCalendarViewByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
