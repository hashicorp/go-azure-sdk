
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteam` Documentation

The `groupteam` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteam"
```


### Client Initialization

```go
client := groupteam.NewGroupTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamClient.CreateGroupByIdTeamArchive`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

payload := groupteam.CreateGroupByIdTeamArchiveRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamArchive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.CreateGroupByIdTeamClone`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

payload := groupteam.CreateGroupByIdTeamCloneRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamClone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.CreateGroupByIdTeamCompleteMigration`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdTeamCompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.CreateGroupByIdTeamSendActivityNotification`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

payload := groupteam.CreateGroupByIdTeamSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.CreateGroupByIdTeamUnarchive`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

read, err := client.CreateGroupByIdTeamUnarchive(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.CreateUpdateGroupByIdTeam`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

payload := groupteam.Team{
	// ...
}


read, err := client.CreateUpdateGroupByIdTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.DeleteGroupByIdTeam`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

read, err := client.DeleteGroupByIdTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamClient.GetGroupByIdTeam`

```go
ctx := context.TODO()
id := groupteam.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeam(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
