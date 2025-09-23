
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/cdnwebapplicationfirewallpolicies` Documentation

The `cdnwebapplicationfirewallpolicies` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/cdnwebapplicationfirewallpolicies"
```


### Client Initialization

```go
client := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CdnWebApplicationFirewallPoliciesClient.PoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cdnWebApplicationFirewallPolicyName")

payload := cdnwebapplicationfirewallpolicies.CdnWebApplicationFirewallPolicy{
	// ...
}


if err := client.PoliciesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CdnWebApplicationFirewallPoliciesClient.PoliciesDelete`

```go
ctx := context.TODO()
id := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cdnWebApplicationFirewallPolicyName")

read, err := client.PoliciesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CdnWebApplicationFirewallPoliciesClient.PoliciesGet`

```go
ctx := context.TODO()
id := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cdnWebApplicationFirewallPolicyName")

read, err := client.PoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CdnWebApplicationFirewallPoliciesClient.PoliciesList`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.PoliciesList(ctx, id)` can be used to do batched pagination
items, err := client.PoliciesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CdnWebApplicationFirewallPoliciesClient.PoliciesUpdate`

```go
ctx := context.TODO()
id := cdnwebapplicationfirewallpolicies.NewCdnWebApplicationFirewallPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cdnWebApplicationFirewallPolicyName")

payload := cdnwebapplicationfirewallpolicies.CdnWebApplicationFirewallPolicyPatchParameters{
	// ...
}


if err := client.PoliciesUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
