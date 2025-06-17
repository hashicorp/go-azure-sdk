
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelplannerplan` Documentation

The `teamchannelplannerplan` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelplannerplan"
```


### Client Initialization

```go
client := teamchannelplannerplan.NewTeamChannelPlannerPlanClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelPlannerPlanClient.CreateTeamChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannelplannerplan.PlannerPlan{
	// ...
}


read, err := client.CreateTeamChannelPlannerPlan(ctx, id, payload, teamchannelplannerplan.DefaultCreateTeamChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.CreateTeamChannelPlannerPlanArchive`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

payload := teamchannelplannerplan.CreateTeamChannelPlannerPlanArchiveRequest{
	// ...
}


read, err := client.CreateTeamChannelPlannerPlanArchive(ctx, id, payload, teamchannelplannerplan.DefaultCreateTeamChannelPlannerPlanArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.CreateTeamChannelPlannerPlanUnarchive`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

payload := teamchannelplannerplan.CreateTeamChannelPlannerPlanUnarchiveRequest{
	// ...
}


read, err := client.CreateTeamChannelPlannerPlanUnarchive(ctx, id, payload, teamchannelplannerplan.DefaultCreateTeamChannelPlannerPlanUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.DeleteTeamChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

read, err := client.DeleteTeamChannelPlannerPlan(ctx, id, teamchannelplannerplan.DefaultDeleteTeamChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.GetTeamChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

read, err := client.GetTeamChannelPlannerPlan(ctx, id, teamchannelplannerplan.DefaultGetTeamChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.GetTeamChannelPlannerPlansCount`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.GetTeamChannelPlannerPlansCount(ctx, id, teamchannelplannerplan.DefaultGetTeamChannelPlannerPlansCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.ListTeamChannelPlannerPlans`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelID("groupId", "channelId")

// alternatively `client.ListTeamChannelPlannerPlans(ctx, id, teamchannelplannerplan.DefaultListTeamChannelPlannerPlansOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelPlannerPlansComplete(ctx, id, teamchannelplannerplan.DefaultListTeamChannelPlannerPlansOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelPlannerPlanClient.MoveTeamChannelPlannerPlanToContainer`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

payload := teamchannelplannerplan.MoveTeamChannelPlannerPlanToContainerRequest{
	// ...
}


read, err := client.MoveTeamChannelPlannerPlanToContainer(ctx, id, payload, teamchannelplannerplan.DefaultMoveTeamChannelPlannerPlanToContainerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelPlannerPlanClient.UpdateTeamChannelPlannerPlan`

```go
ctx := context.TODO()
id := teamchannelplannerplan.NewGroupIdTeamChannelIdPlannerPlanID("groupId", "channelId", "plannerPlanId")

payload := teamchannelplannerplan.PlannerPlan{
	// ...
}


read, err := client.UpdateTeamChannelPlannerPlan(ctx, id, payload, teamchannelplannerplan.DefaultUpdateTeamChannelPlannerPlanOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
