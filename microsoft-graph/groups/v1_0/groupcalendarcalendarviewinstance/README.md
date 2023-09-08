
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarcalendarviewinstance` Documentation

The `groupcalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendarcalendarviewinstance"
```


### Client Initialization

```go
client := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendarcalendarviewinstance.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.GetGroupByIdCalendarCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.GetGroupByIdCalendarCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarCalendarViewInstanceClient.ListGroupByIdCalendarCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
