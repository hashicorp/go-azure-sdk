
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teammember` Documentation

The `teammember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teammember"
```


### Client Initialization

```go
client := teammember.NewTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamMemberClient.AddTeamMembers`

```go
ctx := context.TODO()
id := teammember.NewGroupID("groupIdValue")

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
id := teammember.NewGroupID("groupIdValue")

payload := teammember.ConversationMember{
	// ...
}


read, err := client.CreateTeamMember(ctx, id, payload)
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
id := teammember.NewGroupIdTeamMemberID("groupIdValue", "conversationMemberIdValue")

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
id := teammember.NewGroupIdTeamMemberID("groupIdValue", "conversationMemberIdValue")

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
id := teammember.NewGroupID("groupIdValue")

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
id := teammember.NewGroupID("groupIdValue")

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
id := teammember.NewGroupIdTeamMemberID("groupIdValue", "conversationMemberIdValue")

payload := teammember.ConversationMember{
	// ...
}


read, err := client.UpdateTeamMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
