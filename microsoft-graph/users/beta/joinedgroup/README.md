
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/joinedgroup` Documentation

The `joinedgroup` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/joinedgroup"
```


### Client Initialization

```go
client := joinedgroup.NewJoinedGroupClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedGroupClient.EvaluateJoinedGroupsDynamicMembership`

```go
ctx := context.TODO()
id := joinedgroup.NewUserID("userId")

payload := joinedgroup.EvaluateJoinedGroupsDynamicMembershipRequest{
	// ...
}


read, err := client.EvaluateJoinedGroupsDynamicMembership(ctx, id, payload, joinedgroup.DefaultEvaluateJoinedGroupsDynamicMembershipOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedGroupClient.ListJoinedGroups`

```go
ctx := context.TODO()
id := joinedgroup.NewUserID("userId")

// alternatively `client.ListJoinedGroups(ctx, id, joinedgroup.DefaultListJoinedGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedGroupsComplete(ctx, id, joinedgroup.DefaultListJoinedGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
