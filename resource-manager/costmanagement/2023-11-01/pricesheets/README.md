
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/pricesheets` Documentation

The `pricesheets` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/pricesheets"
```


### Client Initialization

```go
client := pricesheets.NewPriceSheetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheetsClient.PriceSheetDownloadByBillingAccount`

```go
ctx := context.TODO()
id := pricesheets.NewBillingPeriodID("billingAccountIdValue", "billingPeriodValue")

if err := client.PriceSheetDownloadByBillingAccountThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PriceSheetsClient.PriceSheetDownloadByBillingProfile`

```go
ctx := context.TODO()
id := pricesheets.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

if err := client.PriceSheetDownloadByBillingProfileThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PriceSheetsClient.PriceSheetDownloadByInvoice`

```go
ctx := context.TODO()
id := pricesheets.NewInvoiceID("billingAccountValue", "billingProfileValue", "invoiceValue")

if err := client.PriceSheetDownloadByInvoiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```
