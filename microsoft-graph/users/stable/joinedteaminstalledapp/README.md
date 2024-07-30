
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledapp` Documentation

The `joinedteaminstalledapp` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledapp"
```


### Client Initialization

```go
client := joinedteaminstalledapp.NewJoinedTeamInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamInstalledAppClient.CreateJoinedTeamInstalledApp`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateJoinedTeamInstalledApp(ctx, id, payload)
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
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userIdValue", "teamIdValue", "teamsAppInstallationIdValue")

payload := joinedteaminstalledapp.CreateJoinedTeamInstalledAppUpgradeRequest{
	// ...
}


read, err := client.CreateJoinedTeamInstalledAppUpgrade(ctx, id, payload)
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
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userIdValue", "teamIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteJoinedTeamInstalledApp(ctx, id)
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
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userIdValue", "teamIdValue", "teamsAppInstallationIdValue")

read, err := client.GetJoinedTeamInstalledApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamInstalledAppClient.GetJoinedTeamInstalledAppCount`

```go
ctx := context.TODO()
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeamInstalledAppCount(ctx, id)
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
id := joinedteaminstalledapp.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListJoinedTeamInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamInstalledAppsComplete(ctx, id)
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
id := joinedteaminstalledapp.NewUserIdJoinedTeamIdInstalledAppID("userIdValue", "teamIdValue", "teamsAppInstallationIdValue")

payload := joinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateJoinedTeamInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
