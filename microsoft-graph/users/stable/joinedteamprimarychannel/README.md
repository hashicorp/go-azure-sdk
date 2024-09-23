
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannel` Documentation

The `joinedteamprimarychannel` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannel"
```


### Client Initialization

```go
client := joinedteamprimarychannel.NewJoinedTeamPrimaryChannelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelClient.CreateJoinedTeamPrimaryChannelArchive`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannel.CreateJoinedTeamPrimaryChannelArchiveRequest{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelArchive(ctx, id, payload, joinedteamprimarychannel.DefaultCreateJoinedTeamPrimaryChannelArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.CreateJoinedTeamPrimaryChannelCompleteMigration`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.CreateJoinedTeamPrimaryChannelCompleteMigration(ctx, id, joinedteamprimarychannel.DefaultCreateJoinedTeamPrimaryChannelCompleteMigrationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.CreateJoinedTeamPrimaryChannelUnarchive`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.CreateJoinedTeamPrimaryChannelUnarchive(ctx, id, joinedteamprimarychannel.DefaultCreateJoinedTeamPrimaryChannelUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.DeleteJoinedTeamPrimaryChannel`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.DeleteJoinedTeamPrimaryChannel(ctx, id, joinedteamprimarychannel.DefaultDeleteJoinedTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.GetJoinedTeamPrimaryChannel`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamPrimaryChannel(ctx, id, joinedteamprimarychannel.DefaultGetJoinedTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.ProvisionJoinedTeamPrimaryChannelEmail`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.ProvisionJoinedTeamPrimaryChannelEmail(ctx, id, joinedteamprimarychannel.DefaultProvisionJoinedTeamPrimaryChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.RemoveJoinedTeamPrimaryChannelEmail`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.RemoveJoinedTeamPrimaryChannelEmail(ctx, id, joinedteamprimarychannel.DefaultRemoveJoinedTeamPrimaryChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelClient.UpdateJoinedTeamPrimaryChannel`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannel.Channel{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannel(ctx, id, payload, joinedteamprimarychannel.DefaultUpdateJoinedTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
