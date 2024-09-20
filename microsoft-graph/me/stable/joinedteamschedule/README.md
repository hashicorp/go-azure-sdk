
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamschedule` Documentation

The `joinedteamschedule` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamschedule"
```


### Client Initialization

```go
client := joinedteamschedule.NewJoinedTeamScheduleClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamScheduleClient.DeleteJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewMeJoinedTeamID("teamId")

read, err := client.DeleteJoinedTeamSchedule(ctx, id, joinedteamschedule.DefaultDeleteJoinedTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleClient.GetJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewMeJoinedTeamID("teamId")

read, err := client.GetJoinedTeamSchedule(ctx, id, joinedteamschedule.DefaultGetJoinedTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleClient.SetJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewMeJoinedTeamID("teamId")

payload := joinedteamschedule.Schedule{
	// ...
}


read, err := client.SetJoinedTeamSchedule(ctx, id, payload, joinedteamschedule.DefaultSetJoinedTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleClient.ShareJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewMeJoinedTeamID("teamId")

payload := joinedteamschedule.ShareJoinedTeamScheduleRequest{
	// ...
}


read, err := client.ShareJoinedTeamSchedule(ctx, id, payload, joinedteamschedule.DefaultShareJoinedTeamScheduleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
