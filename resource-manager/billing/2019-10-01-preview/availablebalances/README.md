
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/availablebalances` Documentation

The `availablebalances` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/availablebalances"
```


### Client Initialization

```go
client := availablebalances.NewAvailableBalancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableBalancesClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := availablebalances.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
