
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamschedule` Documentation

The `joinedteamschedule` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamschedule"
```


### Client Initialization

```go
client := joinedteamschedule.NewJoinedTeamScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamScheduleClient.CreateJoinedTeamScheduleShare`

```go
ctx := context.TODO()
id := joinedteamschedule.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamschedule.CreateJoinedTeamScheduleShareRequest{
	// ...
}


read, err := client.CreateJoinedTeamScheduleShare(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleClient.CreateUpdateJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamschedule.Schedule{
	// ...
}


read, err := client.CreateUpdateJoinedTeamSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleClient.DeleteJoinedTeamSchedule`

```go
ctx := context.TODO()
id := joinedteamschedule.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.DeleteJoinedTeamSchedule(ctx, id)
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
id := joinedteamschedule.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeamSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
