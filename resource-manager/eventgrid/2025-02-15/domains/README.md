
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/domains` Documentation

The `domains` SDK allows for interaction with Azure Resource Manager `eventgrid` (API Version `2025-02-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2025-02-15/domains"
```


### Client Initialization

```go
client := domains.NewDomainsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DomainsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

payload := domains.Domain{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DomainsClient.Delete`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DomainsClient.Get`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, domains.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, domains.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, domains.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, domains.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainsClient.ListSharedAccessKeys`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

read, err := client.ListSharedAccessKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.RegenerateKey`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

payload := domains.DomainRegenerateKeyRequest{
	// ...
}


read, err := client.RegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.Update`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainName")

payload := domains.DomainUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
