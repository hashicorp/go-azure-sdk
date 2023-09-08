
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteam` Documentation

The `userjoinedteam` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/userjoinedteam"
```


### Client Initialization

```go
client := userjoinedteam.NewUserJoinedTeamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeam`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserID("userIdValue")

payload := userjoinedteam.Team{
	// ...
}


read, err := client.CreateUserByIdJoinedTeam(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeamByIdArchive`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteam.CreateUserByIdJoinedTeamByIdArchiveRequest{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdArchive(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeamByIdClone`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteam.CreateUserByIdJoinedTeamByIdCloneRequest{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdClone(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeamByIdCompleteMigration`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdCompleteMigration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeamByIdSendActivityNotification`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteam.CreateUserByIdJoinedTeamByIdSendActivityNotificationRequest{
	// ...
}


read, err := client.CreateUserByIdJoinedTeamByIdSendActivityNotification(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.CreateUserByIdJoinedTeamByIdUnarchive`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateUserByIdJoinedTeamByIdUnarchive(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.DeleteUserByIdJoinedTeamById`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.DeleteUserByIdJoinedTeamById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.GetUserByIdJoinedTeamById`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetUserByIdJoinedTeamById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.GetUserByIdJoinedTeamCount`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserID("userIdValue")

read, err := client.GetUserByIdJoinedTeamCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserJoinedTeamClient.ListUserByIdJoinedTeams`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserID("userIdValue")

// alternatively `client.ListUserByIdJoinedTeams(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdJoinedTeamsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserJoinedTeamClient.UpdateUserByIdJoinedTeamById`

```go
ctx := context.TODO()
id := userjoinedteam.NewUserJoinedTeamID("userIdValue", "teamIdValue")

payload := userjoinedteam.Team{
	// ...
}


read, err := client.UpdateUserByIdJoinedTeamById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
