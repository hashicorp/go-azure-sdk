
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpersonalschedule` Documentation

The `scalingplanpersonalschedule` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2024-04-08-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/scalingplanpersonalschedule"
```


### Client Initialization

```go
client := scalingplanpersonalschedule.NewScalingPlanPersonalScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScalingPlanPersonalScheduleClient.Create`

```go
ctx := context.TODO()
id := scalingplanpersonalschedule.NewPersonalScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "personalScheduleName")

payload := scalingplanpersonalschedule.ScalingPlanPersonalSchedule{
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


### Example Usage: `ScalingPlanPersonalScheduleClient.Delete`

```go
ctx := context.TODO()
id := scalingplanpersonalschedule.NewPersonalScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "personalScheduleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanPersonalScheduleClient.Get`

```go
ctx := context.TODO()
id := scalingplanpersonalschedule.NewPersonalScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "personalScheduleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScalingPlanPersonalScheduleClient.List`

```go
ctx := context.TODO()
id := scalingplanpersonalschedule.NewScalingPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName")

// alternatively `client.List(ctx, id, scalingplanpersonalschedule.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, scalingplanpersonalschedule.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScalingPlanPersonalScheduleClient.Update`

```go
ctx := context.TODO()
id := scalingplanpersonalschedule.NewPersonalScheduleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "scalingPlanName", "personalScheduleName")

payload := scalingplanpersonalschedule.ScalingPlanPersonalSchedulePatch{
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
