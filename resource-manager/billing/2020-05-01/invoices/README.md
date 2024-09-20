
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/invoices` Documentation

The `invoices` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/invoices"
```


### Client Initialization

```go
client := invoices.NewInvoicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InvoicesClient.DownloadBillingSubscriptionInvoice`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionInvoiceID("subscriptionId", "invoiceName")

if err := client.DownloadBillingSubscriptionInvoiceThenPoll(ctx, id, invoices.DefaultDownloadBillingSubscriptionInvoiceOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.DownloadInvoice`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountInvoiceID("billingAccountName", "invoiceName")

if err := client.DownloadInvoiceThenPoll(ctx, id, invoices.DefaultDownloadInvoiceOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.DownloadMultipleBillingProfileInvoices`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountID("billingAccountName")
var payload []string

if err := client.DownloadMultipleBillingProfileInvoicesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.DownloadMultipleBillingSubscriptionInvoices`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionID("subscriptionId")
var payload []string

if err := client.DownloadMultipleBillingSubscriptionInvoicesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.Get`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountInvoiceID("billingAccountName", "invoiceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.GetById`

```go
ctx := context.TODO()
id := invoices.NewInvoiceID("invoiceName")

read, err := client.GetById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.GetBySubscriptionAndInvoiceId`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionInvoiceID("subscriptionId", "invoiceName")

read, err := client.GetBySubscriptionAndInvoiceId(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, invoices.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, invoices.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InvoicesClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := invoices.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id, invoices.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, invoices.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InvoicesClient.ListByBillingSubscription`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionID("subscriptionId")

// alternatively `client.ListByBillingSubscription(ctx, id, invoices.DefaultListByBillingSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingSubscriptionComplete(ctx, id, invoices.DefaultListByBillingSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
