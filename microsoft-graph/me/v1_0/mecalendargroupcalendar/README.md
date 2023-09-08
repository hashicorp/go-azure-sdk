
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendar` Documentation

The `mecalendargroupcalendar` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendargroupcalendar"
```


### Client Initialization

```go
client := mecalendargroupcalendar.NewMeCalendarGroupCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarGroupCalendarClient.CreateMeCalendarGroupByIdCalendar`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupID("calendarGroupIdValue")

payload := mecalendargroupcalendar.Calendar{
	// ...
}


read, err := client.CreateMeCalendarGroupByIdCalendar(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarClient.DeleteMeCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.DeleteMeCalendarGroupByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarClient.GetMeCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarClient.GetMeCalendarGroupByIdCalendarByIdSchedule`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

payload := mecalendargroupcalendar.GetMeCalendarGroupByIdCalendarByIdScheduleRequest{
	// ...
}


read, err := client.GetMeCalendarGroupByIdCalendarByIdSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarClient.GetMeCalendarGroupByIdCalendarCount`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupID("calendarGroupIdValue")

read, err := client.GetMeCalendarGroupByIdCalendarCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarGroupCalendarClient.ListMeCalendarGroupByIdCalendars`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupID("calendarGroupIdValue")

// alternatively `client.ListMeCalendarGroupByIdCalendars(ctx, id)` can be used to do batched pagination
items, err := client.ListMeCalendarGroupByIdCalendarsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarGroupCalendarClient.UpdateMeCalendarGroupByIdCalendarById`

```go
ctx := context.TODO()
id := mecalendargroupcalendar.NewMeCalendarGroupCalendarID("calendarGroupIdValue", "calendarIdValue")

payload := mecalendargroupcalendar.Calendar{
	// ...
}


read, err := client.UpdateMeCalendarGroupByIdCalendarById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
