
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarview` Documentation

The `calendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarview"
```


### Client Initialization

```go
client := calendarcalendarview.NewCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewAccept`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewCancel`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewCancelRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewDecline`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewDeclineRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewDismissReminder`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.CreateCalendarCalendarViewDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewForward`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewForwardRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewSnoozeReminder`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewSnoozeReminderRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.CreateCalendarCalendarViewTentativelyAccept`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := calendarcalendarview.CreateCalendarCalendarViewTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateCalendarCalendarViewTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.GetCalendarCalendarView`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupIdCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetCalendarCalendarView(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.GetCalendarCalendarViewCount`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupID("groupIdValue")

read, err := client.GetCalendarCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarCalendarViewClient.ListCalendarCalendarViews`

```go
ctx := context.TODO()
id := calendarcalendarview.NewGroupID("groupIdValue")

// alternatively `client.ListCalendarCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListCalendarCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
