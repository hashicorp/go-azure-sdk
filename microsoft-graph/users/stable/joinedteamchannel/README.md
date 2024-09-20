
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannel` Documentation

The `joinedteamchannel` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannel"
```


### Client Initialization

```go
client := joinedteamchannel.NewJoinedTeamChannelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelClient.CreateJoinedTeamChannel`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamchannel.Channel{
	// ...
}


read, err := client.CreateJoinedTeamChannel(ctx, id, payload, joinedteamchannel.DefaultCreateJoinedTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.CreateJoinedTeamChannelArchive`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannel.CreateJoinedTeamChannelArchiveRequest{
	// ...
}


read, err := client.CreateJoinedTeamChannelArchive(ctx, id, payload, joinedteamchannel.DefaultCreateJoinedTeamChannelArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.CreateJoinedTeamChannelCompleteMigration`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.CreateJoinedTeamChannelCompleteMigration(ctx, id, joinedteamchannel.DefaultCreateJoinedTeamChannelCompleteMigrationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.CreateJoinedTeamChannelUnarchive`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.CreateJoinedTeamChannelUnarchive(ctx, id, joinedteamchannel.DefaultCreateJoinedTeamChannelUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.DeleteJoinedTeamChannel`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.DeleteJoinedTeamChannel(ctx, id, joinedteamchannel.DefaultDeleteJoinedTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.GetJoinedTeamChannel`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.GetJoinedTeamChannel(ctx, id, joinedteamchannel.DefaultGetJoinedTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.GetJoinedTeamChannelsCount`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamChannelsCount(ctx, id, joinedteamchannel.DefaultGetJoinedTeamChannelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.ListJoinedTeamChannels`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamID("userId", "teamId")

// alternatively `client.ListJoinedTeamChannels(ctx, id, joinedteamchannel.DefaultListJoinedTeamChannelsOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamChannelsComplete(ctx, id, joinedteamchannel.DefaultListJoinedTeamChannelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamChannelClient.ProvisionJoinedTeamChannelEmail`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.ProvisionJoinedTeamChannelEmail(ctx, id, joinedteamchannel.DefaultProvisionJoinedTeamChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.RemoveJoinedTeamChannelEmail`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

read, err := client.RemoveJoinedTeamChannelEmail(ctx, id, joinedteamchannel.DefaultRemoveJoinedTeamChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamChannelClient.UpdateJoinedTeamChannel`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userId", "teamId", "channelId")

payload := joinedteamchannel.Channel{
	// ...
}


read, err := client.UpdateJoinedTeamChannel(ctx, id, payload, joinedteamchannel.DefaultUpdateJoinedTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
