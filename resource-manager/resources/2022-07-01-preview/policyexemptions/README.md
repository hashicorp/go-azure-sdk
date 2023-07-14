
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-07-01-preview/policyexemptions` Documentation

The `policyexemptions` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-07-01-preview/policyexemptions"
```


### Client Initialization

```go
client := policyexemptions.NewPolicyExemptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyExemptionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := policyexemptions.NewScopedPolicyExemptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "policyExemptionValue")

payload := policyexemptions.PolicyExemption{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyExemptionsClient.Delete`

```go
ctx := context.TODO()
id := policyexemptions.NewScopedPolicyExemptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "policyExemptionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyExemptionsClient.Get`

```go
ctx := context.TODO()
id := policyexemptions.NewScopedPolicyExemptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "policyExemptionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyExemptionsClient.List`

```go
ctx := context.TODO()
id := policyexemptions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, policyexemptions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, policyexemptions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyExemptionsClient.ListForManagementGroup`

```go
ctx := context.TODO()
id := policyexemptions.NewManagementGroupID("groupIdValue")

// alternatively `client.ListForManagementGroup(ctx, id, policyexemptions.DefaultListForManagementGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListForManagementGroupComplete(ctx, id, policyexemptions.DefaultListForManagementGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyExemptionsClient.ListForResource`

```go
ctx := context.TODO()
id := policyexemptions.NewScopedResourceID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "resourceValue")

// alternatively `client.ListForResource(ctx, id, policyexemptions.DefaultListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceComplete(ctx, id, policyexemptions.DefaultListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyExemptionsClient.ListForResourceGroup`

```go
ctx := context.TODO()
id := policyexemptions.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListForResourceGroup(ctx, id, policyexemptions.DefaultListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceGroupComplete(ctx, id, policyexemptions.DefaultListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyExemptionsClient.Update`

```go
ctx := context.TODO()
id := policyexemptions.NewScopedPolicyExemptionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "policyExemptionValue")

payload := policyexemptions.PolicyExemptionUpdate{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
