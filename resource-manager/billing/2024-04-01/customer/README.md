
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/customer` Documentation

The `customer` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/customer"
```


### Client Initialization

```go
client := customer.NewCustomerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomerClient.Get`

```go
ctx := context.TODO()
id := customer.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomerClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := customer.NewCustomerID("billingAccountValue", "customerValue")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomerClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := customer.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, customer.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, customer.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CustomerClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := customer.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, customer.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, customer.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
