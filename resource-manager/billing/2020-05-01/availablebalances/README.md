
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/availablebalances` Documentation

The `availablebalances` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/availablebalances"
```


### Client Initialization

```go
client := availablebalances.NewAvailableBalancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableBalancesClient.Get`

```go
ctx := context.TODO()
id := availablebalances.NewBillingProfileID("billingAccountName", "billingProfileName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
