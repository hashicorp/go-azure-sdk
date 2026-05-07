
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.AttestationsCreateOrUpdateAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationName")

payload := openapis.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtResourceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.AttestationsCreateOrUpdateAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationName")

payload := openapis.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtResourceGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.AttestationsCreateOrUpdateAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationName")

payload := openapis.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.AttestationsDeleteAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationName")

read, err := client.AttestationsDeleteAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsDeleteAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationName")

read, err := client.AttestationsDeleteAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsDeleteAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationName")

read, err := client.AttestationsDeleteAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsGetAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationName")

read, err := client.AttestationsGetAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsGetAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationName")

read, err := client.AttestationsGetAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsGetAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationName")

read, err := client.AttestationsGetAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.AttestationsListForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.AttestationsListForResource(ctx, id, openapis.DefaultAttestationsListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForResourceComplete(ctx, id, openapis.DefaultAttestationsListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.AttestationsListForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.AttestationsListForResourceGroup(ctx, id, openapis.DefaultAttestationsListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForResourceGroupComplete(ctx, id, openapis.DefaultAttestationsListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.AttestationsListForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.AttestationsListForSubscription(ctx, id, openapis.DefaultAttestationsListForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForSubscriptionComplete(ctx, id, openapis.DefaultAttestationsListForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

read, err := client.ComponentPolicyStatesListQueryResultsForPolicyDefinition(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.ComponentPolicyStatesListQueryResultsForResource(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ComponentPolicyStatesListQueryResultsForResourceGroup(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

read, err := client.ComponentPolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ComponentPolicyStatesListQueryResultsForSubscription(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ComponentPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

read, err := client.ComponentPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, openapis.DefaultComponentPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewManagementGroupID("managementGroupName")

// alternatively `client.PolicyEventsListQueryResultsForManagementGroup(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForManagementGroupComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

// alternatively `client.PolicyEventsListQueryResultsForPolicyDefinition(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForPolicyDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForPolicyDefinitionComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForPolicySetDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName")

// alternatively `client.PolicyEventsListQueryResultsForPolicySetDefinition(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForPolicySetDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForPolicySetDefinitionComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.PolicyEventsListQueryResultsForResource(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForResourceComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.PolicyEventsListQueryResultsForResourceGroup(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForResourceGroupComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

// alternatively `client.PolicyEventsListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.PolicyEventsListQueryResultsForSubscription(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForSubscriptionComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyEventsListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

// alternatively `client.PolicyEventsListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.PolicyEventsListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx, id, openapis.DefaultPolicyEventsListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyMetadataGetResource`

```go
ctx := context.TODO()
id := openapis.NewPolicyMetadataID("policyMetadataName")

read, err := client.PolicyMetadataGetResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyMetadataList`

```go
ctx := context.TODO()


// alternatively `client.PolicyMetadataList(ctx, openapis.DefaultPolicyMetadataListOperationOptions())` can be used to do batched pagination
items, err := client.PolicyMetadataListComplete(ctx, openapis.DefaultPolicyMetadataListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyRestrictionsCheckAtManagementGroupScope`

```go
ctx := context.TODO()
id := openapis.NewManagementGroupID("managementGroupName")

payload := openapis.CheckManagementGroupRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtManagementGroupScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyRestrictionsCheckAtResourceGroupScope`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.CheckRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtResourceGroupScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyRestrictionsCheckAtSubscriptionScope`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.CheckRestrictionsRequest{
	// ...
}


read, err := client.PolicyRestrictionsCheckAtSubscriptionScope(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2PolicyStatePolicyStatesResourceID("managementGroupName", "default")

// alternatively `client.PolicyStatesListQueryResultsForManagementGroup(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForManagementGroupComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForPolicyDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicyDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyDefinitionName", "default")

// alternatively `client.PolicyStatesListQueryResultsForPolicyDefinition(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForPolicyDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForPolicyDefinitionComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForPolicySetDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName", "default")

// alternatively `client.PolicyStatesListQueryResultsForPolicySetDefinition(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForPolicySetDefinitionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForPolicySetDefinitionComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForResource`

```go
ctx := context.TODO()
id := openapis.NewScopedPolicyStatesResourceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "default")

// alternatively `client.PolicyStatesListQueryResultsForResource(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForResourceComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewPolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "default")

// alternatively `client.PolicyStatesListQueryResultsForResourceGroup(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForResourceGroupComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName", "default")

// alternatively `client.PolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignmentComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForSubscription`

```go
ctx := context.TODO()
id := openapis.NewPolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "default")

// alternatively `client.PolicyStatesListQueryResultsForSubscription(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForSubscriptionComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID("12345678-1234-9876-4563-123456789012", "policyAssignmentName", "default")

// alternatively `client.PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())` can be used to do batched pagination
items, err := client.PolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentComplete(ctx, id, openapis.DefaultPolicyStatesListQueryResultsForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewManagementGroupID("managementGroupName")

read, err := client.PolicyStatesSummarizeForManagementGroup(ctx, id, openapis.DefaultPolicyStatesSummarizeForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForPolicyDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicyDefinitionID("12345678-1234-9876-4563-123456789012", "policyDefinitionName")

read, err := client.PolicyStatesSummarizeForPolicyDefinition(ctx, id, openapis.DefaultPolicyStatesSummarizeForPolicyDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForPolicySetDefinition`

```go
ctx := context.TODO()
id := openapis.NewPolicySetDefinitionID("12345678-1234-9876-4563-123456789012", "policySetDefinitionName")

read, err := client.PolicyStatesSummarizeForPolicySetDefinition(ctx, id, openapis.DefaultPolicyStatesSummarizeForPolicySetDefinitionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.PolicyStatesSummarizeForResource(ctx, id, openapis.DefaultPolicyStatesSummarizeForResourceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.PolicyStatesSummarizeForResourceGroup(ctx, id, openapis.DefaultPolicyStatesSummarizeForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForResourceGroupLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewAuthorizationNamespacePolicyAssignmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "policyAssignmentName")

read, err := client.PolicyStatesSummarizeForResourceGroupLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyStatesSummarizeForResourceGroupLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.PolicyStatesSummarizeForSubscription(ctx, id, openapis.DefaultPolicyStatesSummarizeForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesSummarizeForSubscriptionLevelPolicyAssignment`

```go
ctx := context.TODO()
id := openapis.NewPolicyAssignmentID("12345678-1234-9876-4563-123456789012", "policyAssignmentName")

read, err := client.PolicyStatesSummarizeForSubscriptionLevelPolicyAssignment(ctx, id, openapis.DefaultPolicyStatesSummarizeForSubscriptionLevelPolicyAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.PolicyStatesTriggerResourceGroupEvaluation`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

if err := client.PolicyStatesTriggerResourceGroupEvaluationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.PolicyStatesTriggerSubscriptionEvaluation`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

if err := client.PolicyStatesTriggerSubscriptionEvaluationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.RemediationsCancelAtManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2RemediationID("managementGroupId", "remediationName")

read, err := client.RemediationsCancelAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCancelAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedRemediationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "remediationName")

read, err := client.RemediationsCancelAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCancelAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderRemediationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remediationName")

read, err := client.RemediationsCancelAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCancelAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewRemediationID("12345678-1234-9876-4563-123456789012", "remediationName")

read, err := client.RemediationsCancelAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCreateOrUpdateAtManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2RemediationID("managementGroupId", "remediationName")

payload := openapis.Remediation{
	// ...
}


read, err := client.RemediationsCreateOrUpdateAtManagementGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCreateOrUpdateAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedRemediationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "remediationName")

payload := openapis.Remediation{
	// ...
}


read, err := client.RemediationsCreateOrUpdateAtResource(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCreateOrUpdateAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderRemediationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remediationName")

payload := openapis.Remediation{
	// ...
}


read, err := client.RemediationsCreateOrUpdateAtResourceGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsCreateOrUpdateAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewRemediationID("12345678-1234-9876-4563-123456789012", "remediationName")

payload := openapis.Remediation{
	// ...
}


read, err := client.RemediationsCreateOrUpdateAtSubscription(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsDeleteAtManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2RemediationID("managementGroupId", "remediationName")

read, err := client.RemediationsDeleteAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsDeleteAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedRemediationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "remediationName")

read, err := client.RemediationsDeleteAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsDeleteAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderRemediationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remediationName")

read, err := client.RemediationsDeleteAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsDeleteAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewRemediationID("12345678-1234-9876-4563-123456789012", "remediationName")

read, err := client.RemediationsDeleteAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsGetAtManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2RemediationID("managementGroupId", "remediationName")

read, err := client.RemediationsGetAtManagementGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsGetAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedRemediationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "remediationName")

read, err := client.RemediationsGetAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsGetAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderRemediationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remediationName")

read, err := client.RemediationsGetAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsGetAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewRemediationID("12345678-1234-9876-4563-123456789012", "remediationName")

read, err := client.RemediationsGetAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RemediationsListDeploymentsAtManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewProviders2RemediationID("managementGroupId", "remediationName")

// alternatively `client.RemediationsListDeploymentsAtManagementGroup(ctx, id, openapis.DefaultRemediationsListDeploymentsAtManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListDeploymentsAtManagementGroupComplete(ctx, id, openapis.DefaultRemediationsListDeploymentsAtManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListDeploymentsAtResource`

```go
ctx := context.TODO()
id := openapis.NewScopedRemediationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "remediationName")

// alternatively `client.RemediationsListDeploymentsAtResource(ctx, id, openapis.DefaultRemediationsListDeploymentsAtResourceOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListDeploymentsAtResourceComplete(ctx, id, openapis.DefaultRemediationsListDeploymentsAtResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListDeploymentsAtResourceGroup`

```go
ctx := context.TODO()
id := openapis.NewProviderRemediationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remediationName")

// alternatively `client.RemediationsListDeploymentsAtResourceGroup(ctx, id, openapis.DefaultRemediationsListDeploymentsAtResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListDeploymentsAtResourceGroupComplete(ctx, id, openapis.DefaultRemediationsListDeploymentsAtResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListDeploymentsAtSubscription`

```go
ctx := context.TODO()
id := openapis.NewRemediationID("12345678-1234-9876-4563-123456789012", "remediationName")

// alternatively `client.RemediationsListDeploymentsAtSubscription(ctx, id, openapis.DefaultRemediationsListDeploymentsAtSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListDeploymentsAtSubscriptionComplete(ctx, id, openapis.DefaultRemediationsListDeploymentsAtSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListForManagementGroup`

```go
ctx := context.TODO()
id := openapis.NewManagementGroupID("managementGroupName")

// alternatively `client.RemediationsListForManagementGroup(ctx, id, openapis.DefaultRemediationsListForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListForManagementGroupComplete(ctx, id, openapis.DefaultRemediationsListForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.RemediationsListForResource(ctx, id, openapis.DefaultRemediationsListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListForResourceComplete(ctx, id, openapis.DefaultRemediationsListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.RemediationsListForResourceGroup(ctx, id, openapis.DefaultRemediationsListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListForResourceGroupComplete(ctx, id, openapis.DefaultRemediationsListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.RemediationsListForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.RemediationsListForSubscription(ctx, id, openapis.DefaultRemediationsListForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.RemediationsListForSubscriptionComplete(ctx, id, openapis.DefaultRemediationsListForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
