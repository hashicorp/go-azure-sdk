
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policyevents` Documentation

The `policyevents` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policyevents"
```


### Client Initialization

```go
client := policyevents.NewPolicyEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForManagementGroup`

```go
ctx := context.TODO()
id := policyevents.NewManagementGroupID("managementGroupId")

// alternatively `client.ListQueryResultsForManagementGroup(ctx, id, policyevents.DefaultListQueryResultsForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForManagementGroupComplete(ctx, id, policyevents.DefaultListQueryResultsForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := policyevents.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

// alternatively `client.ListQueryResultsForPolicyDefinition(ctx, id, policyevents.DefaultListQueryResultsForPolicyDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForPolicyDefinitionComplete(ctx, id, policyevents.DefaultListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForPolicySetDefinition`

```go
ctx := context.TODO()
id := policyevents.NewPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName")

// alternatively `client.ListQueryResultsForPolicySetDefinition(ctx, id, policyevents.DefaultListQueryResultsForPolicySetDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForPolicySetDefinitionComplete(ctx, id, policyevents.DefaultListQueryResultsForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListQueryResultsForResource(ctx, id, policyevents.DefaultListQueryResultsForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceComplete(ctx, id, policyevents.DefaultListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListQueryResultsForResourceGroup(ctx, id, policyevents.DefaultListQueryResultsForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceGroupComplete(ctx, id, policyevents.DefaultListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policyevents.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

// alternatively `client.ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, policyevents.DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx, id, policyevents.DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListQueryResultsForSubscription(ctx, id, policyevents.DefaultListQueryResultsForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForSubscriptionComplete(ctx, id, policyevents.DefaultListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyEventsClient.ListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policyevents.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

// alternatively `client.ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, policyevents.DefaultListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx, id, policyevents.DefaultListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
