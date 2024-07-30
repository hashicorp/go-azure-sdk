
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


### Example Usage: `TeamScheduleClient.CreateTeamScheduleShare`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupIdValue")

payload := teamschedule.CreateTeamScheduleShareRequest{
	// ...
}


read, err := client.CreateTeamScheduleShare(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleClient.CreateUpdateTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupIdValue")

payload := teamschedule.Schedule{
	// ...
}


read, err := client.CreateUpdateTeamSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleClient.DeleteTeamSchedule`

```go
ctx := context.TODO()
id := teamschedule.NewGroupID("groupIdValue")

read, err := client.DeleteTeamSchedule(ctx, id)
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

read, err := client.GetTeamSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
