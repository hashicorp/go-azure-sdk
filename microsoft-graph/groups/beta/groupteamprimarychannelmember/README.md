
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamprimarychannelmember` Documentation

The `groupteamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamprimarychannelmember"
```


### Client Initialization

```go
client := groupteamprimarychannelmember.NewGroupTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.AddGroupByIdTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupID("groupIdValue")

payload := groupteamprimarychannelmember.AddGroupByIdTeamPrimaryChannelMemberRequest{
	// ...
}


read, err := client.AddGroupByIdTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.CreateGroupByIdTeamPrimaryChannelMember`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupID("groupIdValue")

payload := groupteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateGroupByIdTeamPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.DeleteGroupByIdTeamPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.DeleteGroupByIdTeamPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.GetGroupByIdTeamPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.GetGroupByIdTeamPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.GetGroupByIdTeamPrimaryChannelMemberCount`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamPrimaryChannelMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.ListGroupByIdTeamPrimaryChannelMembers`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamPrimaryChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamPrimaryChannelMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamPrimaryChannelMemberClient.UpdateGroupByIdTeamPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := groupteamprimarychannelmember.NewGroupTeamPrimaryChannelMemberID("groupIdValue", "conversationMemberIdValue")

payload := groupteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateGroupByIdTeamPrimaryChannelMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
