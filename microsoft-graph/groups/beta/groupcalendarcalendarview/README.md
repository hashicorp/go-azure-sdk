
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarview` Documentation

The `groupcalendarcalendarview` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendarcalendarview"
```


### Client Initialization

```go
client := groupcalendarcalendarview.NewGroupCalendarCalendarViewClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.CreateGroupByIdCalendarCalendarViewByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdForward`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.CreateGroupByIdCalendarCalendarViewByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

payload := groupcalendarcalendarview.CreateGroupByIdCalendarCalendarViewByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.GetGroupByIdCalendarCalendarViewById`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.GetGroupByIdCalendarCalendarViewCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewClient.ListGroupByIdCalendarCalendarViews`

```go
ctx := context.TODO()
id := groupcalendarcalendarview.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdCalendarCalendarViews(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
