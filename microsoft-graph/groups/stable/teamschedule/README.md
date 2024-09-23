
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamschedule` Documentation

The `teamschedule` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamschedule"
```


### Client Initialization

```go
client := teamschedule.NewTeamScheduleClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamScheduleClient.DeleteTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupId")

read, err := client.DeleteTeamSchedule(ctx, id, teamschedule.DefaultDeleteTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleClient.GetTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupId")

read, err := client.GetTeamSchedule(ctx, id, teamschedule.DefaultGetTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleClient.SetTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupId")

payload := teamschedule.Schedule{
	// ...
}


read, err := client.SetTeamSchedule(ctx, id, payload, teamschedule.DefaultSetTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleClient.ShareTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupId")

payload := teamschedule.ShareTeamScheduleRequest{
	// ...
}


read, err := client.ShareTeamSchedule(ctx, id, payload, teamschedule.DefaultShareTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
