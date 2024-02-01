
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/domains` Documentation

The `domains` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2023-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-01-01/domains"
```


### Client Initialization

```go
client := domains.NewDomainsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DomainsClient.CheckAvailability`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := domains.NameIdentifier{
	// ...
}


read, err := client.CheckAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

payload := domains.Domain{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DomainsClient.CreateOrUpdateOwnershipIdentifier`

```go
ctx := context.TODO()
id := domains.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue", "domainOwnershipIdentifierValue")

payload := domains.DomainOwnershipIdentifier{
	// ...
}


read, err := client.CreateOrUpdateOwnershipIdentifier(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.Delete`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

read, err := client.Delete(ctx, id, domains.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.DeleteOwnershipIdentifier`

```go
ctx := context.TODO()
id := domains.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue", "domainOwnershipIdentifierValue")

read, err := client.DeleteOwnershipIdentifier(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.Get`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.GetControlCenterSsoRequest`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.GetControlCenterSsoRequest(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.GetOwnershipIdentifier`

```go
ctx := context.TODO()
id := domains.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue", "domainOwnershipIdentifierValue")

read, err := client.GetOwnershipIdentifier(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.List`

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


### Example Usage: `DomainsClient.ListByResourceGroup`

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


### Example Usage: `DomainsClient.ListOwnershipIdentifiers`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

// alternatively `client.ListOwnershipIdentifiers(ctx, id)` can be used to do batched pagination
items, err := client.ListOwnershipIdentifiersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainsClient.ListRecommendations`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := domains.DomainRecommendationSearchParameters{
	// ...
}


// alternatively `client.ListRecommendations(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListRecommendationsComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainsClient.Renew`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

read, err := client.Renew(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainsClient.TransferOut`

```go
ctx := context.TODO()
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

read, err := client.TransferOut(ctx, id)
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
id := domains.NewDomainID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue")

payload := domains.DomainPatchResource{
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


### Example Usage: `DomainsClient.UpdateOwnershipIdentifier`

```go
ctx := context.TODO()
id := domains.NewDomainOwnershipIdentifierID("12345678-1234-9876-4563-123456789012", "example-resource-group", "domainValue", "domainOwnershipIdentifierValue")

payload := domains.DomainOwnershipIdentifier{
	// ...
}


read, err := client.UpdateOwnershipIdentifier(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
