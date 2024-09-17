
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transaction` Documentation

The `transaction` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transaction"
```


### Client Initialization

```go
client := transaction.NewTransactionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TransactionClient.GetTransactionSummaryByInvoice`

```go
ctx := context.TODO()
id := transaction.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

read, err := client.GetTransactionSummaryByInvoice(ctx, id, transaction.DefaultGetTransactionSummaryByInvoiceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransactionClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := transaction.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, transaction.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, transaction.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionClient.ListByCustomer`

```go
ctx := context.TODO()
id := transaction.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id, transaction.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, transaction.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionClient.ListByInvoice`

```go
ctx := context.TODO()
id := transaction.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

// alternatively `client.ListByInvoice(ctx, id, transaction.DefaultListByInvoiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceComplete(ctx, id, transaction.DefaultListByInvoiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := transaction.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, transaction.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, transaction.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionClient.TransactionsDownloadByInvoice`

```go
ctx := context.TODO()
id := transaction.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

if err := client.TransactionsDownloadByInvoiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```
