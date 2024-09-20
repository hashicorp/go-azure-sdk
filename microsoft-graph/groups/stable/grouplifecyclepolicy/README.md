
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/grouplifecyclepolicy` Documentation

The `grouplifecyclepolicy` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/grouplifecyclepolicy"
```


### Client Initialization

```go
client := grouplifecyclepolicy.NewGroupLifecyclePolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupLifecyclePolicyClient.AddGroupLifecyclePolicyGroup`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupId", "groupLifecyclePolicyId")

payload := grouplifecyclepolicy.AddGroupLifecyclePolicyGroupRequest{
	// ...
}


read, err := client.AddGroupLifecyclePolicyGroup(ctx, id, payload, grouplifecyclepolicy.DefaultAddGroupLifecyclePolicyGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.CreateGroupLifecyclePolicy`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupID("groupId")

payload := grouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.CreateGroupLifecyclePolicy(ctx, id, payload, grouplifecyclepolicy.DefaultCreateGroupLifecyclePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.DeleteGroupLifecyclePolicy`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupId", "groupLifecyclePolicyId")

read, err := client.DeleteGroupLifecyclePolicy(ctx, id, grouplifecyclepolicy.DefaultDeleteGroupLifecyclePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.GetGroupLifecyclePoliciesCount`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupID("groupId")

read, err := client.GetGroupLifecyclePoliciesCount(ctx, id, grouplifecyclepolicy.DefaultGetGroupLifecyclePoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.GetGroupLifecyclePolicy`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupId", "groupLifecyclePolicyId")

read, err := client.GetGroupLifecyclePolicy(ctx, id, grouplifecyclepolicy.DefaultGetGroupLifecyclePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.ListGroupLifecyclePolicies`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupID("groupId")

// alternatively `client.ListGroupLifecyclePolicies(ctx, id, grouplifecyclepolicy.DefaultListGroupLifecyclePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupLifecyclePoliciesComplete(ctx, id, grouplifecyclepolicy.DefaultListGroupLifecyclePoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupLifecyclePolicyClient.RemoveGroupLifecyclePolicyGroup`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupId", "groupLifecyclePolicyId")

payload := grouplifecyclepolicy.RemoveGroupLifecyclePolicyGroupRequest{
	// ...
}


read, err := client.RemoveGroupLifecyclePolicyGroup(ctx, id, payload, grouplifecyclepolicy.DefaultRemoveGroupLifecyclePolicyGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.UpdateGroupLifecyclePolicy`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupId", "groupLifecyclePolicyId")

payload := grouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.UpdateGroupLifecyclePolicy(ctx, id, payload, grouplifecyclepolicy.DefaultUpdateGroupLifecyclePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
