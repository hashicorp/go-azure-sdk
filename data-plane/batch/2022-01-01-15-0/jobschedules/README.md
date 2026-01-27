
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/jobschedules` Documentation

The `jobschedules` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/jobschedules"
```


### Client Initialization

```go
client := jobschedules.NewJobSchedulesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobSchedulesClient.JobScheduleAdd`

```go
ctx := context.TODO()

payload := jobschedules.JobScheduleAddParameter{
	// ...
}


read, err := client.JobScheduleAdd(ctx, payload, jobschedules.DefaultJobScheduleAddOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleDelete`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleDelete(ctx, id, jobschedules.DefaultJobScheduleDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleDisable`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleDisable(ctx, id, jobschedules.DefaultJobScheduleDisableOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleEnable`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleEnable(ctx, id, jobschedules.DefaultJobScheduleEnableOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleExists`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleExists(ctx, id, jobschedules.DefaultJobScheduleExistsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleGet`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleGet(ctx, id, jobschedules.DefaultJobScheduleGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleList`

```go
ctx := context.TODO()


// alternatively `client.JobScheduleList(ctx, jobschedules.DefaultJobScheduleListOperationOptions())` can be used to do batched pagination
items, err := client.JobScheduleListComplete(ctx, jobschedules.DefaultJobScheduleListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JobSchedulesClient.JobSchedulePatch`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

payload := jobschedules.JobSchedulePatchParameter{
	// ...
}


read, err := client.JobSchedulePatch(ctx, id, payload, jobschedules.DefaultJobSchedulePatchOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleTerminate`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

read, err := client.JobScheduleTerminate(ctx, id, jobschedules.DefaultJobScheduleTerminateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobSchedulesClient.JobScheduleUpdate`

```go
ctx := context.TODO()
id := jobschedules.NewJobscheduleID("jobScheduleId")

payload := jobschedules.JobScheduleUpdateParameter{
	// ...
}


read, err := client.JobScheduleUpdate(ctx, id, payload, jobschedules.DefaultJobScheduleUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
