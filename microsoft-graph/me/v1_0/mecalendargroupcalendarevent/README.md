
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendarevent` Documentation

The `mecalendargroupcalendarevent` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendarevent"
```


### Client Initialization

```go
client := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEvent`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

payload := mecalendargroupcalendarevent.Event{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.CreateMeCalendarGroupByIdCalendarByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.DeleteMeCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.GetMeCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.GetMeCalendarGroupByIdCalendarByIdEventCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.ListMeCalendarGroupByIdCalendarByIdEvents`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarGroupCalendarEventClient.UpdateMeCalendarGroupByIdCalendarByIdEventById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarevent.Event{
	// ...
}


read, err := client.UpdateMeCalendarGroupByIdCalendarByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
