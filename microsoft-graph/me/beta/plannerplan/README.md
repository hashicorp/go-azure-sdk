
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/plannerplan` Documentation

The `plannerplan` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/plannerplan"
```


### Client Initialization

```go
client := plannerplan.NewPlannerPlanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlannerPlanClient.CreatePlannerPlan`

```go
ctx := context.TODO()

payload := plannerplan.PlannerPlan{
	// ...
}


read, err := client.CreatePlannerPlan(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.CreatePlannerPlanMoveToContainer`

```go
ctx := context.TODO()
id := plannerplan.NewMePlannerPlanID("plannerPlanIdValue")

payload := plannerplan.CreatePlannerPlanMoveToContainerRequest{
	// ...
}


read, err := client.CreatePlannerPlanMoveToContainer(ctx, id, payload)
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
id := plannerplan.NewMePlannerPlanID("plannerPlanIdValue")

read, err := client.DeletePlannerPlan(ctx, id)
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
id := plannerplan.NewMePlannerPlanID("plannerPlanIdValue")

read, err := client.GetPlannerPlan(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlannerPlanClient.GetPlannerPlanCount`

```go
ctx := context.TODO()


read, err := client.GetPlannerPlanCount(ctx)
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


// alternatively `client.ListPlannerPlans(ctx)` can be used to do batched pagination
items, err := client.ListPlannerPlansComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PlannerPlanClient.UpdatePlannerPlan`

```go
ctx := context.TODO()
id := plannerplan.NewMePlannerPlanID("plannerPlanIdValue")

payload := plannerplan.PlannerPlan{
	// ...
}


read, err := client.UpdatePlannerPlan(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
