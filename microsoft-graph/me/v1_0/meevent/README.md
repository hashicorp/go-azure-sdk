
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meevent` Documentation

The `meevent` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meevent"
```


### Client Initialization

```go
client := meevent.NewMeEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventClient.CreateMeEvent`

```go
ctx := context.TODO()

payload := meevent.Event{
	// ...
}


read, err := client.CreateMeEvent(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdAccept`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdCancel`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdCancelRequest{
	// ...
}


read, err := client.CreateMeEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdDecline`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdDismissReminder`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

read, err := client.CreateMeEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdForward`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdForwardRequest{
	// ...
}


read, err := client.CreateMeEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.CreateMeEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.CreateMeEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.DeleteMeEventById`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

read, err := client.DeleteMeEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.GetMeEventById`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

read, err := client.GetMeEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.GetMeEventCount`

```go
ctx := context.TODO()


read, err := client.GetMeEventCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventClient.ListMeEvents`

```go
ctx := context.TODO()


// alternatively `client.ListMeEvents(ctx)` can be used to do batched pagination
items, err := client.ListMeEventsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeEventClient.UpdateMeEventById`

```go
ctx := context.TODO()
id := meevent.NewMeEventID("eventIdValue")

payload := meevent.Event{
	// ...
}


read, err := client.UpdateMeEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
