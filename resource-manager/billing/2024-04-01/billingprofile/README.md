
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingprofile` Documentation

The `billingprofile` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingprofile"
```


### Client Initialization

```go
client := billingprofile.NewBillingProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingProfileClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := billingprofile.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingprofile.BillingProfile{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingProfileClient.Delete`

```go
ctx := context.TODO()
id := billingprofile.NewBillingProfileID("billingAccountValue", "billingProfileValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BillingProfileClient.Get`

```go
ctx := context.TODO()
id := billingprofile.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingProfileClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingprofile.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, billingprofile.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, billingprofile.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingProfileClient.ValidateDeleteEligibility`

```go
ctx := context.TODO()
id := billingprofile.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.ValidateDeleteEligibility(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
