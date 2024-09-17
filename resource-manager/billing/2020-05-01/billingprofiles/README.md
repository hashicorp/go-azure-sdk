
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingprofiles` Documentation

The `billingprofiles` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingprofiles"
```


### Client Initialization

```go
client := billingprofiles.NewBillingProfilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingProfilesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingprofiles.BillingProfile{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingProfilesClient.Get`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.Get(ctx, id, billingprofiles.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingProfilesClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, billingprofiles.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, billingprofiles.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
