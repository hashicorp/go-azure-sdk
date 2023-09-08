
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendar` Documentation

The `mecalendar` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mecalendar"
```


### Client Initialization

```go
client := mecalendar.NewMeCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeCalendarClient.CreateMeCalendar`

```go
ctx := context.TODO()

payload := mecalendar.Calendar{
	// ...
}


read, err := client.CreateMeCalendar(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.DeleteMeCalendarById`

```go
ctx := context.TODO()
id := mecalendar.NewMeCalendarID("calendarIdValue")

read, err := client.DeleteMeCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.GetMeCalendar`

```go
ctx := context.TODO()


read, err := client.GetMeCalendar(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.GetMeCalendarById`

```go
ctx := context.TODO()
id := mecalendar.NewMeCalendarID("calendarIdValue")

read, err := client.GetMeCalendarById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.GetMeCalendarByIdSchedule`

```go
ctx := context.TODO()
id := mecalendar.NewMeCalendarID("calendarIdValue")

payload := mecalendar.GetMeCalendarByIdScheduleRequest{
	// ...
}


read, err := client.GetMeCalendarByIdSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.GetMeCalendarCount`

```go
ctx := context.TODO()


read, err := client.GetMeCalendarCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.GetMeCalendarSchedule`

```go
ctx := context.TODO()

payload := mecalendar.GetMeCalendarScheduleRequest{
	// ...
}


read, err := client.GetMeCalendarSchedule(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.ListMeCalendars`

```go
ctx := context.TODO()


// alternatively `client.ListMeCalendars(ctx)` can be used to do batched pagination
items, err := client.ListMeCalendarsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeCalendarClient.UpdateMeCalendar`

```go
ctx := context.TODO()

payload := mecalendar.Calendar{
	// ...
}


read, err := client.UpdateMeCalendar(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeCalendarClient.UpdateMeCalendarById`

```go
ctx := context.TODO()
id := mecalendar.NewMeCalendarID("calendarIdValue")

payload := mecalendar.Calendar{
	// ...
}


read, err := client.UpdateMeCalendarById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
