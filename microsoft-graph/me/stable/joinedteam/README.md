
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteam` Documentation

The `joinedteam` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteam"
```


### Client Initialization

```go
client := joinedteam.NewJoinedTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamClient.CreateJoinedTeam`

```go
ctx := context.TODO()

payload := joinedteam.Team{
	// ...
}


read, err := client.CreateJoinedTeam(ctx, payload)
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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.CreateJoinedTeamCompleteMigration(ctx, id)
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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.DeleteJoinedTeam(ctx, id, joinedteam.DefaultDeleteJoinedTeamOperationOptions())
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
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetJoinedTeam(ctx, id, joinedteam.DefaultGetJoinedTeamOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.GetJoinedTeamsCount`

```go
ctx := context.TODO()


read, err := client.GetJoinedTeamsCount(ctx, joinedteam.DefaultGetJoinedTeamsCountOperationOptions())
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


// alternatively `client.ListJoinedTeams(ctx, joinedteam.DefaultListJoinedTeamsOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamsComplete(ctx, joinedteam.DefaultListJoinedTeamsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamClient.SendJoinedTeamActivityNotification`

```go
ctx := context.TODO()
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

payload := joinedteam.SendJoinedTeamActivityNotificationRequest{
	// ...
}


read, err := client.SendJoinedTeamActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamClient.UpdateJoinedTeam`

```go
ctx := context.TODO()
id := joinedteam.NewMeJoinedTeamID("teamIdValue")

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
