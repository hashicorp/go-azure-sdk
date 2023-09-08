
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendar` Documentation

The `groupcalendar` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupcalendar"
```


### Client Initialization

```go
client := groupcalendar.NewGroupCalendarClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupCalendarClient.GetGroupByIdCalendar`

```go
ctx := context.TODO()
id := groupcalendar.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdCalendar(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupCalendarClient.GetGroupByIdCalendarSchedule`

```go
ctx := context.TODO()
id := groupcalendar.NewGroupID("groupIdValue")

payload := groupcalendar.GetGroupByIdCalendarScheduleRequest{
	// ...
}


read, err := client.GetGroupByIdCalendarSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
