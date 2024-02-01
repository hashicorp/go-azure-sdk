
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/pricings` Documentation

The `pricings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/pricings"
```


### Client Initialization

```go
client := pricings.NewPricingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PricingsClient.CreateOrUpdateResourceGroupPricing`

```go
ctx := context.TODO()
id := pricings.NewProviderPricingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "pricingValue")

payload := pricings.Pricing{
	// ...
}


read, err := client.CreateOrUpdateResourceGroupPricing(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PricingsClient.GetResourceGroupPricing`

```go
ctx := context.TODO()
id := pricings.NewProviderPricingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "pricingValue")

read, err := client.GetResourceGroupPricing(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PricingsClient.GetSubscriptionPricing`

```go
ctx := context.TODO()
id := pricings.NewPricingID("12345678-1234-9876-4563-123456789012", "pricingValue")

read, err := client.GetSubscriptionPricing(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PricingsClient.List`

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


### Example Usage: `PricingsClient.ListByResourceGroup`

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


### Example Usage: `PricingsClient.UpdateSubscriptionPricing`

```go
ctx := context.TODO()
id := pricings.NewPricingID("12345678-1234-9876-4563-123456789012", "pricingValue")

payload := pricings.Pricing{
	// ...
}


read, err := client.UpdateSubscriptionPricing(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
