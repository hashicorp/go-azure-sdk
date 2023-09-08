
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventinstanceexceptionoccurrence` Documentation

The `mecalendareventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendareventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := mecalendareventinstanceexceptionoccurrence.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventInstanceExceptionOccurrenceClient.ListMeCalendarEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeCalendarEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
