
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupevent` Documentation

The `groupevent` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupevent"
```


### Client Initialization

```go
client := groupevent.NewGroupEventClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventClient.CreateGroupByIdEvent`

```go
ctx := context.TODO()
id := groupevent.NewGroupID("groupIdValue")

payload := groupevent.Event{
	// ...
}


read, err := client.CreateGroupByIdEvent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdAccept`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdCancel`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdDecline`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdDismissReminder`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.CreateGroupByIdEventByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdForward`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.CreateGroupByIdEventByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.CreateGroupByIdEventByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.DeleteGroupByIdEventById`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.DeleteGroupByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.GetGroupByIdEventById`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdEventById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.GetGroupByIdEventCount`

```go
ctx := context.TODO()
id := groupevent.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdEventCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventClient.ListGroupByIdEvents`

```go
ctx := context.TODO()
id := groupevent.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdEvents(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupEventClient.UpdateGroupByIdEventById`

```go
ctx := context.TODO()
id := groupevent.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupevent.Event{
	// ...
}


read, err := client.UpdateGroupByIdEventById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
