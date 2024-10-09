
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/paymentmethods` Documentation

The `paymentmethods` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/paymentmethods"
```


### Client Initialization

```go
client := paymentmethods.NewPaymentMethodsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PaymentMethodsClient.DeleteByUser`

```go
ctx := context.TODO()
id := paymentmethods.NewPaymentMethodID("paymentMethodName")

read, err := client.DeleteByUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PaymentMethodsClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := paymentmethods.NewBillingAccountPaymentMethodID("billingAccountName", "paymentMethodName")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PaymentMethodsClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := paymentmethods.NewPaymentMethodLinkID("billingAccountName", "billingProfileName", "paymentMethodLinkName")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PaymentMethodsClient.GetByUser`

```go
ctx := context.TODO()
id := paymentmethods.NewPaymentMethodID("paymentMethodName")

read, err := client.GetByUser(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PaymentMethodsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := paymentmethods.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PaymentMethodsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := paymentmethods.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PaymentMethodsClient.ListByUser`

```go
ctx := context.TODO()


// alternatively `client.ListByUser(ctx)` can be used to do batched pagination
items, err := client.ListByUserComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
