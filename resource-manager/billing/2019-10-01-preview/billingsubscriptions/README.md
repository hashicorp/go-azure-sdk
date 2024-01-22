
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingsubscriptions` Documentation

The `billingsubscriptions` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingsubscriptions"
```


### Client Initialization

```go
client := billingsubscriptions.NewBillingSubscriptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingSubscriptionsClient.Get`

```go
ctx := context.TODO()
id := billingsubscriptions.NewInvoiceSectionBillingSubscriptionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingSubscriptionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsClient.GetByCustomer`

```go
ctx := context.TODO()
id := billingsubscriptions.NewCustomerBillingSubscriptionID("billingAccountValue", "customerValue", "billingSubscriptionValue")

read, err := client.GetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.ListByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingsubscriptions.NewCustomerID("billingAccountValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingsubscriptions.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

read, err := client.ListByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
