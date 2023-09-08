
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarview` Documentation

The `groupcalendarview` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarview"
```


### Client Initialization

```go
client := groupcalendarview.NewGroupCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.CreateGroupByIdCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdForward`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.CreateGroupByIdCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarview.CreateGroupByIdCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.GetGroupByIdCalendarViewById`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.GetGroupByIdCalendarViewCount`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewClient.ListGroupByIdCalendarViews`

```go
ctx := context.TODO()
id := groupcalendarview.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
