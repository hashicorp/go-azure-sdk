
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamschedule` Documentation

The `teamschedule` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamschedule"
```


### Client Initialization

```go
client := teamschedule.NewTeamScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamScheduleClient.DeleteTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupIdValue")

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
id := teamschedule.NewGroupID("groupIdValue")

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
id := teamschedule.NewGroupID("groupIdValue")

payload := teamschedule.Schedule{
	// ...
}


read, err := client.SetTeamSchedule(ctx, id, payload)
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
id := teamschedule.NewGroupID("groupIdValue")

payload := teamschedule.ShareTeamScheduleRequest{
	// ...
}


read, err := client.ShareTeamSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
