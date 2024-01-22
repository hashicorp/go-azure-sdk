
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptions` Documentation

The `billingsubscriptions` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptions"
```


### Client Initialization

```go
client := billingsubscriptions.NewBillingSubscriptionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingSubscriptionsClient.BillingSubscriptionsMerge`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscriptions.BillingSubscriptionMergeRequest{
	// ...
}


if err := client.BillingSubscriptionsMergeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionsClient.BillingSubscriptionsMove`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscriptions.MoveBillingSubscriptionRequest{
	// ...
}


if err := client.BillingSubscriptionsMoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionsClient.BillingSubscriptionsSplit`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscriptions.BillingSubscriptionSplitRequest{
	// ...
}


if err := client.BillingSubscriptionsSplitThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionsClient.BillingSubscriptionsValidateMoveEligibility`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscriptions.MoveBillingSubscriptionRequest{
	// ...
}


read, err := client.BillingSubscriptionsValidateMoveEligibility(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsClient.Delete`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BillingSubscriptionsClient.Get`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingSubscriptionsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingSubscriptionsClient.Update`

```go
ctx := context.TODO()
id := billingsubscriptions.NewBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

payload := billingsubscriptions.BillingSubscription{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
