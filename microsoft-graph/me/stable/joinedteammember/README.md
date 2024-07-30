
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteammember` Documentation

The `joinedteammember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteammember"
```


### Client Initialization

```go
client := joinedteammember.NewJoinedTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamMemberClient.AddMeJoinedTeamMember`

```go
ctx := context.TODO()
id := joinedteammember.NewMeJoinedTeamID("teamIdValue")

payload := joinedteammember.AddMeJoinedTeamMemberRequest{
	// ...
}


read, err := client.AddMeJoinedTeamMember(ctx, id, payload)
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
id := joinedteammember.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteammember.NewMeJoinedTeamIdMemberID("teamIdValue", "conversationMemberIdValue")

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
id := joinedteammember.NewMeJoinedTeamIdMemberID("teamIdValue", "conversationMemberIdValue")

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
id := joinedteammember.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteammember.NewMeJoinedTeamID("teamIdValue")

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
id := joinedteammember.NewMeJoinedTeamIdMemberID("teamIdValue", "conversationMemberIdValue")

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
