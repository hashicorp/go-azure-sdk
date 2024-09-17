
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplanorder` Documentation

The `savingsplanorder` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplanorder"
```


### Client Initialization

```go
client := savingsplanorder.NewSavingsPlanOrderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SavingsPlanOrderClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := savingsplanorder.NewSavingsPlanOrderID("billingAccountValue", "savingsPlanOrderIdValue")

read, err := client.GetByBillingAccount(ctx, id, savingsplanorder.DefaultGetByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SavingsPlanOrderClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := savingsplanorder.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, savingsplanorder.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, savingsplanorder.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
