
## `github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverdomainlists` Documentation

The `dnsresolverdomainlists` SDK allows for interaction with Azure Resource Manager `dnsresolver` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dnsresolver/2025-05-01/dnsresolverdomainlists"
```


### Client Initialization

```go
client := dnsresolverdomainlists.NewDnsResolverDomainListsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DnsResolverDomainListsClient.Bulk`

```go
ctx := context.TODO()
id := dnsresolverdomainlists.NewDnsResolverDomainListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverDomainListName")

payload := dnsresolverdomainlists.DnsResolverDomainListBulk{
	// ...
}


if err := client.BulkThenPoll(ctx, id, payload, dnsresolverdomainlists.DefaultBulkOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverDomainListsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dnsresolverdomainlists.NewDnsResolverDomainListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverDomainListName")

payload := dnsresolverdomainlists.DnsResolverDomainList{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, dnsresolverdomainlists.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverDomainListsClient.Delete`

```go
ctx := context.TODO()
id := dnsresolverdomainlists.NewDnsResolverDomainListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverDomainListName")

if err := client.DeleteThenPoll(ctx, id, dnsresolverdomainlists.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DnsResolverDomainListsClient.Get`

```go
ctx := context.TODO()
id := dnsresolverdomainlists.NewDnsResolverDomainListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverDomainListName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DnsResolverDomainListsClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, dnsresolverdomainlists.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, dnsresolverdomainlists.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsResolverDomainListsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, dnsresolverdomainlists.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, dnsresolverdomainlists.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DnsResolverDomainListsClient.Update`

```go
ctx := context.TODO()
id := dnsresolverdomainlists.NewDnsResolverDomainListID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dnsResolverDomainListName")

payload := dnsresolverdomainlists.DnsResolverDomainListPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, dnsresolverdomainlists.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
