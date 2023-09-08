
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarevent` Documentation

The `mecalendarevent` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarevent"
```


### Client Initialization

```go
client := mecalendarevent.NewMeCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEvent`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarID("calendarIdValue")

payload := mecalendarevent.Event{
	// ...
}


read, err := client.CreateMeCalendarByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdAccept`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdCancel`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdDecline`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.CreateMeCalendarByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdForward`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEvent`

```go
ctx := context.TODO()

payload := mecalendarevent.Event{
	// ...
}


read, err := client.CreateMeCalendarEvent(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdAccept`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdCancel`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdDecline`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.CreateMeCalendarEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdForward`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.CreateMeCalendarEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.CreateMeCalendarEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.DeleteMeCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.DeleteMeCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.DeleteMeCalendarEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.DeleteMeCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.GetMeCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.GetMeCalendarByIdEventCount`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarID("calendarIdValue")

read, err := client.GetMeCalendarByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.GetMeCalendarEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

read, err := client.GetMeCalendarEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.GetMeCalendarEventCount`

```go
ctx := context.TODO()


read, err := client.GetMeCalendarEventCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.ListMeCalendarByIdEvents`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarID("calendarIdValue")

// alternatively `client.ListMeCalendarByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventClient.ListMeCalendarEvents`

```go
ctx := context.TODO()


// alternatively `client.ListMeCalendarEvents(ctx)` can be used to do batched pagination
items, err := client.ListMeCalendarEventsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventClient.UpdateMeCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.Event{
	// ...
}


read, err := client.UpdateMeCalendarByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventClient.UpdateMeCalendarEventById`

```go
ctx := context.TODO()
id := mecalendarevent.NewMeCalendarEventID("eventIdValue")

payload := mecalendarevent.Event{
	// ...
}


read, err := client.UpdateMeCalendarEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
