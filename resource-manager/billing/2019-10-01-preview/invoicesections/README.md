
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoicesections` Documentation

The `invoicesections` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

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
id := invoicesections.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

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
id := invoicesections.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

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
id := invoicesections.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

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
id := invoicesections.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.ListByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoiceSectionsClient.Update`

```go
ctx := context.TODO()
id := invoicesections.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

payload := invoicesections.InvoiceSection{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
