
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendar` Documentation

The `calendar` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendar"
```


### Client Initialization

```go
client := calendar.NewCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarClient.GetCalendar`

```go
ctx := context.TODO()
id := calendar.NewGroupID("groupIdValue")

read, err := client.GetCalendar(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetGroupCalendarSchedule`

```go
ctx := context.TODO()
id := calendar.NewGroupID("groupIdValue")

payload := calendar.GetGroupCalendarScheduleRequest{
	// ...
}


read, err := client.GetGroupCalendarSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
