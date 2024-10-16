
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policystates` Documentation

The `policystates` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policystates"
```


### Client Initialization

```go
client := policystates.NewPolicyStatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForManagementGroup`

```go
ctx := context.TODO()
id := policystates.NewProviders2PolicyStatePolicyStatesResourceID("managementGroupName", "default")

// alternatively `client.ListQueryResultsForManagementGroup(ctx, id, policystates.DefaultListQueryResultsForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForManagementGroupComplete(ctx, id, policystates.DefaultListQueryResultsForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := policystates.NewPolicyDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyDefinitionName", "default")

// alternatively `client.ListQueryResultsForPolicyDefinition(ctx, id, policystates.DefaultListQueryResultsForPolicyDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForPolicyDefinitionComplete(ctx, id, policystates.DefaultListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForPolicySetDefinition`

```go
ctx := context.TODO()
id := policystates.NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "default")

// alternatively `client.ListQueryResultsForPolicySetDefinition(ctx, id, policystates.DefaultListQueryResultsForPolicySetDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForPolicySetDefinitionComplete(ctx, id, policystates.DefaultListQueryResultsForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForResource`

```go
ctx := context.TODO()
id := policystates.NewScopedPolicyStatesResourceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "default")

// alternatively `client.ListQueryResultsForResource(ctx, id, policystates.DefaultListQueryResultsForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceComplete(ctx, id, policystates.DefaultListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := policystates.NewPolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "default")

// alternatively `client.ListQueryResultsForResourceGroup(ctx, id, policystates.DefaultListQueryResultsForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceGroupComplete(ctx, id, policystates.DefaultListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policystates.NewAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName", "default")

// alternatively `client.ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, policystates.DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx, id, policystates.DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := policystates.NewPolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "default")

// alternatively `client.ListQueryResultsForSubscription(ctx, id, policystates.DefaultListQueryResultsForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForSubscriptionComplete(ctx, id, policystates.DefaultListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policystates.NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyAssignmentName", "default")

// alternatively `client.ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, policystates.DefaultListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.ListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx, id, policystates.DefaultListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyStatesClient.SummarizeForManagementGroup`

```go
ctx := context.TODO()
id := policystates.NewManagementGroupID("managementGroupId")

read, err := client.SummarizeForManagementGroup(ctx, id, policystates.DefaultSummarizeForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForPolicyDefinition`

```go
ctx := context.TODO()
id := policystates.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

read, err := client.SummarizeForPolicyDefinition(ctx, id, policystates.DefaultSummarizeForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForPolicySetDefinition`

```go
ctx := context.TODO()
id := policystates.NewPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName")

read, err := client.SummarizeForPolicySetDefinition(ctx, id, policystates.DefaultSummarizeForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.SummarizeForResource(ctx, id, policystates.DefaultSummarizeForResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.SummarizeForResourceGroup(ctx, id, policystates.DefaultSummarizeForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policystates.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

read, err := client.SummarizeForResourceGroupLevelPolicyAssignment(ctx, id, policystates.DefaultSummarizeForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.SummarizeForSubscription(ctx, id, policystates.DefaultSummarizeForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.SummarizeForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := policystates.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

read, err := client.SummarizeForSubscriptionLevelPolicyAssignment(ctx, id, policystates.DefaultSummarizeForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyStatesClient.TriggerResourceGroupEvaluation`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

if err := client.TriggerResourceGroupEvaluationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PolicyStatesClient.TriggerSubscriptionEvaluation`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

if err := client.TriggerSubscriptionEvaluationThenPoll(ctx, id); err != nil {
	// handle the error
}
```
