
## `github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `dnsresolver` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.DnsForwardingRulesetsListByVirtualNetwork`

```go
ctx := context.TODO()
id := commonids.NewVirtualNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkName")

// alternatively `client.DnsForwardingRulesetsListByVirtualNetwork(ctx, id, openapis.DefaultDnsForwardingRulesetsListByVirtualNetworkOperationOptions())` can be used to do batched pagination
items, err := client.DnsForwardingRulesetsListByVirtualNetworkComplete(ctx, id, openapis.DefaultDnsForwardingRulesetsListByVirtualNetworkOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.DnsResolverPoliciesListByVirtualNetwork`

```go
ctx := context.TODO()
id := commonids.NewVirtualNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkName")

// alternatively `client.DnsResolverPoliciesListByVirtualNetwork(ctx, id)` can be used to do batched pagination
items, err := client.DnsResolverPoliciesListByVirtualNetworkComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.DnsResolversListByVirtualNetwork`

```go
ctx := context.TODO()
id := commonids.NewVirtualNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkName")

// alternatively `client.DnsResolversListByVirtualNetwork(ctx, id, openapis.DefaultDnsResolversListByVirtualNetworkOperationOptions())` can be used to do batched pagination
items, err := client.DnsResolversListByVirtualNetworkComplete(ctx, id, openapis.DefaultDnsResolversListByVirtualNetworkOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
