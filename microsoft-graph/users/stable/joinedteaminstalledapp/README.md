
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledapp` Documentation

The `joinedteaminstalledapp` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledapp"
```


### Client Initialization

```go
client := joinedteaminstalledapp.NewJoinedTeamInstalledAppClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamInstalledAppClient.CreateJoinedTeamInstalledApp`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateJoinedTeamInstalledApp(ctx, id, payload, joinedteaminstalledapp.DefaultCreateJoinedTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.CreateJoinedTeamInstalledAppUpgrade`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId")

payload := joinedteaminstalledapp.CreateJoinedTeamInstalledAppUpgradeRequest{
	// ...
}


read, err := client.CreateJoinedTeamInstalledAppUpgrade(ctx, id, payload, joinedteaminstalledapp.DefaultCreateJoinedTeamInstalledAppUpgradeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.DeleteJoinedTeamInstalledApp`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId")

read, err := client.DeleteJoinedTeamInstalledApp(ctx, id, joinedteaminstalledapp.DefaultDeleteJoinedTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.GetJoinedTeamInstalledApp`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId")

read, err := client.GetJoinedTeamInstalledApp(ctx, id, joinedteaminstalledapp.DefaultGetJoinedTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.GetJoinedTeamInstalledAppsCount`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamInstalledAppsCount(ctx, id, joinedteaminstalledapp.DefaultGetJoinedTeamInstalledAppsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.ListJoinedTeamInstalledApps`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userId", "teamId")

// alternatively `client.ListJoinedTeamInstalledApps(ctx, id, joinedteaminstalledapp.DefaultListJoinedTeamInstalledAppsOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamInstalledAppsComplete(ctx, id, joinedteaminstalledapp.DefaultListJoinedTeamInstalledAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamInstalledAppClient.UpdateJoinedTeamInstalledApp`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userId", "teamId", "teamsAppInstallationId")

payload := joinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateJoinedTeamInstalledApp(ctx, id, payload, joinedteaminstalledapp.DefaultUpdateJoinedTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
