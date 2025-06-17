
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelallmember` Documentation

The `joinedteamprimarychannelallmember` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelallmember"
```


### Client Initialization

```go
client := joinedteamprimarychannelallmember.NewJoinedTeamPrimaryChannelAllMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.AddJoinedTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelallmember.AddJoinedTeamPrimaryChannelAllMembersRequest{
	// ...
}


// alternatively `client.AddJoinedTeamPrimaryChannelAllMembers(ctx, id, payload, joinedteamprimarychannelallmember.DefaultAddJoinedTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddJoinedTeamPrimaryChannelAllMembersComplete(ctx, id, payload, joinedteamprimarychannelallmember.DefaultAddJoinedTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.CreateJoinedTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelallmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelAllMember(ctx, id, payload, joinedteamprimarychannelallmember.DefaultCreateJoinedTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.DeleteJoinedTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamIdPrimaryChannelAllMemberID("userId", "teamId", "conversationMemberId")

read, err := client.DeleteJoinedTeamPrimaryChannelAllMember(ctx, id, joinedteamprimarychannelallmember.DefaultDeleteJoinedTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.GetJoinedTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamIdPrimaryChannelAllMemberID("userId", "teamId", "conversationMemberId")

read, err := client.GetJoinedTeamPrimaryChannelAllMember(ctx, id, joinedteamprimarychannelallmember.DefaultGetJoinedTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.GetJoinedTeamPrimaryChannelAllMembersCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamPrimaryChannelAllMembersCount(ctx, id, joinedteamprimarychannelallmember.DefaultGetJoinedTeamPrimaryChannelAllMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.ListJoinedTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamID("userId", "teamId")

// alternatively `client.ListJoinedTeamPrimaryChannelAllMembers(ctx, id, joinedteamprimarychannelallmember.DefaultListJoinedTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelAllMembersComplete(ctx, id, joinedteamprimarychannelallmember.DefaultListJoinedTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.RemoveJoinedTeamPrimaryChannelAllMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelallmember.RemoveJoinedTeamPrimaryChannelAllMembersRequest{
	// ...
}


// alternatively `client.RemoveJoinedTeamPrimaryChannelAllMembers(ctx, id, payload, joinedteamprimarychannelallmember.DefaultRemoveJoinedTeamPrimaryChannelAllMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveJoinedTeamPrimaryChannelAllMembersComplete(ctx, id, payload, joinedteamprimarychannelallmember.DefaultRemoveJoinedTeamPrimaryChannelAllMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelAllMemberClient.UpdateJoinedTeamPrimaryChannelAllMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelallmember.NewUserIdJoinedTeamIdPrimaryChannelAllMemberID("userId", "teamId", "conversationMemberId")

payload := joinedteamprimarychannelallmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelAllMember(ctx, id, payload, joinedteamprimarychannelallmember.DefaultUpdateJoinedTeamPrimaryChannelAllMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
