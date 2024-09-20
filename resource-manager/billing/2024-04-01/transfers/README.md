
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transfers` Documentation

The `transfers` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transfers"
```


### Client Initialization

```go
client := transfers.NewTransfersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TransfersClient.Cancel`

```go
ctx := context.TODO()
id := transfers.NewInvoiceSectionTransferID("billingAccountName", "billingProfileName", "invoiceSectionName", "transferName")

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
id := transfers.NewInvoiceSectionTransferID("billingAccountName", "billingProfileName", "invoiceSectionName", "transferName")

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
id := transfers.NewInvoiceSectionTransferID("billingAccountName", "billingProfileName", "invoiceSectionName", "transferName")

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
id := transfers.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

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
id := transfers.NewCustomerTransferID("billingAccountName", "billingProfileName", "customerName", "transferName")

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
id := transfers.NewCustomerTransferID("billingAccountName", "billingProfileName", "customerName", "transferName")

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
id := transfers.NewCustomerTransferID("billingAccountName", "billingProfileName", "customerName", "transferName")

payload := transfers.PartnerInitiateTransferRequest{
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
id := transfers.NewBillingProfileCustomerID("billingAccountName", "billingProfileName", "customerName")

// alternatively `client.PartnerTransfersList(ctx, id)` can be used to do batched pagination
items, err := client.PartnerTransfersListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
