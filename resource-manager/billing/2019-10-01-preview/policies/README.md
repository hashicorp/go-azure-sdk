
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/policies` Documentation

The `policies` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/policies"
```


### Client Initialization

```go
client := policies.NewPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PoliciesClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := policies.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoliciesClient.GetByCustomer`

```go
ctx := context.TODO()
id := policies.NewCustomerID("billingAccountValue", "customerValue")

read, err := client.GetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoliciesClient.Update`

```go
ctx := context.TODO()
id := policies.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := policies.Policy{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PoliciesClient.UpdateCustomer`

```go
ctx := context.TODO()
id := policies.NewCustomerID("billingAccountValue", "customerValue")

payload := policies.CustomerPolicy{
	// ...
}


read, err := client.UpdateCustomer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
