
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelallmember` Documentation

The `teamprimarychannelallmember` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelallmember"
```


### Client Initialization

```go
client := teamprimarychannelallmember.NewTeamPrimaryChannelAllMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.AddTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupID("groupId")

payload := teamprimarychannelallmember.AddTeamPrimaryChannelAllMembersRequest{
	// ...
}


// alternatively `client.AddTeamPrimaryChannelAllMembers(ctx, id, payload, teamprimarychannelallmember.DefaultAddTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddTeamPrimaryChannelAllMembersComplete(ctx, id, payload, teamprimarychannelallmember.DefaultAddTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.CreateTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupID("groupId")

payload := teamprimarychannelallmember.ConversationMember{
	// ...
}


read, err := client.CreateTeamPrimaryChannelAllMember(ctx, id, payload, teamprimarychannelallmember.DefaultCreateTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.DeleteTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupIdTeamPrimaryChannelAllMemberID("groupId", "conversationMemberId")

read, err := client.DeleteTeamPrimaryChannelAllMember(ctx, id, teamprimarychannelallmember.DefaultDeleteTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.GetTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupIdTeamPrimaryChannelAllMemberID("groupId", "conversationMemberId")

read, err := client.GetTeamPrimaryChannelAllMember(ctx, id, teamprimarychannelallmember.DefaultGetTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.GetTeamPrimaryChannelAllMembersCount`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupID("groupId")

read, err := client.GetTeamPrimaryChannelAllMembersCount(ctx, id, teamprimarychannelallmember.DefaultGetTeamPrimaryChannelAllMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.ListTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupID("groupId")

// alternatively `client.ListTeamPrimaryChannelAllMembers(ctx, id, teamprimarychannelallmember.DefaultListTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamPrimaryChannelAllMembersComplete(ctx, id, teamprimarychannelallmember.DefaultListTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.RemoveTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupID("groupId")

payload := teamprimarychannelallmember.RemoveTeamPrimaryChannelAllMembersRequest{
	// ...
}


// alternatively `client.RemoveTeamPrimaryChannelAllMembers(ctx, id, payload, teamprimarychannelallmember.DefaultRemoveTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveTeamPrimaryChannelAllMembersComplete(ctx, id, payload, teamprimarychannelallmember.DefaultRemoveTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamPrimaryChannelAllMemberClient.UpdateTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := teamprimarychannelallmember.NewGroupIdTeamPrimaryChannelAllMemberID("groupId", "conversationMemberId")

payload := teamprimarychannelallmember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamPrimaryChannelAllMember(ctx, id, payload, teamprimarychannelallmember.DefaultUpdateTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
