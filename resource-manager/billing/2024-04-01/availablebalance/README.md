
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/availablebalance` Documentation

The `availablebalance` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/availablebalance"
```


### Client Initialization

```go
client := availablebalance.NewAvailableBalanceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableBalanceClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := availablebalance.NewBillingAccountID("billingAccountName")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AvailableBalanceClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := availablebalance.NewBillingProfileID("billingAccountName", "billingProfileName")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
