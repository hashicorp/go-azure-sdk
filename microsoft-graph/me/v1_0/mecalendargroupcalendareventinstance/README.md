
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendareventinstance` Documentation

The `mecalendargroupcalendareventinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendareventinstance"
```


### Client Initialization

```go
client := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendareventinstance.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdEventByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdEventByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarEventInstanceClient.ListMeCalendarGroupByIdCalendarByIdEventByIdInstances`

```go
ctx := context.TODO()
id := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdEventByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
