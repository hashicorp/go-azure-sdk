
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/benefitutilizationsummaries` Documentation

The `benefitutilizationsummaries` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/benefitutilizationsummaries"
```


### Client Initialization

```go
client := benefitutilizationsummaries.NewBenefitUtilizationSummariesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BenefitUtilizationSummariesClient.ListByBillingAccountId`

```go
ctx := context.TODO()
id := benefitutilizationsummaries.NewBillingAccountID("billingAccountIdValue")

// alternatively `client.ListByBillingAccountId(ctx, id, benefitutilizationsummaries.DefaultListByBillingAccountIdOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountIdComplete(ctx, id, benefitutilizationsummaries.DefaultListByBillingAccountIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BenefitUtilizationSummariesClient.ListByBillingProfileId`

```go
ctx := context.TODO()
id := benefitutilizationsummaries.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

// alternatively `client.ListByBillingProfileId(ctx, id, benefitutilizationsummaries.DefaultListByBillingProfileIdOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileIdComplete(ctx, id, benefitutilizationsummaries.DefaultListByBillingProfileIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BenefitUtilizationSummariesClient.ListBySavingsPlanId`

```go
ctx := context.TODO()
id := benefitutilizationsummaries.NewSavingsPlanID("savingsPlanOrderIdValue", "savingsPlanIdValue")

// alternatively `client.ListBySavingsPlanId(ctx, id, benefitutilizationsummaries.DefaultListBySavingsPlanIdOperationOptions())` can be used to do batched pagination
items, err := client.ListBySavingsPlanIdComplete(ctx, id, benefitutilizationsummaries.DefaultListBySavingsPlanIdOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BenefitUtilizationSummariesClient.ListBySavingsPlanOrder`

```go
ctx := context.TODO()
id := benefitutilizationsummaries.NewSavingsPlanOrderID("savingsPlanOrderIdValue")

// alternatively `client.ListBySavingsPlanOrder(ctx, id, benefitutilizationsummaries.DefaultListBySavingsPlanOrderOperationOptions())` can be used to do batched pagination
items, err := client.ListBySavingsPlanOrderComplete(ctx, id, benefitutilizationsummaries.DefaultListBySavingsPlanOrderOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
