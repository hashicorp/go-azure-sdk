
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmember` Documentation

The `joinedteamprimarychannelmember` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmember"
```


### Client Initialization

```go
client := joinedteamprimarychannelmember.NewJoinedTeamPrimaryChannelMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.AddJoinedTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelmember.AddJoinedTeamPrimaryChannelMembersRequest{
	// ...
}


// alternatively `client.AddJoinedTeamPrimaryChannelMembers(ctx, id, payload, joinedteamprimarychannelmember.DefaultAddJoinedTeamPrimaryChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddJoinedTeamPrimaryChannelMembersComplete(ctx, id, payload, joinedteamprimarychannelmember.DefaultAddJoinedTeamPrimaryChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.CreateJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMember(ctx, id, payload, joinedteamprimarychannelmember.DefaultCreateJoinedTeamPrimaryChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.DeleteJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userId", "teamId", "conversationMemberId")

read, err := client.DeleteJoinedTeamPrimaryChannelMember(ctx, id, joinedteamprimarychannelmember.DefaultDeleteJoinedTeamPrimaryChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.GetJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userId", "teamId", "conversationMemberId")

read, err := client.GetJoinedTeamPrimaryChannelMember(ctx, id, joinedteamprimarychannelmember.DefaultGetJoinedTeamPrimaryChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.GetJoinedTeamPrimaryChannelMembersCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamPrimaryChannelMembersCount(ctx, id, joinedteamprimarychannelmember.DefaultGetJoinedTeamPrimaryChannelMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.ListJoinedTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userId", "teamId")

// alternatively `client.ListJoinedTeamPrimaryChannelMembers(ctx, id, joinedteamprimarychannelmember.DefaultListJoinedTeamPrimaryChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMembersComplete(ctx, id, joinedteamprimarychannelmember.DefaultListJoinedTeamPrimaryChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.RemoveJoinedTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteamprimarychannelmember.RemoveJoinedTeamPrimaryChannelMembersRequest{
	// ...
}


// alternatively `client.RemoveJoinedTeamPrimaryChannelMembers(ctx, id, payload, joinedteamprimarychannelmember.DefaultRemoveJoinedTeamPrimaryChannelMembersOperationOptions())` can be used to do batched pagination
items, err := client.RemoveJoinedTeamPrimaryChannelMembersComplete(ctx, id, payload, joinedteamprimarychannelmember.DefaultRemoveJoinedTeamPrimaryChannelMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.UpdateJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userId", "teamId", "conversationMemberId")

payload := joinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMember(ctx, id, payload, joinedteamprimarychannelmember.DefaultUpdateJoinedTeamPrimaryChannelMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
