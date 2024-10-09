
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/costallocationrules` Documentation

The `costallocationrules` SDK allows for interaction with Azure Resource Manager `costmanagement` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2023-11-01/costallocationrules"
```


### Client Initialization

```go
client := costallocationrules.NewCostAllocationRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CostAllocationRulesClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := costallocationrules.NewBillingAccountID("billingAccountId")

payload := costallocationrules.CostAllocationRuleCheckNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := costallocationrules.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

payload := costallocationrules.CostAllocationRuleDefinition{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRulesClient.Delete`

```go
ctx := context.TODO()
id := costallocationrules.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRulesClient.Get`

```go
ctx := context.TODO()
id := costallocationrules.NewCostAllocationRuleID("billingAccountId", "costAllocationRuleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostAllocationRulesClient.List`

```go
ctx := context.TODO()
id := costallocationrules.NewBillingAccountID("billingAccountId")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
