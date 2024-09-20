
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingaccount` Documentation

The `billingaccount` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingaccount"
```


### Client Initialization

```go
client := billingaccount.NewBillingAccountClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingAccountClient.AddPaymentTerms`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")
var payload []PaymentTerm

if err := client.AddPaymentTermsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingAccountClient.AddressValidate`

```go
ctx := context.TODO()

payload := billingaccount.AddressDetails{
	// ...
}


read, err := client.AddressValidate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingAccountClient.CancelPaymentTerms`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")
var payload string

if err := client.CancelPaymentTermsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingAccountClient.ConfirmTransition`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")

read, err := client.ConfirmTransition(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingAccountClient.Get`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingAccountClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, billingaccount.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, billingaccount.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingAccountClient.ListInvoiceSectionsByCreateSubscriptionPermission`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")

// alternatively `client.ListInvoiceSectionsByCreateSubscriptionPermission(ctx, id, billingaccount.DefaultListInvoiceSectionsByCreateSubscriptionPermissionOperationOptions())` can be used to do batched pagination
items, err := client.ListInvoiceSectionsByCreateSubscriptionPermissionComplete(ctx, id, billingaccount.DefaultListInvoiceSectionsByCreateSubscriptionPermissionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingAccountClient.Update`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")

payload := billingaccount.BillingAccountPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingAccountClient.ValidatePaymentTerms`

```go
ctx := context.TODO()
id := billingaccount.NewBillingAccountID("billingAccountName")
var payload []PaymentTerm

read, err := client.ValidatePaymentTerms(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
