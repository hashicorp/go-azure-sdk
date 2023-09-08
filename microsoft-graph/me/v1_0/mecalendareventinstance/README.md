
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventinstance` Documentation

The `mecalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendareventinstance"
```


### Client Initialization

```go
client := mecalendareventinstance.NewMeCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.CreateMeCalendarEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendareventinstance.CreateMeCalendarEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.GetMeCalendarByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.GetMeCalendarByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.GetMeCalendarEventByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventInstanceID("calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.GetMeCalendarEventByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventID("calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarEventInstanceClient.ListMeCalendarByIdEventByIdInstances`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarEventInstanceClient.ListMeCalendarEventByIdInstances`

```go
ctx := context.TODO()
id := mecalendareventinstance.NewMeCalendarEventID("calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
