
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/pricesheets` Documentation

The `pricesheets` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-08-01/pricesheets"
```


### Client Initialization

```go
client := pricesheets.NewPriceSheetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheetsClient.PriceSheetDownload`

```go
ctx := context.TODO()
id := pricesheets.NewInvoiceID("billingAccountName", "billingProfileName", "invoiceName")

if err := client.PriceSheetDownloadThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PriceSheetsClient.PriceSheetDownloadByBillingProfile`

```go
ctx := context.TODO()
id := pricesheets.NewBillingProfileID("billingAccountId", "billingProfileId")

if err := client.PriceSheetDownloadByBillingProfileThenPoll(ctx, id); err != nil {
	// handle the error
}
```
