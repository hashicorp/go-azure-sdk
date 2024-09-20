
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teaminstalledapp` Documentation

The `teaminstalledapp` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teaminstalledapp"
```


### Client Initialization

```go
client := teaminstalledapp.NewTeamInstalledAppClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamInstalledAppClient.CreateTeamInstalledApp`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupID("groupId")

payload := teaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateTeamInstalledApp(ctx, id, payload, teaminstalledapp.DefaultCreateTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.CreateTeamInstalledAppUpgrade`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId")

payload := teaminstalledapp.CreateTeamInstalledAppUpgradeRequest{
	// ...
}


read, err := client.CreateTeamInstalledAppUpgrade(ctx, id, payload, teaminstalledapp.DefaultCreateTeamInstalledAppUpgradeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.DeleteTeamInstalledApp`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId")

read, err := client.DeleteTeamInstalledApp(ctx, id, teaminstalledapp.DefaultDeleteTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.GetTeamInstalledApp`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId")

read, err := client.GetTeamInstalledApp(ctx, id, teaminstalledapp.DefaultGetTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.GetTeamInstalledAppsCount`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupID("groupId")

read, err := client.GetTeamInstalledAppsCount(ctx, id, teaminstalledapp.DefaultGetTeamInstalledAppsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.ListTeamInstalledApps`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupID("groupId")

// alternatively `client.ListTeamInstalledApps(ctx, id, teaminstalledapp.DefaultListTeamInstalledAppsOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamInstalledAppsComplete(ctx, id, teaminstalledapp.DefaultListTeamInstalledAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamInstalledAppClient.UpdateTeamInstalledApp`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupId", "teamsAppInstallationId")

payload := teaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateTeamInstalledApp(ctx, id, payload, teaminstalledapp.DefaultUpdateTeamInstalledAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
