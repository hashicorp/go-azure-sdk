
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/pricesheet` Documentation

The `pricesheet` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/pricesheet"
```


### Client Initialization

```go
client := pricesheet.NewPriceSheetClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheetClient.Download`

```go
ctx := context.TODO()
id := pricesheet.NewBillingProfileInvoiceID("billingAccountValue", "billingProfileValue", "invoiceValue")

if err := client.DownloadThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PriceSheetClient.DownloadByBillingProfile`

```go
ctx := context.TODO()
id := pricesheet.NewBillingProfileID("billingAccountValue", "billingProfileValue")

if err := client.DownloadByBillingProfileThenPoll(ctx, id); err != nil {
	// handle the error
}
```
