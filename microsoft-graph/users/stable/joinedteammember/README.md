
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteammember` Documentation

The `joinedteammember` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteammember"
```


### Client Initialization

```go
client := joinedteammember.NewJoinedTeamMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamMemberClient.AddJoinedTeamMembers`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteammember.AddJoinedTeamMembersRequest{
	// ...
}


// alternatively `client.AddJoinedTeamMembers(ctx, id, payload, joinedteammember.DefaultAddJoinedTeamMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddJoinedTeamMembersComplete(ctx, id, payload, joinedteammember.DefaultAddJoinedTeamMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamMemberClient.CreateJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userId", "teamId")

payload := joinedteammember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamMember(ctx, id, payload, joinedteammember.DefaultCreateJoinedTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.DeleteJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userId", "teamId", "conversationMemberId")

read, err := client.DeleteJoinedTeamMember(ctx, id, joinedteammember.DefaultDeleteJoinedTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.GetJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userId", "teamId", "conversationMemberId")

read, err := client.GetJoinedTeamMember(ctx, id, joinedteammember.DefaultGetJoinedTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.GetJoinedTeamMembersCount`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userId", "teamId")

read, err := client.GetJoinedTeamMembersCount(ctx, id, joinedteammember.DefaultGetJoinedTeamMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.ListJoinedTeamMembers`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userId", "teamId")

// alternatively `client.ListJoinedTeamMembers(ctx, id, joinedteammember.DefaultListJoinedTeamMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamMembersComplete(ctx, id, joinedteammember.DefaultListJoinedTeamMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamMemberClient.UpdateJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userId", "teamId", "conversationMemberId")

payload := joinedteammember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamMember(ctx, id, payload, joinedteammember.DefaultUpdateJoinedTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
