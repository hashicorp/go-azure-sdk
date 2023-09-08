
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventinstance` Documentation

The `groupeventinstance` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventinstance"
```


### Client Initialization

```go
client := groupeventinstance.NewGroupEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateGroupByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.CreateGroupByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstance.CreateGroupByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.GetGroupByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.GetGroupByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceClient.ListGroupByIdEventByIdInstances`

```go
ctx := context.TODO()
id := groupeventinstance.NewGroupEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
