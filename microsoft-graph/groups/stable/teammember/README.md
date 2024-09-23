
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teammember` Documentation

The `teammember` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teammember"
```


### Client Initialization

```go
client := teammember.NewTeamMemberClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamMemberClient.AddTeamMembers`

```go
ctx := context.TODO()
id := teammember.NewGroupID("groupId")

payload := teammember.AddTeamMembersRequest{
	// ...
}


// alternatively `client.AddTeamMembers(ctx, id, payload, teammember.DefaultAddTeamMembersOperationOptions())` can be used to do batched pagination
items, err := client.AddTeamMembersComplete(ctx, id, payload, teammember.DefaultAddTeamMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamMemberClient.CreateTeamMember`

```go
ctx := context.TODO()
id := teammember.NewGroupID("groupId")

payload := teammember.ConversationMember{
	// ...
}


read, err := client.CreateTeamMember(ctx, id, payload, teammember.DefaultCreateTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamMemberClient.DeleteTeamMember`

```go
ctx := context.TODO()
id := teammember.NewGroupIdTeamMemberID("groupId", "conversationMemberId")

read, err := client.DeleteTeamMember(ctx, id, teammember.DefaultDeleteTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamMemberClient.GetTeamMember`

```go
ctx := context.TODO()
id := teammember.NewGroupIdTeamMemberID("groupId", "conversationMemberId")

read, err := client.GetTeamMember(ctx, id, teammember.DefaultGetTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamMemberClient.GetTeamMembersCount`

```go
ctx := context.TODO()
id := teammember.NewGroupID("groupId")

read, err := client.GetTeamMembersCount(ctx, id, teammember.DefaultGetTeamMembersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamMemberClient.ListTeamMembers`

```go
ctx := context.TODO()
id := teammember.NewGroupID("groupId")

// alternatively `client.ListTeamMembers(ctx, id, teammember.DefaultListTeamMembersOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamMembersComplete(ctx, id, teammember.DefaultListTeamMembersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamMemberClient.UpdateTeamMember`

```go
ctx := context.TODO()
id := teammember.NewGroupIdTeamMemberID("groupId", "conversationMemberId")

payload := teammember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamMember(ctx, id, payload, teammember.DefaultUpdateTeamMemberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
