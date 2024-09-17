
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transfers` Documentation

The `transfers` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transfers"
```


### Client Initialization

```go
client := transfers.NewTransfersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TransfersClient.Cancel`

```go
ctx := context.TODO()
id := transfers.NewInvoiceSectionTransferID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "transferValue")

read, err := client.Cancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.Get`

```go
ctx := context.TODO()
id := transfers.NewInvoiceSectionTransferID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "transferValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.Initiate`

```go
ctx := context.TODO()
id := transfers.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

payload := transfers.InitiateTransferRequest{
	// ...
}


read, err := client.Initiate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.List`

```go
ctx := context.TODO()
id := transfers.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TransfersClient.PartnerTransfersCancel`

```go
ctx := context.TODO()
id := transfers.NewCustomerTransferID("billingAccountValue", "billingProfileValue", "customerValue", "transferValue")

read, err := client.PartnerTransfersCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.PartnerTransfersGet`

```go
ctx := context.TODO()
id := transfers.NewCustomerTransferID("billingAccountValue", "billingProfileValue", "customerValue", "transferValue")

read, err := client.PartnerTransfersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.PartnerTransfersInitiate`

```go
ctx := context.TODO()
id := transfers.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

payload := transfers.InitiateTransferRequest{
	// ...
}


read, err := client.PartnerTransfersInitiate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TransfersClient.PartnerTransfersList`

```go
ctx := context.TODO()
id := transfers.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.PartnerTransfersList(ctx, id)` can be used to do batched pagination
items, err := client.PartnerTransfersListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
