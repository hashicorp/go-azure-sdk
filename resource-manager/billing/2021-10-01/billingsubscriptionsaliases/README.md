
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptionsaliases` Documentation

The `billingsubscriptionsaliases` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptionsaliases"
```


### Client Initialization

```go
client := billingsubscriptionsaliases.NewBillingSubscriptionsAliasesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingSubscriptionsAliasesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := billingsubscriptionsaliases.NewBillingSubscriptionAliasID("billingAccountName", "billingSubscriptionAliasName")

payload := billingsubscriptionsaliases.BillingSubscriptionAlias{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionsAliasesClient.Get`

```go
ctx := context.TODO()
id := billingsubscriptionsaliases.NewBillingSubscriptionAliasID("billingAccountName", "billingSubscriptionAliasName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsAliasesClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingsubscriptionsaliases.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
