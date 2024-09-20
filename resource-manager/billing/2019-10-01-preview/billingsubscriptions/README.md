
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingsubscriptions` Documentation

The `billingsubscriptions` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

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
id := billingsubscriptions.NewInvoiceSectionBillingSubscriptionID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingSubscriptionName")

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
id := billingsubscriptions.NewCustomerBillingSubscriptionID("billingAccountName", "customerName", "billingSubscriptionName")

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
id := billingsubscriptions.NewBillingAccountID("billingAccountName")

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
id := billingsubscriptions.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionsClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingsubscriptions.NewCustomerID("billingAccountName", "customerName")

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
id := billingsubscriptions.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
