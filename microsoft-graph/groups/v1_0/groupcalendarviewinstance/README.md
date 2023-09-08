
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarviewinstance` Documentation

The `groupcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarviewinstance"
```


### Client Initialization

```go
client := groupcalendarviewinstance.NewGroupCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.CreateGroupByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarviewinstance.CreateGroupByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.GetGroupByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.GetGroupByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarViewInstanceClient.ListGroupByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := groupcalendarviewinstance.NewGroupCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
