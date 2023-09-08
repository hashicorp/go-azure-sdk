
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamschedule` Documentation

The `groupteamschedule` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamschedule"
```


### Client Initialization

```go
client := groupteamschedule.NewGroupTeamScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamScheduleClient.CreateGroupByIdTeamScheduleShare`

```go
ctx := context.TODO()
id := groupteamschedule.NewGroupID("groupIdValue")

payload := groupteamschedule.CreateGroupByIdTeamScheduleShareRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamScheduleShare(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleClient.CreateUpdateGroupByIdTeamSchedule`

```go
ctx := context.TODO()
id := groupteamschedule.NewGroupID("groupIdValue")

payload := groupteamschedule.Schedule{
	// ...
}


read, err := client.CreateUpdateGroupByIdTeamSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleClient.DeleteGroupByIdTeamSchedule`

```go
ctx := context.TODO()
id := groupteamschedule.NewGroupID("groupIdValue")

read, err := client.DeleteGroupByIdTeamSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleClient.GetGroupByIdTeamSchedule`

```go
ctx := context.TODO()
id := groupteamschedule.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
