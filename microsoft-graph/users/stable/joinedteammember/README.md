
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteammember` Documentation

The `joinedteammember` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteammember"
```


### Client Initialization

```go
client := joinedteammember.NewJoinedTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamMemberClient.AddUserJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteammember.AddUserJoinedTeamMemberRequest{
	// ...
}


read, err := client.AddUserJoinedTeamMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.CreateJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

payload := joinedteammember.ConversationMember{
	// ...
}


read, err := client.CreateJoinedTeamMember(ctx, id, payload)
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
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteJoinedTeamMember(ctx, id)
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
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

read, err := client.GetJoinedTeamMember(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamMemberClient.GetJoinedTeamMemberCount`

```go
ctx := context.TODO()
id := joinedteammember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

read, err := client.GetJoinedTeamMemberCount(ctx, id)
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
id := joinedteammember.NewUserIdJoinedTeamID("userIdValue", "teamIdValue")

// alternatively `client.ListJoinedTeamMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListJoinedTeamMembersComplete(ctx, id)
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
id := joinedteammember.NewUserIdJoinedTeamIdMemberID("userIdValue", "teamIdValue", "conversationMemberIdValue")

payload := joinedteammember.ConversationMember{
	// ...
}


read, err := client.UpdateJoinedTeamMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
