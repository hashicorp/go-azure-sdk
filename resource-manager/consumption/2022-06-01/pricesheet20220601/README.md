
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-06-01/pricesheet20220601` Documentation

The `pricesheet20220601` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2022-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-06-01/pricesheet20220601"
```


### Client Initialization

```go
client := pricesheet20220601.NewPriceSheet20220601ClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheet20220601Client.PriceSheetsGet`

```go
ctx := context.TODO()
id := pricesheet20220601.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.PriceSheetsGet(ctx, id, pricesheet20220601.DefaultPriceSheetsGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PriceSheet20220601Client.PriceSheetsGetByBillingPeriod`

```go
ctx := context.TODO()
id := pricesheet20220601.NewBillingPeriodID("12345678-1234-9876-4563-123456789012", "billingPeriodValue")

read, err := client.PriceSheetsGetByBillingPeriod(ctx, id, pricesheet20220601.DefaultPriceSheetsGetByBillingPeriodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
