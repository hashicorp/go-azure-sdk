
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-01-01/pricings` Documentation

The `pricings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2024-01-01/pricings"
```


### Client Initialization

```go
client := pricings.NewPricingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PricingsClient.Delete`

```go
ctx := context.TODO()
id := pricings.NewPricingID("scopeIdValue", "pricingValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PricingsClient.Get`

```go
ctx := context.TODO()
id := pricings.NewPricingID("scopeIdValue", "pricingValue")

read, err := client.Get(ctx, id)
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
id := pricings.NewScopeIdID("scopeIdValue")

read, err := client.List(ctx, id, pricings.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PricingsClient.Update`

```go
ctx := context.TODO()
id := pricings.NewPricingID("scopeIdValue", "pricingValue")

payload := pricings.Pricing{
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
