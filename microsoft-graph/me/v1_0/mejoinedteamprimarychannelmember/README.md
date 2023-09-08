
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmember` Documentation

The `mejoinedteamprimarychannelmember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteamprimarychannelmember"
```


### Client Initialization

```go
client := mejoinedteamprimarychannelmember.NewMeJoinedTeamPrimaryChannelMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.AddMeJoinedTeamByIdPrimaryChannelMember`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteamprimarychannelmember.AddMeJoinedTeamByIdPrimaryChannelMemberRequest{
	// ...
}


read, err := client.AddMeJoinedTeamByIdPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.CreateMeJoinedTeamByIdPrimaryChannelMember`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdPrimaryChannelMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.DeleteMeJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamPrimaryChannelMemberID("teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteMeJoinedTeamByIdPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.GetMeJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamPrimaryChannelMemberID("teamIdValue", "conversationMemberIdValue")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.GetMeJoinedTeamByIdPrimaryChannelMemberCount`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamByIdPrimaryChannelMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.ListMeJoinedTeamByIdPrimaryChannelMembers`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.ListMeJoinedTeamByIdPrimaryChannelMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdPrimaryChannelMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPrimaryChannelMemberClient.UpdateMeJoinedTeamByIdPrimaryChannelMemberById`

```go
ctx := context.TODO()
id := mejoinedteamprimarychannelmember.NewMeJoinedTeamPrimaryChannelMemberID("teamIdValue", "conversationMemberIdValue")

payload := mejoinedteamprimarychannelmember.ConversationMember{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdPrimaryChannelMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
