
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstance` Documentation

The `mecalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mecalendarcalendarviewinstance"
```


### Client Initialization

```go
client := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.CreateMeCalendarCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

payload := mecalendarcalendarviewinstance.CreateMeCalendarCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.GetMeCalendarByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.GetMeCalendarByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.GetMeCalendarCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceID("eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.GetMeCalendarCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.ListMeCalendarByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarCalendarViewInstanceClient.ListMeCalendarCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
