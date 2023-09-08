
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarviewinstance` Documentation

The `mecalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendarviewinstance"
```


### Client Initialization

```go
client := mecalendarviewinstance.NewMeCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.CreateMeCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarviewinstance.CreateMeCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.GetMeCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.GetMeCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewID("eventIdValue")

read, err := client.GetMeCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarViewInstanceClient.ListMeCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := mecalendarviewinstance.NewMeCalendarViewID("eventIdValue")

// alternatively `client.ListMeCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
