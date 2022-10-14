
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/placementpolicies` Documentation

The `placementpolicies` SDK allows for interaction with the Azure Resource Manager Service `vmware` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/placementpolicies"
```


### Client Initialization

```go
client := placementpolicies.NewPlacementPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlacementPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := placementpolicies.NewPlacementPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue", "placementPolicyValue")

payload := placementpolicies.PlacementPolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PlacementPoliciesClient.Delete`

```go
ctx := context.TODO()
id := placementpolicies.NewPlacementPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue", "placementPolicyValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PlacementPoliciesClient.Get`

```go
ctx := context.TODO()
id := placementpolicies.NewPlacementPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue", "placementPolicyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PlacementPoliciesClient.List`

```go
ctx := context.TODO()
id := placementpolicies.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PlacementPoliciesClient.Update`

```go
ctx := context.TODO()
id := placementpolicies.NewPlacementPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue", "placementPolicyValue")

payload := placementpolicies.PlacementPolicyUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PlacementPoliciesClient.VirtualMachinesRestrictMovement`

```go
ctx := context.TODO()
id := placementpolicies.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "clusterValue", "virtualMachineIdValue")

payload := placementpolicies.VirtualMachineRestrictMovement{
	// ...
}


if err := client.VirtualMachinesRestrictMovementThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
