
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceexceptionoccurrence` Documentation

The `meeventinstanceexceptionoccurrence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meeventinstanceexceptionoccurrence"
```


### Client Initialization

```go
client := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAccept`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdCancel`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdDecline`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdForward`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

payload := meeventinstanceexceptionoccurrence.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdExceptionOccurrenceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.GetMeEventByIdInstanceByIdExceptionOccurrenceById`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceID("eventIdValue", "eventId1Value", "eventId2Value")

read, err := client.GetMeEventByIdInstanceByIdExceptionOccurrenceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.GetMeEventByIdInstanceByIdExceptionOccurrenceCount`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdInstanceByIdExceptionOccurrenceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceExceptionOccurrenceClient.ListMeEventByIdInstanceByIdExceptionOccurrences`

```go
ctx := context.TODO()
id := meeventinstanceexceptionoccurrence.NewMeEventInstanceID("eventIdValue", "eventId1Value")

// alternatively `client.ListMeEventByIdInstanceByIdExceptionOccurrences(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdInstanceByIdExceptionOccurrencesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
