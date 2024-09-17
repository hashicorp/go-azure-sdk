
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/jitnetworkaccesspolicies` Documentation

The `jitnetworkaccesspolicies` SDK allows for interaction with Azure Resource Manager `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/jitnetworkaccesspolicies"
```


### Client Initialization

```go
client := jitnetworkaccesspolicies.NewJitNetworkAccessPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JitNetworkAccessPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewJitNetworkAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "jitNetworkAccessPolicyValue")

payload := jitnetworkaccesspolicies.JitNetworkAccessPolicy{
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


### Example Usage: `JitNetworkAccessPoliciesClient.Delete`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewJitNetworkAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "jitNetworkAccessPolicyValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.Get`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewJitNetworkAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "jitNetworkAccessPolicyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.Initiate`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewJitNetworkAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "jitNetworkAccessPolicyValue")

payload := jitnetworkaccesspolicies.JitNetworkAccessPolicyInitiateRequest{
	// ...
}


read, err := client.Initiate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.ListByRegion`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListByRegion(ctx, id)` can be used to do batched pagination
items, err := client.ListByRegionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JitNetworkAccessPoliciesClient.ListByResourceGroupAndRegion`

```go
ctx := context.TODO()
id := jitnetworkaccesspolicies.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

// alternatively `client.ListByResourceGroupAndRegion(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupAndRegionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
