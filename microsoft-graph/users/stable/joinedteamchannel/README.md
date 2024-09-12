
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannel` Documentation

The `joinedteamchannel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannel"
```


### Client Initialization

```go
client := joinedteamchannel.NewJoinedTeamChannelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamChannelClient.CreateJoinedTeamChannel`

```go
ctx := context.TODO()
id := joinedteamchannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamchannel.Channel{
	// ...
}


read, err := client.CreateJoinedTeamChannel(ctx, id, payload)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := joinedteamchannel.CreateJoinedTeamChannelArchiveRequest{
	// ...
}


read, err := client.CreateJoinedTeamChannelArchive(ctx, id, payload)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.CreateJoinedTeamChannelCompleteMigration(ctx, id)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.CreateJoinedTeamChannelUnarchive(ctx, id)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

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
id := joinedteamchannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

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
id := joinedteamchannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.ProvisionJoinedTeamChannelEmail(ctx, id)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

read, err := client.RemoveJoinedTeamChannelEmail(ctx, id)
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
id := joinedteamchannel.NewUserIdJoinedTeamIdChannelID("userIdValue", "teamIdValue", "channelIdValue")

payload := joinedteamchannel.Channel{
	// ...
}


read, err := client.UpdateJoinedTeamChannel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
