
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannel` Documentation

The `teamprimarychannel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannel"
```


### Client Initialization

```go
client := teamprimarychannel.NewTeamPrimaryChannelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelClient.CreateTeamPrimaryChannelArchive`

```go
ctx := context.TODO()
id := teamprimarychannel.NewGroupID("groupIdValue")

payload := teamprimarychannel.CreateTeamPrimaryChannelArchiveRequest{
	// ...
}


read, err := client.CreateTeamPrimaryChannelArchive(ctx, id, payload)
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
id := teamprimarychannel.NewGroupID("groupIdValue")

read, err := client.CreateTeamPrimaryChannelCompleteMigration(ctx, id)
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
id := teamprimarychannel.NewGroupID("groupIdValue")

read, err := client.CreateTeamPrimaryChannelUnarchive(ctx, id)
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
id := teamprimarychannel.NewGroupID("groupIdValue")

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
id := teamprimarychannel.NewGroupID("groupIdValue")

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
id := teamprimarychannel.NewGroupID("groupIdValue")

read, err := client.ProvisionTeamPrimaryChannelEmail(ctx, id)
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
id := teamprimarychannel.NewGroupID("groupIdValue")

read, err := client.RemoveTeamPrimaryChannelEmail(ctx, id)
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
id := teamprimarychannel.NewGroupID("groupIdValue")

payload := teamprimarychannel.Channel{
	// ...
}


read, err := client.UpdateTeamPrimaryChannel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
