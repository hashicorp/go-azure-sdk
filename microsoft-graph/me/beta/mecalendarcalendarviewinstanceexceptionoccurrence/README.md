
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstanceexceptionoccurrence` Documentation

The `mecalendarcalendarviewinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendarcalendarviewinstanceexceptionoccurrence.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewInstanceExceptionOccurrenceClient.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
