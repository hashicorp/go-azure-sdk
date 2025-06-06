
## `github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverpolicyvirtualnetworklinks` Documentation

The `dnsresolverpolicyvirtualnetworklinks` SDK allows for interaction with Azure Resource Manager `dnsresolver` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverpolicyvirtualnetworklinks"
```


### Client Initialization

```go
client := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyVirtualNetworkLinksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DnsResolverPolicyVirtualNetworkLinksClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "virtualNetworkLinkName")

payload := dnsresolverpolicyvirtualnetworklinks.DnsResolverPolicyVirtualNetworkLink{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, dnsresolverpolicyvirtualnetworklinks.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverPolicyVirtualNetworkLinksClient.Delete`

```go
ctx := context.TODO()
id := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "virtualNetworkLinkName")

if err := client.DeleteThenPoll(ctx, id, dnsresolverpolicyvirtualnetworklinks.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverPolicyVirtualNetworkLinksClient.Get`

```go
ctx := context.TODO()
id := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "virtualNetworkLinkName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DnsResolverPolicyVirtualNetworkLinksClient.List`

```go
ctx := context.TODO()
id := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName")

// alternatively `client.List(ctx, id, dnsresolverpolicyvirtualnetworklinks.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dnsresolverpolicyvirtualnetworklinks.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsResolverPolicyVirtualNetworkLinksClient.Update`

```go
ctx := context.TODO()
id := dnsresolverpolicyvirtualnetworklinks.NewDnsResolverPolicyVirtualNetworkLinkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverPolicyName", "virtualNetworkLinkName")

payload := dnsresolverpolicyvirtualnetworklinks.DnsResolverPolicyVirtualNetworkLinkPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, dnsresolverpolicyvirtualnetworklinks.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
