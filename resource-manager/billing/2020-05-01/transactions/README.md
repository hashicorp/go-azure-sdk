
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/transactions` Documentation

The `transactions` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/transactions"
```


### Client Initialization

```go
client := transactions.NewTransactionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TransactionsClient.ListByInvoice`

```go
ctx := context.TODO()
id := transactions.NewBillingAccountInvoiceID("billingAccountValue", "invoiceValue")

// alternatively `client.ListByInvoice(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
