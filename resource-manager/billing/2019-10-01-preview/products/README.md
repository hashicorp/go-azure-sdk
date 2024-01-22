
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/products` Documentation

The `products` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/products"
```


### Client Initialization

```go
client := products.NewProductsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProductsClient.Get`

```go
ctx := context.TODO()
id := products.NewInvoiceSectionProductID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "productValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.GetByCustomer`

```go
ctx := context.TODO()
id := products.NewProductID("billingAccountValue", "customerValue", "productValue")

read, err := client.GetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := products.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, products.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, products.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductsClient.ListByCustomer`

```go
ctx := context.TODO()
id := products.NewCustomerID("billingAccountValue", "customerValue")

read, err := client.ListByCustomer(ctx, id, products.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := products.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

read, err := client.ListByInvoiceSection(ctx, id, products.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.Transfer`

```go
ctx := context.TODO()
id := products.NewInvoiceSectionProductID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "productValue")

payload := products.TransferProductRequestProperties{
	// ...
}


read, err := client.Transfer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.UpdateAutoRenewByInvoiceSection`

```go
ctx := context.TODO()
id := products.NewInvoiceSectionProductID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "productValue")

payload := products.UpdateAutoRenewRequest{
	// ...
}


read, err := client.UpdateAutoRenewByInvoiceSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
