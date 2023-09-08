
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteammember` Documentation

The `mejoinedteammember` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteammember"
```


### Client Initialization

```go
client := mejoinedteammember.NewMeJoinedTeamMemberClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamMemberClient.AddMeJoinedTeamByIdMember`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteammember.AddMeJoinedTeamByIdMemberRequest{
	// ...
}


read, err := client.AddMeJoinedTeamByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamMemberClient.CreateMeJoinedTeamByIdMember`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteammember.ConversationMember{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdMember(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamMemberClient.DeleteMeJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamMemberID("teamIdValue", "conversationMemberIdValue")

read, err := client.DeleteMeJoinedTeamByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamMemberClient.GetMeJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamMemberID("teamIdValue", "conversationMemberIdValue")

read, err := client.GetMeJoinedTeamByIdMemberById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamMemberClient.GetMeJoinedTeamByIdMemberCount`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamByIdMemberCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamMemberClient.ListMeJoinedTeamByIdMembers`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.ListMeJoinedTeamByIdMembers(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdMembersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamMemberClient.UpdateMeJoinedTeamByIdMemberById`

```go
ctx := context.TODO()
id := mejoinedteammember.NewMeJoinedTeamMemberID("teamIdValue", "conversationMemberIdValue")

payload := mejoinedteammember.ConversationMember{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdMemberById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
