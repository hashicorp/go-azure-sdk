
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/policy` Documentation

The `policy` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/policy"
```


### Client Initialization

```go
client := policy.NewPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyClient.PoliciesCreateOrUpdateByBillingAccount`

```go
ctx := context.TODO()
id := policy.NewBillingAccountID("billingAccountValue")

payload := policy.BillingAccountPolicy{
	// ...
}


if err := client.PoliciesCreateOrUpdateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PolicyClient.PoliciesCreateOrUpdateByBillingProfile`

```go
ctx := context.TODO()
id := policy.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := policy.BillingProfilePolicy{
	// ...
}


if err := client.PoliciesCreateOrUpdateByBillingProfileThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PolicyClient.PoliciesCreateOrUpdateByCustomer`

```go
ctx := context.TODO()
id := policy.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

payload := policy.CustomerPolicy{
	// ...
}


if err := client.PoliciesCreateOrUpdateByCustomerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PolicyClient.PoliciesCreateOrUpdateByCustomerAtBillingAccount`

```go
ctx := context.TODO()
id := policy.NewCustomerID("billingAccountValue", "customerValue")

payload := policy.CustomerPolicy{
	// ...
}


if err := client.PoliciesCreateOrUpdateByCustomerAtBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PolicyClient.PoliciesGetByBillingAccount`

```go
ctx := context.TODO()
id := policy.NewBillingAccountID("billingAccountValue")

read, err := client.PoliciesGetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.PoliciesGetByBillingProfile`

```go
ctx := context.TODO()
id := policy.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.PoliciesGetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.PoliciesGetByCustomer`

```go
ctx := context.TODO()
id := policy.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

read, err := client.PoliciesGetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.PoliciesGetByCustomerAtBillingAccount`

```go
ctx := context.TODO()
id := policy.NewCustomerID("billingAccountValue", "customerValue")

read, err := client.PoliciesGetByCustomerAtBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.PoliciesGetBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.PoliciesGetBySubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
