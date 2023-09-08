
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarview` Documentation

The `mecalendarview` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarview"
```


### Client Initialization

```go
client := mecalendarview.NewMeCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

read, err := client.CreateMeCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdForward`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.CreateMeCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

payload := mecalendarview.CreateMeCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.GetMeCalendarViewById`

```go
ctx := context.TODO()
id := mecalendarview.NewMeCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.GetMeCalendarViewCount`

```go
ctx := context.TODO()


read, err := client.GetMeCalendarViewCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewClient.ListMeCalendarViews`

```go
ctx := context.TODO()


// alternatively `client.ListMeCalendarViews(ctx)` can be used to do batched pagination
items, err := client.ListMeCalendarViewsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
