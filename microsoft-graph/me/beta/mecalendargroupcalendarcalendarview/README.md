
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarview` Documentation

The `mecalendargroupcalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendargroupcalendarcalendarview"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

payload := mecalendargroupcalendarcalendarview.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewClient.ListMeCalendarGroupByIdCalendarByIdCalendarViews`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
