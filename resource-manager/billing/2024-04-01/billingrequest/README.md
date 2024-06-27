
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingrequest` Documentation

The `billingrequest` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingrequest"
```


### Client Initialization

```go
client := billingrequest.NewBillingRequestClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRequestClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := billingrequest.NewBillingRequestID("billingRequestValue")

payload := billingrequest.BillingRequest{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRequestClient.Get`

```go
ctx := context.TODO()
id := billingrequest.NewBillingRequestID("billingRequestValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRequestClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingrequest.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, billingrequest.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, billingrequest.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRequestClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingrequest.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, billingrequest.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, billingrequest.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRequestClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingrequest.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id, billingrequest.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, billingrequest.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRequestClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingrequest.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, billingrequest.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, billingrequest.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRequestClient.ListByUser`

```go
ctx := context.TODO()


// alternatively `client.ListByUser(ctx, billingrequest.DefaultListByUserOperationOptions())` can be used to do batched pagination
items, err := client.ListByUserComplete(ctx, billingrequest.DefaultListByUserOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
