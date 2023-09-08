
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteaminstalledapp` Documentation

The `groupteaminstalledapp` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteaminstalledapp"
```


### Client Initialization

```go
client := groupteaminstalledapp.NewGroupTeamInstalledAppClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamInstalledAppClient.CreateGroupByIdTeamInstalledApp`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupID("groupIdValue")

payload := groupteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.CreateGroupByIdTeamInstalledApp(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamInstalledAppClient.CreateGroupByIdTeamInstalledAppByIdUpgrade`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

payload := groupteaminstalledapp.CreateGroupByIdTeamInstalledAppByIdUpgradeRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamInstalledAppByIdUpgrade(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamInstalledAppClient.DeleteGroupByIdTeamInstalledAppById`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

read, err := client.DeleteGroupByIdTeamInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamInstalledAppClient.GetGroupByIdTeamInstalledAppById`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

read, err := client.GetGroupByIdTeamInstalledAppById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamInstalledAppClient.GetGroupByIdTeamInstalledAppCount`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamInstalledAppCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamInstalledAppClient.ListGroupByIdTeamInstalledApps`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamInstalledApps(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamInstalledAppsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamInstalledAppClient.UpdateGroupByIdTeamInstalledAppById`

```go
ctx := context.TODO()
id := groupteaminstalledapp.NewGroupTeamInstalledAppID("groupIdValue", "teamsAppInstallationIdValue")

payload := groupteaminstalledapp.TeamsAppInstallation{
	// ...
}


read, err := client.UpdateGroupByIdTeamInstalledAppById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
