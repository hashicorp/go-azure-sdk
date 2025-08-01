
## `github.com/hashicorp/go-azure-sdk/resource-manager/aad/2025-06-01/domainservices` Documentation

The `domainservices` SDK allows for interaction with Azure Resource Manager `aad` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/aad/2025-06-01/domainservices"
```


### Client Initialization

```go
client := domainservices.NewDomainServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DomainServicesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := domainservices.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceName")

payload := domainservices.DomainService{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DomainServicesClient.Delete`

```go
ctx := context.TODO()
id := domainservices.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DomainServicesClient.Get`

```go
ctx := context.TODO()
id := domainservices.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainServicesClient.List`

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


### Example Usage: `DomainServicesClient.ListByResourceGroup`

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


### Example Usage: `DomainServicesClient.Unsuspend`

```go
ctx := context.TODO()
id := domainservices.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceName")

read, err := client.Unsuspend(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainServicesClient.Update`

```go
ctx := context.TODO()
id := domainservices.NewDomainServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainServiceName")

payload := domainservices.DomainService{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
