
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoicesections` Documentation

The `invoicesections` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoicesections"
```


### Client Initialization

```go
client := invoicesections.NewInvoiceSectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InvoiceSectionsClient.Create`

```go
ctx := context.TODO()
id := invoicesections.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

payload := invoicesections.InvoiceSectionCreationRequest{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoiceSectionsClient.ElevateToBillingProfile`

```go
ctx := context.TODO()
id := invoicesections.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

read, err := client.ElevateToBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceSectionsClient.Get`

```go
ctx := context.TODO()
id := invoicesections.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceSectionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := invoicesections.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InvoiceSectionsClient.Update`

```go
ctx := context.TODO()
id := invoicesections.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

payload := invoicesections.InvoiceSection{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
