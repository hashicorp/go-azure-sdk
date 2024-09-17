
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/balances` Documentation

The `balances` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-05-01/balances"
```


### Client Initialization

```go
client := balances.NewBalancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BalancesClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := balances.NewBillingAccountID("billingAccountIdValue")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BalancesClient.GetForBillingPeriodByBillingAccount`

```go
ctx := context.TODO()
id := balances.NewBillingAccountBillingPeriodID("billingAccountIdValue", "billingPeriodValue")

read, err := client.GetForBillingPeriodByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
