
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-09-01/webapplicationfirewallmanagedrulesets` Documentation

The `webapplicationfirewallmanagedrulesets` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2024-09-01/webapplicationfirewallmanagedrulesets"
```


### Client Initialization

```go
client := webapplicationfirewallmanagedrulesets.NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WebApplicationFirewallManagedRuleSetsClient.ManagedRuleSetsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ManagedRuleSetsList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedRuleSetsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
