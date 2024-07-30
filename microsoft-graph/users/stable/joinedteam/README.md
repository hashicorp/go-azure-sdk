
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteam` Documentation

The `joinedteam` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteam"
```


### Client Initialization

```go
client := joinedteam.NewJoinedTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeam`

```go
ctx := context.TODO()
id := joinedteam.NewUserID("userIdValue")

payload := joinedteam.Team{
	// ...
}


read, err := client.CreateJoinedTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeamArchive`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteam.CreateJoinedTeamArchiveRequest{
	// ...
}


read, err := client.CreateJoinedTeamArchive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeamClone`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteam.CreateJoinedTeamCloneRequest{
	// ...
}


read, err := client.CreateJoinedTeamClone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeamCompleteMigration`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateJoinedTeamCompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeamSendActivityNotification`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteam.CreateJoinedTeamSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateJoinedTeamSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeamUnarchive`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateJoinedTeamUnarchive(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.DeleteJoinedTeam`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.DeleteJoinedTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.GetJoinedTeam`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.GetJoinedTeamCount`

```go
ctx := context.TODO()
id := joinedteam.NewUserID("userIdValue")

read, err := client.GetJoinedTeamCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.ListJoinedTeams`

```go
ctx := context.TODO()
id := joinedteam.NewUserID("userIdValue")

// alternatively `client.ListJoinedTeams(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamClient.UpdateJoinedTeam`

```go
ctx := context.TODO()
id := joinedteam.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteam.Team{
	// ...
}


read, err := client.UpdateJoinedTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
