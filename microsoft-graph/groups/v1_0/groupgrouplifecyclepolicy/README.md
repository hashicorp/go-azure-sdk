
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupgrouplifecyclepolicy` Documentation

The `groupgrouplifecyclepolicy` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupgrouplifecyclepolicy"
```


### Client Initialization

```go
client := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupGroupLifecyclePolicyClient.AddGroupByIdGroupLifecyclePolicyByIdGroup`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := groupgrouplifecyclepolicy.AddGroupByIdGroupLifecyclePolicyByIdGroupRequest{
	// ...
}


read, err := client.AddGroupByIdGroupLifecyclePolicyByIdGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.CreateGroupByIdGroupLifecyclePolicy`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupID("groupIdValue")

payload := groupgrouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.CreateGroupByIdGroupLifecyclePolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.DeleteGroupByIdGroupLifecyclePolicyById`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

read, err := client.DeleteGroupByIdGroupLifecyclePolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.GetGroupByIdGroupLifecyclePolicyById`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

read, err := client.GetGroupByIdGroupLifecyclePolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.GetGroupByIdGroupLifecyclePolicyCount`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdGroupLifecyclePolicyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.ListGroupByIdGroupLifecyclePolicies`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdGroupLifecyclePolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdGroupLifecyclePoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.RemoveGroupByIdGroupLifecyclePolicyByIdGroup`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := groupgrouplifecyclepolicy.RemoveGroupByIdGroupLifecyclePolicyByIdGroupRequest{
	// ...
}


read, err := client.RemoveGroupByIdGroupLifecyclePolicyByIdGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupGroupLifecyclePolicyClient.UpdateGroupByIdGroupLifecyclePolicyById`

```go
ctx := context.TODO()
id := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyID("groupIdValue", "groupLifecyclePolicyIdValue")

payload := groupgrouplifecyclepolicy.GroupLifecyclePolicy{
	// ...
}


read, err := client.UpdateGroupByIdGroupLifecyclePolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
