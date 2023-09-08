
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteammember` Documentation

The `groupteammember` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteammember"
```


### Client Initialization

```go
client := groupteammember.NewGroupTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamMemberClient.AddGroupByIdTeamMember`

```go
ctx := context.TODO()
id := groupteammember.NewGroupID("groupIdValue")

payload := groupteammember.AddGroupByIdTeamMemberRequest{
	// ...
}


read, err := client.AddGroupByIdTeamMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamMemberClient.CreateGroupByIdTeamMember`

```go
ctx := context.TODO()
id := groupteammember.NewGroupID("groupIdValue")

payload := groupteammember.ConversationMember{
	// ...
}


read, err := client.CreateGroupByIdTeamMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamMemberClient.DeleteGroupByIdTeamMemberById`

```go
ctx := context.TODO()
id := groupteammember.NewGroupTeamMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.DeleteGroupByIdTeamMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamMemberClient.GetGroupByIdTeamMemberById`

```go
ctx := context.TODO()
id := groupteammember.NewGroupTeamMemberID("groupIdValue", "conversationMemberIdValue")

read, err := client.GetGroupByIdTeamMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamMemberClient.GetGroupByIdTeamMemberCount`

```go
ctx := context.TODO()
id := groupteammember.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamMemberClient.ListGroupByIdTeamMembers`

```go
ctx := context.TODO()
id := groupteammember.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamMemberClient.UpdateGroupByIdTeamMemberById`

```go
ctx := context.TODO()
id := groupteammember.NewGroupTeamMemberID("groupIdValue", "conversationMemberIdValue")

payload := groupteammember.ConversationMember{
	// ...
}


read, err := client.UpdateGroupByIdTeamMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
