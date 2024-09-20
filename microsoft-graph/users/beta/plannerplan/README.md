
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplan` Documentation

The `plannerplan` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplan"
```


### Client Initialization

```go
client := plannerplan.NewPlannerPlanClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlannerPlanClient.CreatePlannerPlan`

```go
ctx := context.TODO()
id := plannerplan.NewUserID("userId")

payload := plannerplan.PlannerPlan{
	// ...
}


read, err := client.CreatePlannerPlan(ctx, id, payload, plannerplan.DefaultCreatePlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.CreatePlannerPlanArchive`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

payload := plannerplan.CreatePlannerPlanArchiveRequest{
	// ...
}


read, err := client.CreatePlannerPlanArchive(ctx, id, payload, plannerplan.DefaultCreatePlannerPlanArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.CreatePlannerPlanUnarchive`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

payload := plannerplan.CreatePlannerPlanUnarchiveRequest{
	// ...
}


read, err := client.CreatePlannerPlanUnarchive(ctx, id, payload, plannerplan.DefaultCreatePlannerPlanUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.DeletePlannerPlan`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

read, err := client.DeletePlannerPlan(ctx, id, plannerplan.DefaultDeletePlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.GetPlannerPlan`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

read, err := client.GetPlannerPlan(ctx, id, plannerplan.DefaultGetPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.GetPlannerPlansCount`

```go
ctx := context.TODO()
id := plannerplan.NewUserID("userId")

read, err := client.GetPlannerPlansCount(ctx, id, plannerplan.DefaultGetPlannerPlansCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.ListPlannerPlans`

```go
ctx := context.TODO()
id := plannerplan.NewUserID("userId")

// alternatively `client.ListPlannerPlans(ctx, id, plannerplan.DefaultListPlannerPlansOperationOptions())` can be used to do batched pagination
items, err := client.ListPlannerPlansComplete(ctx, id, plannerplan.DefaultListPlannerPlansOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PlannerPlanClient.MovePlannerPlanToContainer`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

payload := plannerplan.MovePlannerPlanToContainerRequest{
	// ...
}


read, err := client.MovePlannerPlanToContainer(ctx, id, payload, plannerplan.DefaultMovePlannerPlanToContainerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.UpdatePlannerPlan`

```go
ctx := context.TODO()
id := plannerplan.NewUserIdPlannerPlanID("userId", "plannerPlanId")

payload := plannerplan.PlannerPlan{
	// ...
}


read, err := client.UpdatePlannerPlan(ctx, id, payload, plannerplan.DefaultUpdatePlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
