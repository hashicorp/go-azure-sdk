
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/team` Documentation

The `team` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/team"
```


### Client Initialization

```go
client := team.NewTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamClient.CreateTeamArchive`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

payload := team.CreateTeamArchiveRequest{
	// ...
}


read, err := client.CreateTeamArchive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.CreateTeamClone`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

payload := team.CreateTeamCloneRequest{
	// ...
}


read, err := client.CreateTeamClone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.CreateTeamCompleteMigration`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

read, err := client.CreateTeamCompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.CreateTeamSendActivityNotification`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

payload := team.CreateTeamSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateTeamSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.CreateTeamUnarchive`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

read, err := client.CreateTeamUnarchive(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.CreateUpdateTeam`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

payload := team.Team{
	// ...
}


read, err := client.CreateUpdateTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.DeleteTeam`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

read, err := client.DeleteTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamClient.GetTeam`

```go
ctx := context.TODO()
id := team.NewGroupID("groupIdValue")

read, err := client.GetTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
