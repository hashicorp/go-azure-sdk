
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/schedule` Documentation

The `schedule` SDK allows for interaction with the Azure Resource Manager Service `labservices` (API Version `2021-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/schedule"
```


### Client Initialization

```go
client := schedule.NewScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ScheduleClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "scheduleValue")

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
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "scheduleValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ScheduleClient.Get`

```go
ctx := context.TODO()
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "scheduleValue")
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
id := schedule.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
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
id := schedule.NewScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "scheduleValue")

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
