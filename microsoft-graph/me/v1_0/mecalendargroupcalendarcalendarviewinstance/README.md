
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendarcalendarviewinstance` Documentation

The `mecalendargroupcalendarcalendarviewinstance` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendarcalendarviewinstance"
```


### Client Initialization

```go
client := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancel`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancelRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdCancel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDecline`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDeclineRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDecline(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDismissReminder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForward`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForwardRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdForward(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminderRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdSnoozeReminder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

payload := mecalendargroupcalendarcalendarviewinstance.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAcceptRequest{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdTentativelyAccept(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceById`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceID("calendarGroupIdValue", "calendarIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarCalendarViewInstanceClient.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstances`

```go
ctx := context.TODO()
id := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewID("calendarGroupIdValue", "calendarIdValue", "eventIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstances(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstancesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
