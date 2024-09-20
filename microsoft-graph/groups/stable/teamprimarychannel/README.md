
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannel` Documentation

The `teamprimarychannel` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannel"
```


### Client Initialization

```go
client := teamprimarychannel.NewTeamPrimaryChannelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelClient.CreateTeamPrimaryChannelArchive`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

payload := teamprimarychannel.CreateTeamPrimaryChannelArchiveRequest{
	// ...
}


read, err := client.CreateTeamPrimaryChannelArchive(ctx, id, payload, teamprimarychannel.DefaultCreateTeamPrimaryChannelArchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.CreateTeamPrimaryChannelCompleteMigration`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.CreateTeamPrimaryChannelCompleteMigration(ctx, id, teamprimarychannel.DefaultCreateTeamPrimaryChannelCompleteMigrationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.CreateTeamPrimaryChannelUnarchive`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.CreateTeamPrimaryChannelUnarchive(ctx, id, teamprimarychannel.DefaultCreateTeamPrimaryChannelUnarchiveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.DeleteTeamPrimaryChannel`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.DeleteTeamPrimaryChannel(ctx, id, teamprimarychannel.DefaultDeleteTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.GetTeamPrimaryChannel`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.GetTeamPrimaryChannel(ctx, id, teamprimarychannel.DefaultGetTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.ProvisionTeamPrimaryChannelEmail`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.ProvisionTeamPrimaryChannelEmail(ctx, id, teamprimarychannel.DefaultProvisionTeamPrimaryChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.RemoveTeamPrimaryChannelEmail`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

read, err := client.RemoveTeamPrimaryChannelEmail(ctx, id, teamprimarychannel.DefaultRemoveTeamPrimaryChannelEmailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelClient.UpdateTeamPrimaryChannel`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupId")

payload := teamprimarychannel.Channel{
	// ...
}


read, err := client.UpdateTeamPrimaryChannel(ctx, id, payload, teamprimarychannel.DefaultUpdateTeamPrimaryChannelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
