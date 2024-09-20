
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transactions` Documentation

The `transactions` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transactions"
```


### Client Initialization

```go
client := transactions.NewTransactionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TransactionsClient.Get`

```go
ctx := context.TODO()
id := transactions.NewTransactionID("billingAccountName", "billingProfileName", "transactionName")

read, err := client.Get(ctx, id, transactions.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransactionsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := transactions.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, transactions.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, transactions.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := transactions.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id, transactions.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, transactions.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionsClient.ListByCustomer`

```go
ctx := context.TODO()
id := transactions.NewCustomerID("billingAccountName", "customerName")

// alternatively `client.ListByCustomer(ctx, id, transactions.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, transactions.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionsClient.ListByInvoice`

```go
ctx := context.TODO()
id := transactions.NewBillingProfileInvoiceID("billingAccountName", "billingProfileName", "invoiceName")

// alternatively `client.ListByInvoice(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransactionsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := transactions.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

// alternatively `client.ListByInvoiceSection(ctx, id, transactions.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, transactions.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
