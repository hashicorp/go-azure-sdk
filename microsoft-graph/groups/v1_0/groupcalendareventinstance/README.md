
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendareventinstance` Documentation

The `groupcalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupcalendareventinstance"
```


### Client Initialization

```go
client := groupcalendareventinstance.NewGroupCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.CreateGroupByIdCalendarEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupcalendareventinstance.CreateGroupByIdCalendarEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdCalendarEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.GetGroupByIdCalendarEventByIdInstanceById`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdCalendarEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.GetGroupByIdCalendarEventByIdInstanceCount`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdCalendarEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarEventInstanceClient.ListGroupByIdCalendarEventByIdInstances`

```go
ctx := context.TODO()
id := groupcalendareventinstance.NewGroupCalendarEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdCalendarEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdCalendarEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
