
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteam` Documentation

The `mejoinedteam` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteam"
```


### Client Initialization

```go
client := mejoinedteam.NewMeJoinedTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeam`

```go
ctx := context.TODO()

payload := mejoinedteam.Team{
	// ...
}


read, err := client.CreateMeJoinedTeam(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeamByIdArchive`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteam.CreateMeJoinedTeamByIdArchiveRequest{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdArchive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeamByIdClone`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteam.CreateMeJoinedTeamByIdCloneRequest{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdClone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeamByIdCompleteMigration`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.CreateMeJoinedTeamByIdCompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeamByIdSendActivityNotification`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteam.CreateMeJoinedTeamByIdSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.CreateMeJoinedTeamByIdUnarchive`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.CreateMeJoinedTeamByIdUnarchive(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.DeleteMeJoinedTeamById`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.DeleteMeJoinedTeamById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.GetMeJoinedTeamById`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.GetMeJoinedTeamCount`

```go
ctx := context.TODO()


read, err := client.GetMeJoinedTeamCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamClient.ListMeJoinedTeams`

```go
ctx := context.TODO()


// alternatively `client.ListMeJoinedTeams(ctx)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamClient.UpdateMeJoinedTeamById`

```go
ctx := context.TODO()
id := mejoinedteam.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteam.Team{
	// ...
}


read, err := client.UpdateMeJoinedTeamById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
