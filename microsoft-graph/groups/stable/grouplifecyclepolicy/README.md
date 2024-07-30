
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/grouplifecyclepolicy` Documentation

The `grouplifecyclepolicy` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/grouplifecyclepolicy"
```


### Client Initialization

```go
client := grouplifecyclepolicy.NewGroupLifecyclePolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupLifecyclePolicyClient.AddGroupGroupLifecyclePolicyGroup`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := grouplifecyclepolicy.AddGroupGroupLifecyclePolicyGroupRequest{
	// ...
}


read, err := client.AddGroupGroupLifecyclePolicyGroup(ctx, id, payload)
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
id := grouplifecyclepolicy.NewGroupID("groupIdValue")

payload := grouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.CreateGroupLifecyclePolicy(ctx, id, payload)
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
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

read, err := client.DeleteGroupLifecyclePolicy(ctx, id)
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
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

read, err := client.GetGroupLifecyclePolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupLifecyclePolicyClient.GetGroupLifecyclePolicyCount`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupID("groupIdValue")

read, err := client.GetGroupLifecyclePolicyCount(ctx, id)
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
id := grouplifecyclepolicy.NewGroupID("groupIdValue")

// alternatively `client.ListGroupLifecyclePolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupLifecyclePoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupLifecyclePolicyClient.RemoveGroupGroupLifecyclePolicyGroup`

```go
ctx := context.TODO()
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := grouplifecyclepolicy.RemoveGroupGroupLifecyclePolicyGroupRequest{
	// ...
}


read, err := client.RemoveGroupGroupLifecyclePolicyGroup(ctx, id, payload)
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
id := grouplifecyclepolicy.NewGroupIdGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := grouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.UpdateGroupLifecyclePolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
