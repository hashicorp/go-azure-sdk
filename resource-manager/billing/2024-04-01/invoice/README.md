
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/invoice` Documentation

The `invoice` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/invoice"
```


### Client Initialization

```go
client := invoice.NewInvoiceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InvoiceClient.Amend`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

if err := client.AmendThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.DownloadByBillingAccount`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

if err := client.DownloadByBillingAccountThenPoll(ctx, id, invoice.DefaultDownloadByBillingAccountOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.DownloadByBillingSubscription`

```go
ctx := context.TODO()
id := invoice.NewBillingSubscriptionInvoiceID("subscriptionIdValue", "invoiceValue")

if err := client.DownloadByBillingSubscriptionThenPoll(ctx, id, invoice.DefaultDownloadByBillingSubscriptionOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.DownloadDocumentsByBillingAccount`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountID("billingAccountValue")
var payload []DocumentDownloadRequest

if err := client.DownloadDocumentsByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.DownloadDocumentsByBillingSubscription`

```go
ctx := context.TODO()
id := invoice.NewBillingSubscriptionID("subscriptionIdValue")
var payload []DocumentDownloadRequest

if err := client.DownloadDocumentsByBillingSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.DownloadSummaryByBillingAccount`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

if err := client.DownloadSummaryByBillingAccountThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceClient.Get`

```go
ctx := context.TODO()
id := invoice.NewInvoiceID("invoiceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceClient.GetByBillingSubscription`

```go
ctx := context.TODO()
id := invoice.NewBillingSubscriptionInvoiceID("subscriptionIdValue", "invoiceValue")

read, err := client.GetByBillingSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := invoice.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, invoice.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, invoice.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InvoiceClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := invoice.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, invoice.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, invoice.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InvoiceClient.ListByBillingSubscription`

```go
ctx := context.TODO()
id := invoice.NewBillingSubscriptionID("subscriptionIdValue")

// alternatively `client.ListByBillingSubscription(ctx, id, invoice.DefaultListByBillingSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingSubscriptionComplete(ctx, id, invoice.DefaultListByBillingSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
