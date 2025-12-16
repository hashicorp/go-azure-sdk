
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/costallocationruledefinitions` Documentation

The `costallocationruledefinitions` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2025-03-01/costallocationruledefinitions"
```


### Client Initialization

```go
client := costallocationruledefinitions.NewCostAllocationRuleDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CostAllocationRuleDefinitionsClient.CostAllocationRulesCreateOrUpdate`

```go
ctx := context.TODO()
id := costallocationruledefinitions.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

payload := costallocationruledefinitions.CostAllocationRuleDefinition{
	// ...
}


read, err := client.CostAllocationRulesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRuleDefinitionsClient.CostAllocationRulesDelete`

```go
ctx := context.TODO()
id := costallocationruledefinitions.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

read, err := client.CostAllocationRulesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRuleDefinitionsClient.CostAllocationRulesGet`

```go
ctx := context.TODO()
id := costallocationruledefinitions.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

read, err := client.CostAllocationRulesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRuleDefinitionsClient.CostAllocationRulesList`

```go
ctx := context.TODO()
id := costallocationruledefinitions.NewBillingAccountID("billingAccountId")

// alternatively `client.CostAllocationRulesList(ctx, id)` can be used to do batched pagination
items, err := client.CostAllocationRulesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
