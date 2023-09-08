
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarview` Documentation

The `mecalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarview"
```


### Client Initialization

```go
client := mecalendarcalendarview.NewMeCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.CreateMeCalendarByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.CreateMeCalendarCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.CreateMeCalendarCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

payload := mecalendarcalendarview.CreateMeCalendarCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.GetMeCalendarByIdCalendarViewById`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.GetMeCalendarByIdCalendarViewCount`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarID("calendarIdValue")

read, err := client.GetMeCalendarByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.GetMeCalendarCalendarViewById`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.GetMeCalendarCalendarViewCount`

```go
ctx := context.TODO()


read, err := client.GetMeCalendarCalendarViewCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewClient.ListMeCalendarByIdCalendarViews`

```go
ctx := context.TODO()
id := mecalendarcalendarview.NewMeCalendarID("calendarIdValue")

// alternatively `client.ListMeCalendarByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewClient.ListMeCalendarCalendarViews`

```go
ctx := context.TODO()


// alternatively `client.ListMeCalendarCalendarViews(ctx)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
