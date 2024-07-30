
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teaminstalledapp` Documentation

The `teaminstalledapp` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teaminstalledapp"
```


### Client Initialization

```go
client := teaminstalledapp.NewTeamInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamInstalledAppClient.CreateTeamInstalledApp`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupID("groupIdValue")

payload := teaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateTeamInstalledApp(ctx, id, payload)
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
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

payload := teaminstalledapp.CreateTeamInstalledAppUpgradeRequest{
	// ...
}


read, err := client.CreateTeamInstalledAppUpgrade(ctx, id, payload)
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
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteTeamInstalledApp(ctx, id)
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
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

read, err := client.GetTeamInstalledApp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamInstalledAppClient.GetTeamInstalledAppCount`

```go
ctx := context.TODO()
id := teaminstalledapp.NewGroupID("groupIdValue")

read, err := client.GetTeamInstalledAppCount(ctx, id)
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
id := teaminstalledapp.NewGroupID("groupIdValue")

// alternatively `client.ListTeamInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListTeamInstalledAppsComplete(ctx, id)
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
id := teaminstalledapp.NewGroupIdTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

payload := teaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateTeamInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
