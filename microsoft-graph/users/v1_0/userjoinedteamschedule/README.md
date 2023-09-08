
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamschedule` Documentation

The `userjoinedteamschedule` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteamschedule"
```


### Client Initialization

```go
client := userjoinedteamschedule.NewUserJoinedTeamScheduleClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamScheduleClient.CreateUpdateUserByIdJoinedTeamByIdSchedule`

```go
ctx := context.TODO()
id := userjoinedteamschedule.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteamschedule.Schedule{
	// ...
}


read, err := client.CreateUpdateUserByIdJoinedTeamByIdSchedule(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamScheduleClient.CreateUserByIdJoinedTeamByIdScheduleShare`

```go
ctx := context.TODO()
id := userjoinedteamschedule.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteamschedule.CreateUserByIdJoinedTeamByIdScheduleShareRequest{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdScheduleShare(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamScheduleClient.DeleteUserByIdJoinedTeamByIdSchedule`

```go
ctx := context.TODO()
id := userjoinedteamschedule.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.DeleteUserByIdJoinedTeamByIdSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamScheduleClient.GetUserByIdJoinedTeamByIdSchedule`

```go
ctx := context.TODO()
id := userjoinedteamschedule.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamByIdSchedule(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
