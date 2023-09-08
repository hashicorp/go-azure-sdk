
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meeventinstance` Documentation

The `meeventinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meeventinstance"
```


### Client Initialization

```go
client := meeventinstance.NewMeEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.CreateMeEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

payload := meeventinstance.CreateMeEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.GetMeEventByIdInstanceById`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.GetMeEventByIdInstanceCount`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventID("eventIdValue")

read, err := client.GetMeEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeEventInstanceClient.ListMeEventByIdInstances`

```go
ctx := context.TODO()
id := meeventinstance.NewMeEventID("eventIdValue")

// alternatively `client.ListMeEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
