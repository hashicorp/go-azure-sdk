
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmember` Documentation

The `teamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmember"
```


### Client Initialization

```go
client := teamprimarychannelmember.NewTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelMemberClient.AddTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

payload := teamprimarychannelmember.AddTeamPrimaryChannelMembersRequest{
	// ...
}


// alternatively `client.AddTeamPrimaryChannelMembers(ctx, id, payload, teamprimarychannelmember.DefaultAddTeamPrimaryChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddTeamPrimaryChannelMembersComplete(ctx, id, payload, teamprimarychannelmember.DefaultAddTeamPrimaryChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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

read, err := client.DeleteTeamPrimaryChannelMember(ctx, id, teamprimarychannelmember.DefaultDeleteTeamPrimaryChannelMemberOperationOptions())
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

read, err := client.GetTeamPrimaryChannelMember(ctx, id, teamprimarychannelmember.DefaultGetTeamPrimaryChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelMemberClient.GetTeamPrimaryChannelMembersCount`

```go
ctx := context.TODO()
id := teamprimarychannelmember.NewGroupID("groupIdValue")

read, err := client.GetTeamPrimaryChannelMembersCount(ctx, id, teamprimarychannelmember.DefaultGetTeamPrimaryChannelMembersCountOperationOptions())
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

// alternatively `client.ListTeamPrimaryChannelMembers(ctx, id, teamprimarychannelmember.DefaultListTeamPrimaryChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelMembersComplete(ctx, id, teamprimarychannelmember.DefaultListTeamPrimaryChannelMembersOperationOptions())
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
