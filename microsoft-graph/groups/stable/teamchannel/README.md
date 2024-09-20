
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannel` Documentation

The `teamchannel` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannel"
```


### Client Initialization

```go
client := teamchannel.NewTeamChannelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamChannelClient.CreateTeamChannel`

```go
ctx := context.TODO()
id := teamchannel.NewGroupID("groupId")

payload := teamchannel.Channel{
	// ...
}


read, err := client.CreateTeamChannel(ctx, id, payload, teamchannel.DefaultCreateTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.CreateTeamChannelArchive`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannel.CreateTeamChannelArchiveRequest{
	// ...
}


read, err := client.CreateTeamChannelArchive(ctx, id, payload, teamchannel.DefaultCreateTeamChannelArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.CreateTeamChannelCompleteMigration`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.CreateTeamChannelCompleteMigration(ctx, id, teamchannel.DefaultCreateTeamChannelCompleteMigrationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.CreateTeamChannelUnarchive`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.CreateTeamChannelUnarchive(ctx, id, teamchannel.DefaultCreateTeamChannelUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.DeleteTeamChannel`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.DeleteTeamChannel(ctx, id, teamchannel.DefaultDeleteTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.GetTeamChannel`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.GetTeamChannel(ctx, id, teamchannel.DefaultGetTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.GetTeamChannelsCount`

```go
ctx := context.TODO()
id := teamchannel.NewGroupID("groupId")

read, err := client.GetTeamChannelsCount(ctx, id, teamchannel.DefaultGetTeamChannelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.ListTeamChannels`

```go
ctx := context.TODO()
id := teamchannel.NewGroupID("groupId")

// alternatively `client.ListTeamChannels(ctx, id, teamchannel.DefaultListTeamChannelsOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamChannelsComplete(ctx, id, teamchannel.DefaultListTeamChannelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamChannelClient.ProvisionTeamChannelEmail`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.ProvisionTeamChannelEmail(ctx, id, teamchannel.DefaultProvisionTeamChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.RemoveTeamChannelEmail`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

read, err := client.RemoveTeamChannelEmail(ctx, id, teamchannel.DefaultRemoveTeamChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamChannelClient.UpdateTeamChannel`

```go
ctx := context.TODO()
id := teamchannel.NewGroupIdTeamChannelID("groupId", "channelId")

payload := teamchannel.Channel{
	// ...
}


read, err := client.UpdateTeamChannel(ctx, id, payload, teamchannel.DefaultUpdateTeamChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
