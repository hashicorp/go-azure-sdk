
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrence` Documentation

The `meeventexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventexceptionoccurrence"
```


### Client Initialization

```go
client := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeEventByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.CreateMeEventByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

payload := meeventexceptionoccurrence.CreateMeEventByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.GetMeEventByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.GetMeEventByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventID("eventIdValue")

read, err := client.GetMeEventByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventExceptionOccurrenceClient.ListMeEventByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := meeventexceptionoccurrence.NewMeEventID("eventIdValue")

// alternatively `client.ListMeEventByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
