
## `github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-04-01/denyassignments` Documentation

The `denyassignments` SDK allows for interaction with the Azure Resource Manager Service `authorization` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-04-01/denyassignments"
```


### Client Initialization

```go
client := denyassignments.NewDenyAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DenyAssignmentsClient.Get`

```go
ctx := context.TODO()
id := denyassignments.NewScopedDenyAssignmentIdID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DenyAssignmentsClient.GetById`

```go
ctx := context.TODO()
id := denyassignments.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.GetById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DenyAssignmentsClient.List`

```go
ctx := context.TODO()
id := denyassignments.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, denyassignments.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, denyassignments.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DenyAssignmentsClient.ListForResource`

```go
ctx := context.TODO()
id := denyassignments.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "parentResourcePathValue", "resourceTypeValue", "resourceValue")

// alternatively `client.ListForResource(ctx, id, denyassignments.DefaultListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceComplete(ctx, id, denyassignments.DefaultListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DenyAssignmentsClient.ListForResourceGroup`

```go
ctx := context.TODO()
id := denyassignments.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListForResourceGroup(ctx, id, denyassignments.DefaultListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceGroupComplete(ctx, id, denyassignments.DefaultListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DenyAssignmentsClient.ListForScope`

```go
ctx := context.TODO()
id := denyassignments.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListForScope(ctx, id, denyassignments.DefaultListForScopeOperationOptions())` can be used to do batched pagination
items, err := client.ListForScopeComplete(ctx, id, denyassignments.DefaultListForScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
