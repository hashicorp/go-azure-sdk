
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendar` Documentation

The `calendar` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendar"
```


### Client Initialization

```go
client := calendar.NewCalendarClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CalendarClient.CreateCalendarPermanentDelete`

```go
ctx := context.TODO()
id := calendar.NewGroupID("groupId")

read, err := client.CreateCalendarPermanentDelete(ctx, id, calendar.DefaultCreateCalendarPermanentDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetCalendar`

```go
ctx := context.TODO()
id := calendar.NewGroupID("groupId")

read, err := client.GetCalendar(ctx, id, calendar.DefaultGetCalendarOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CalendarClient.GetCalendarSchedules`

```go
ctx := context.TODO()
id := calendar.NewGroupID("groupId")

payload := calendar.GetCalendarSchedulesRequest{
	// ...
}


// alternatively `client.GetCalendarSchedules(ctx, id, payload, calendar.DefaultGetCalendarSchedulesOperationOptions())` can be used to do batched pagination
items, err := client.GetCalendarSchedulesComplete(ctx, id, payload, calendar.DefaultGetCalendarSchedulesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
