
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/pricesheetresults` Documentation

The `pricesheetresults` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/pricesheetresults"
```


### Client Initialization

```go
client := pricesheetresults.NewPriceSheetResultsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PriceSheetResultsClient.PriceSheetGet`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.PriceSheetGet(ctx, id, pricesheetresults.DefaultPriceSheetGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
