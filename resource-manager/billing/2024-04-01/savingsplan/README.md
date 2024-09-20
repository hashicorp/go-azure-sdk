
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplan` Documentation

The `savingsplan` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplan"
```


### Client Initialization

```go
client := savingsplan.NewSavingsPlanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SavingsPlanClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := savingsplan.NewSavingsPlanID("billingAccountName", "savingsPlanOrderId", "savingsPlanId")

read, err := client.GetByBillingAccount(ctx, id, savingsplan.DefaultGetByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SavingsPlanClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := savingsplan.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, savingsplan.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, savingsplan.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SavingsPlanClient.ListBySavingsPlanOrder`

```go
ctx := context.TODO()
id := savingsplan.NewSavingsPlanOrderID("billingAccountName", "savingsPlanOrderId")

// alternatively `client.ListBySavingsPlanOrder(ctx, id)` can be used to do batched pagination
items, err := client.ListBySavingsPlanOrderComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SavingsPlanClient.UpdateByBillingAccount`

```go
ctx := context.TODO()
id := savingsplan.NewSavingsPlanID("billingAccountName", "savingsPlanOrderId", "savingsPlanId")

payload := savingsplan.SavingsPlanUpdateRequest{
	// ...
}


if err := client.UpdateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SavingsPlanClient.ValidateUpdateByBillingAccount`

```go
ctx := context.TODO()
id := savingsplan.NewSavingsPlanID("billingAccountName", "savingsPlanOrderId", "savingsPlanId")

payload := savingsplan.SavingsPlanUpdateValidateRequest{
	// ...
}


read, err := client.ValidateUpdateByBillingAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
