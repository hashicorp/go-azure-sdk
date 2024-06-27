
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/product` Documentation

The `product` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/product"
```


### Client Initialization

```go
client := product.NewProductClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProductClient.Get`

```go
ctx := context.TODO()
id := product.NewProductID("billingAccountValue", "productValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := product.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, product.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, product.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := product.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, product.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, product.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductClient.ListByCustomer`

```go
ctx := context.TODO()
id := product.NewCustomerID("billingAccountValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id, product.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, product.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := product.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, product.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, product.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductClient.Move`

```go
ctx := context.TODO()
id := product.NewProductID("billingAccountValue", "productValue")

payload := product.MoveProductRequest{
	// ...
}


if err := client.MoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ProductClient.Update`

```go
ctx := context.TODO()
id := product.NewProductID("billingAccountValue", "productValue")

payload := product.ProductPatch{
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


### Example Usage: `ProductClient.ValidateMoveEligibility`

```go
ctx := context.TODO()
id := product.NewProductID("billingAccountValue", "productValue")

payload := product.MoveProductRequest{
	// ...
}


read, err := client.ValidateMoveEligibility(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
