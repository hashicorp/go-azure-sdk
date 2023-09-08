
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteaminstalledapp` Documentation

The `mejoinedteaminstalledapp` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteaminstalledapp"
```


### Client Initialization

```go
client := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamInstalledAppClient.CreateMeJoinedTeamByIdInstalledApp`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.CreateMeJoinedTeamByIdInstalledAppByIdUpgrade`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppID("teamIdValue", "teamsAppInstallationIdValue")

payload := mejoinedteaminstalledapp.CreateMeJoinedTeamByIdInstalledAppByIdUpgradeRequest{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdInstalledAppByIdUpgrade(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.DeleteMeJoinedTeamByIdInstalledAppById`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppID("teamIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteMeJoinedTeamByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.GetMeJoinedTeamByIdInstalledAppById`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppID("teamIdValue", "teamsAppInstallationIdValue")

read, err := client.GetMeJoinedTeamByIdInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.GetMeJoinedTeamByIdInstalledAppCount`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamByIdInstalledAppCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.ListMeJoinedTeamByIdInstalledApps`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.ListMeJoinedTeamByIdInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdInstalledAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamInstalledAppClient.UpdateMeJoinedTeamByIdInstalledAppById`

```go
ctx := context.TODO()
id := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppID("teamIdValue", "teamsAppInstallationIdValue")

payload := mejoinedteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdInstalledAppById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
