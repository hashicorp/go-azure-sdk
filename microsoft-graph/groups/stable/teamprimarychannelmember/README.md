
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmember` Documentation

The `teamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmember"
```


### Client Initialization

```go
client := teamprimarychannelmember.NewTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelMemberClient.AddGroupTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

payload := teamprimarychannelmember.AddGroupTeamPrimaryChannelMemberRequest{
	// ...
}


read, err := client.AddGroupTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.CreateTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

payload := teamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.DeleteTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupIdTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.DeleteTeamPrimaryChannelMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.GetTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupIdTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.GetTeamPrimaryChannelMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.GetTeamPrimaryChannelMemberCount`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

read, err := client.GetTeamPrimaryChannelMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.ListTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

// alternatively `client.ListTeamPrimaryChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.UpdateTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupIdTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

payload := teamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
