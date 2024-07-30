
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmember` Documentation

The `joinedteamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmember"
```


### Client Initialization

```go
client := joinedteamprimarychannelmember.NewJoinedTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.AddUserJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamprimarychannelmember.AddUserJoinedTeamPrimaryChannelMemberRequest{
	// ...
}


read, err := client.AddUserJoinedTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.CreateJoinedTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamPrimaryChannelMember(ctx, id, payload)
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
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteJoinedTeamPrimaryChannelMember(ctx, id)
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
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.GetJoinedTeamPrimaryChannelMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamPrimaryChannelMemberClient.GetJoinedTeamPrimaryChannelMemberCount`

```go
ctx := context.TODO()
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeamPrimaryChannelMemberCount(ctx, id)
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
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListJoinedTeamPrimaryChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamPrimaryChannelMembersComplete(ctx, id)
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
id := joinedteamprimarychannelmember.NewUserIdJoinedTeamIdPrimaryChannelMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

payload := joinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
