
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/schedule` Documentation

The `schedule` SDK allows for interaction with Azure Resource Manager `labservices` (API Version `2023-06-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/schedule"
```


### Client Initialization

```go
client := schedule.NewScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScheduleClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "scheduleName")

payload := schedule.Schedule{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScheduleClient.Delete`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "scheduleName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ScheduleClient.Get`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "scheduleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScheduleClient.ListByLab`

```go
ctx := context.TODO()
id := schedule.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName")

// alternatively `client.ListByLab(ctx, id)` can be used to do batched pagination
items, err := client.ListByLabComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScheduleClient.Update`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "scheduleName")

payload := schedule.ScheduleUpdate{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
