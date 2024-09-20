
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/customers` Documentation

The `customers` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/customers"
```


### Client Initialization

```go
client := customers.NewCustomersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomersClient.Get`

```go
ctx := context.TODO()
id := customers.NewCustomerID("billingAccountName", "customerName")

read, err := client.Get(ctx, id, customers.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomersClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := customers.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, customers.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, customers.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CustomersClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := customers.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id, customers.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, customers.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
