
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpooledschedule` Documentation

The `scalingplanpooledschedule` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2024-04-08-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpooledschedule"
```


### Client Initialization

```go
client := scalingplanpooledschedule.NewScalingPlanPooledScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScalingPlanPooledScheduleClient.Create`

```go
ctx := context.TODO()
id := scalingplanpooledschedule.NewPooledScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "pooledScheduleName")

payload := scalingplanpooledschedule.ScalingPlanPooledSchedule{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanPooledScheduleClient.Delete`

```go
ctx := context.TODO()
id := scalingplanpooledschedule.NewPooledScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "pooledScheduleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanPooledScheduleClient.Get`

```go
ctx := context.TODO()
id := scalingplanpooledschedule.NewPooledScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "pooledScheduleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanPooledScheduleClient.List`

```go
ctx := context.TODO()
id := scalingplanpooledschedule.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName")

// alternatively `client.List(ctx, id, scalingplanpooledschedule.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, scalingplanpooledschedule.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScalingPlanPooledScheduleClient.Update`

```go
ctx := context.TODO()
id := scalingplanpooledschedule.NewPooledScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "pooledScheduleName")

payload := scalingplanpooledschedule.ScalingPlanPooledSchedulePatch{
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
