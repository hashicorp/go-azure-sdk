
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingsubscription` Documentation

The `billingsubscription` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingsubscription"
```


### Client Initialization

```go
client := billingsubscription.NewBillingSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingSubscriptionClient.AliasesCreateOrUpdate`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingSubscriptionAliasID("billingAccountValue", "billingSubscriptionAliasValue")

payload := billingsubscription.BillingSubscriptionAlias{
	// ...
}


if err := client.AliasesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.AliasesGet`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingSubscriptionAliasID("billingAccountValue", "billingSubscriptionAliasValue")

read, err := client.AliasesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionClient.AliasesListByBillingAccount`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountID("billingAccountValue")

// alternatively `client.AliasesListByBillingAccount(ctx, id, billingsubscription.DefaultAliasesListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.AliasesListByBillingAccountComplete(ctx, id, billingsubscription.DefaultAliasesListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.Cancel`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.CancelSubscriptionRequest{
	// ...
}


if err := client.CancelThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.Delete`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.Get`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

read, err := client.Get(ctx, id, billingsubscription.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingProfileBillingSubscriptionID("billingAccountValue", "billingProfileValue", "billingSubscriptionValue")

read, err := client.GetByBillingProfile(ctx, id, billingsubscription.DefaultGetByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, billingsubscription.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, billingsubscription.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, billingsubscription.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, billingsubscription.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id, billingsubscription.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, billingsubscription.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.ListByCustomerAtBillingAccount`

```go
ctx := context.TODO()
id := billingsubscription.NewCustomerID("billingAccountValue", "customerValue")

// alternatively `client.ListByCustomerAtBillingAccount(ctx, id, billingsubscription.DefaultListByCustomerAtBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerAtBillingAccountComplete(ctx, id, billingsubscription.DefaultListByCustomerAtBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingsubscription.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

// alternatively `client.ListByEnrollmentAccount(ctx, id, billingsubscription.DefaultListByEnrollmentAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id, billingsubscription.DefaultListByEnrollmentAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingsubscription.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, billingsubscription.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, billingsubscription.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionClient.Merge`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.BillingSubscriptionMergeRequest{
	// ...
}


if err := client.MergeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.Move`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.MoveBillingSubscriptionRequest{
	// ...
}


if err := client.MoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.Split`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.BillingSubscriptionSplitRequest{
	// ...
}


if err := client.SplitThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.Update`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.BillingSubscriptionPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionClient.ValidateMoveEligibility`

```go
ctx := context.TODO()
id := billingsubscription.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscription.MoveBillingSubscriptionRequest{
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
