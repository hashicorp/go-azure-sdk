
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/componentpolicystates` Documentation

The `componentpolicystates` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/componentpolicystates"
```


### Client Initialization

```go
client := componentpolicystates.NewComponentPolicyStatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := componentpolicystates.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

read, err := client.ListQueryResultsForPolicyDefinition(ctx, id, componentpolicystates.DefaultListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.ListQueryResultsForResource(ctx, id, componentpolicystates.DefaultListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ListQueryResultsForResourceGroup(ctx, id, componentpolicystates.DefaultListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := componentpolicystates.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

read, err := client.ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, componentpolicystates.DefaultListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListQueryResultsForSubscription(ctx, id, componentpolicystates.DefaultListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := componentpolicystates.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

read, err := client.ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, componentpolicystates.DefaultListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
