
## `github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverpolicies` Documentation

The `dnsresolverpolicies` SDK allows for interaction with Azure Resource Manager `dnsresolver` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverpolicies"
```


### Client Initialization

```go
client := dnsresolverpolicies.NewDnsResolverPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DnsResolverPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dnsresolverpolicies.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

payload := dnsresolverpolicies.DnsResolverPolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, dnsresolverpolicies.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverPoliciesClient.Delete`

```go
ctx := context.TODO()
id := dnsresolverpolicies.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

if err := client.DeleteThenPoll(ctx, id, dnsresolverpolicies.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverPoliciesClient.Get`

```go
ctx := context.TODO()
id := dnsresolverpolicies.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DnsResolverPoliciesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, dnsresolverpolicies.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dnsresolverpolicies.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsResolverPoliciesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, dnsresolverpolicies.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, dnsresolverpolicies.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsResolverPoliciesClient.Update`

```go
ctx := context.TODO()
id := dnsresolverpolicies.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

payload := dnsresolverpolicies.DnsResolverPolicyPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, dnsresolverpolicies.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
