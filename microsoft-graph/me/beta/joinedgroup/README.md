
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/joinedgroup` Documentation

The `joinedgroup` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/joinedgroup"
```


### Client Initialization

```go
client := joinedgroup.NewJoinedGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedGroupClient.CreateJoinedGroupEvaluateDynamicMembership`

```go
ctx := context.TODO()

payload := joinedgroup.CreateJoinedGroupEvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.CreateJoinedGroupEvaluateDynamicMembership(ctx, payload)
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


// alternatively `client.ListJoinedGroups(ctx)` can be used to do batched pagination
items, err := client.ListJoinedGroupsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
