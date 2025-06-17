
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelplannerplan` Documentation

The `teamprimarychannelplannerplan` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelplannerplan"
```


### Client Initialization

```go
client := teamprimarychannelplannerplan.NewTeamPrimaryChannelPlannerPlanClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.CreateTeamPrimaryChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupID("groupId")

payload := teamprimarychannelplannerplan.PlannerPlan{
	// ...
}


read, err := client.CreateTeamPrimaryChannelPlannerPlan(ctx, id, payload, teamprimarychannelplannerplan.DefaultCreateTeamPrimaryChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.CreateTeamPrimaryChannelPlannerPlanArchive`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

payload := teamprimarychannelplannerplan.CreateTeamPrimaryChannelPlannerPlanArchiveRequest{
	// ...
}


read, err := client.CreateTeamPrimaryChannelPlannerPlanArchive(ctx, id, payload, teamprimarychannelplannerplan.DefaultCreateTeamPrimaryChannelPlannerPlanArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.CreateTeamPrimaryChannelPlannerPlanUnarchive`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

payload := teamprimarychannelplannerplan.CreateTeamPrimaryChannelPlannerPlanUnarchiveRequest{
	// ...
}


read, err := client.CreateTeamPrimaryChannelPlannerPlanUnarchive(ctx, id, payload, teamprimarychannelplannerplan.DefaultCreateTeamPrimaryChannelPlannerPlanUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.DeleteTeamPrimaryChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

read, err := client.DeleteTeamPrimaryChannelPlannerPlan(ctx, id, teamprimarychannelplannerplan.DefaultDeleteTeamPrimaryChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.GetTeamPrimaryChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

read, err := client.GetTeamPrimaryChannelPlannerPlan(ctx, id, teamprimarychannelplannerplan.DefaultGetTeamPrimaryChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.GetTeamPrimaryChannelPlannerPlansCount`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupID("groupId")

read, err := client.GetTeamPrimaryChannelPlannerPlansCount(ctx, id, teamprimarychannelplannerplan.DefaultGetTeamPrimaryChannelPlannerPlansCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.ListTeamPrimaryChannelPlannerPlans`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupID("groupId")

// alternatively `client.ListTeamPrimaryChannelPlannerPlans(ctx, id, teamprimarychannelplannerplan.DefaultListTeamPrimaryChannelPlannerPlansOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelPlannerPlansComplete(ctx, id, teamprimarychannelplannerplan.DefaultListTeamPrimaryChannelPlannerPlansOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.MoveTeamPrimaryChannelPlannerPlanToContainer`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

payload := teamprimarychannelplannerplan.MoveTeamPrimaryChannelPlannerPlanToContainerRequest{
	// ...
}


read, err := client.MoveTeamPrimaryChannelPlannerPlanToContainer(ctx, id, payload, teamprimarychannelplannerplan.DefaultMoveTeamPrimaryChannelPlannerPlanToContainerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelPlannerPlanClient.UpdateTeamPrimaryChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamprimarychannelplannerplan.NewGroupIdTeamPrimaryChannelPlannerPlanID("groupId", "plannerPlanId")

payload := teamprimarychannelplannerplan.PlannerPlan{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelPlannerPlan(ctx, id, payload, teamprimarychannelplannerplan.DefaultUpdateTeamPrimaryChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
