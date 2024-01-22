
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billings` Documentation

The `billings` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billings"
```


### Client Initialization

```go
client := billings.NewBillingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingsClient.BillingSubscriptionsMove`

```go
ctx := context.TODO()
id := billings.NewBillingAccountBillingSubscriptionID("billingAccountValue", "subscriptionIdValue")

payload := billings.TransferBillingSubscriptionRequestProperties{
	// ...
}


if err := client.BillingSubscriptionsMoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingsClient.BillingSubscriptionsValidateMove`

```go
ctx := context.TODO()
id := billings.NewBillingAccountBillingSubscriptionID("billingAccountValue", "subscriptionIdValue")

payload := billings.TransferBillingSubscriptionRequestProperties{
	// ...
}


read, err := client.BillingSubscriptionsValidateMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingsClient.ProductsValidateMove`

```go
ctx := context.TODO()
id := billings.NewProductID("billingAccountValue", "productValue")

payload := billings.TransferProductRequestProperties{
	// ...
}


read, err := client.ProductsValidateMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
