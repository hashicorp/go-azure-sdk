
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingprofiles` Documentation

The `billingprofiles` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingprofiles"
```


### Client Initialization

```go
client := billingprofiles.NewBillingProfilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingProfilesClient.Create`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingprofiles.BillingProfileCreationRequest{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
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


### Example Usage: `BillingProfilesClient.GetEligibilityToDetachPaymentMethod`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.GetEligibilityToDetachPaymentMethod(ctx, id)
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

read, err := client.ListByBillingAccount(ctx, id, billingprofiles.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingProfilesClient.Update`

```go
ctx := context.TODO()
id := billingprofiles.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingprofiles.BillingProfile{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
