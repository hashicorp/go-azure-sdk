
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannel` Documentation

The `joinedteamprimarychannel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannel"
```


### Client Initialization

```go
client := joinedteamprimarychannel.NewJoinedTeamPrimaryChannelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelClient.CreateJoinedTeamPrimaryChannelArchive`

```go
ctx := context.TODO()
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamprimarychannel.CreateJoinedTeamPrimaryChannelArchiveRequest{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelArchive(ctx, id, payload)
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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateJoinedTeamPrimaryChannelCompleteMigration(ctx, id)
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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.CreateJoinedTeamPrimaryChannelUnarchive(ctx, id)
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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.ProvisionJoinedTeamPrimaryChannelEmail(ctx, id)
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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.RemoveJoinedTeamPrimaryChannelEmail(ctx, id)
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
id := joinedteamprimarychannel.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamprimarychannel.Channel{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
