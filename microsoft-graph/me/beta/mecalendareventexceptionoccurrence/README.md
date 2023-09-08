
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrence` Documentation

The `mecalendareventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.CreateMeCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventexceptionoccurrence.CreateMeCalendarEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.GetMeCalendarByIdEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.GetMeCalendarByIdEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarByIdEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.GetMeCalendarEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.GetMeCalendarEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.ListMeCalendarByIdEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventID("eventIdValue")

// alternatively `client.ListMeCalendarByIdEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventExceptionOccurrenceClient.ListMeCalendarEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendareventexceptionoccurrence.NewMeCalendarEventID("eventIdValue")

// alternatively `client.ListMeCalendarEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
