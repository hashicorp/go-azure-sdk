
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/products` Documentation

The `products` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/products"
```


### Client Initialization

```go
client := products.NewProductsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProductsClient.Get`

```go
ctx := context.TODO()
id := products.NewProductID("billingAccountValue", "productValue")

read, err := client.Get(ctx, id)
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


### Example Usage: `ProductsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := products.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, products.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, products.DefaultListByBillingProfileOperationOptions())
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

// alternatively `client.ListByCustomer(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := products.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, products.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, products.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductsClient.Move`

```go
ctx := context.TODO()
id := products.NewProductID("billingAccountValue", "productValue")

payload := products.TransferProductRequestProperties{
	// ...
}


read, err := client.Move(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductsClient.Update`

```go
ctx := context.TODO()
id := products.NewProductID("billingAccountValue", "productValue")

payload := products.Product{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
